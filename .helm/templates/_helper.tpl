{{/* Define helper functions or variables here */}}

{{- define "go-auth.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name -}}
{{- end }}

{{- define "go-auth.name" -}}
{{- printf "%s" .Release.Name -}}
{{- end }}


{{- define "go-auth.chart" -}}
{{- printf "%s" .Chart.Name -}}
{{- end }}


