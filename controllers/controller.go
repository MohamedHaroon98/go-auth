package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MohamedHaroon98/go-auth/models"
	"github.com/MohamedHaroon98/go-auth/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var collection *mongo.Collection

func init() {
	client := utils.ConnectDB()
	collection = client.Database("accounts").Collection("accounts")
}

func RegisterAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var account models.Account
	_ = json.NewDecoder(r.Body).Decode(&account)

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	account.Password = string(hashedPassword)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, account)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result.InsertedID)
}

func LoginAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginRequest models.Account
	_ = json.NewDecoder(r.Body).Decode(&loginRequest)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var storedAccount models.Account
	err := collection.FindOne(ctx, bson.M{"username": loginRequest.Username}).Decode(&storedAccount)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Account not found")
			return
		}
		log.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedAccount.Password), []byte(loginRequest.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid password")
		return
	}

	fmt.Fprintf(w, "Login successful")
}

func GetAllUsernames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	var usernames []string
	for cursor.Next(ctx) {
		var account models.Account
		err := cursor.Decode(&account)
		if err != nil {
			log.Fatal(err)
		}
		usernames = append(usernames, account.Username)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Usernames:", usernames)
	json.NewEncoder(w).Encode(usernames)
}
