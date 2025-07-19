{{- define "inforce-task.name" -}}
inforce-task
{{- end }}

{{- define "inforce-task.fullname" -}}
{{ include "inforce-task.name" . }}-{{ .Release.Name }}
{{- end }}
