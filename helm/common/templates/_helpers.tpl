
{{- define "common.name" -}}
  {{- default .Release.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "common.fullname" -}}
  {{- if .Values.fullnameOverride -}}
    {{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
  {{- else -}}
    {{- $name := default .Chart.Name .Values.nameOverride -}}
    {{- if contains $name .Release.Name -}}
      {{- .Release.Name | trunc 63 | trimSuffix "-" -}}
    {{- else -}}
      {{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
    {{- end -}}
  {{- end -}}
{{- end -}}

{{- define "common.chart" -}}
  {{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "common.metadata" -}}
name: {{ include "common.fullname" . }}
labels:
  app.kubernetes.io/name: {{ .Chart.Name }}
  helm.sh/chart: {{ include "common.chart" . }}
  app.kubernetes.io/managed-by: {{ .Release.Service }}
  app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{- define "common.handler.environmentVariables" -}}
{{- if .Values.env -}}
env:
  {{- range $key, $val := .Values.env }}
  - name: {{ $key }}
    value: {{ $val }}
  {{- end }}
{{- end -}}
{{- end -}}

{{- define "common.handler.resources" -}}
{{- if not .Values.resources -}}
{{ fail "A valid Values resources is required" }}
{{- end -}}
resources:
  limits:
    cpu: {{ required "A valid Values.resources.limits.cpu is required" .Values.resources.limits.cpu }}
    memory: {{ required "A valid Values.resources.limits.memory is required" .Values.resources.limits.memory }}
  requests:
    cpu: {{ required "A valid Values.resources.requests.cpu is required" .Values.resources.requests.cpu }}
    memory: {{ required "A valid Values.resources.requests.memory is required" .Values.resources.requests.memory }}
{{- end -}}

{{- define "common.handler.volumeMounts" -}}
{{- if or .Values.configmap.enabled .Values.storage.enabled -}}
volumeMounts:
{{- if .Values.configmap.enabled }}
  - name: {{ include "common.fullname" . }}-configmap
    mountPath: {{ .Values.configmap.mountPath }}
{{- end }}
{{- if .Values.storage.enabled }}
  - name: {{ include "common.fullname" . }}-volume
    mountPath: {{ .Values.storage.mountPath }}
{{- end }}
{{- end -}}
{{- end -}}

{{- define "common.handler.volume" -}}
{{- if or .Values.configmap.enabled .Values.storage.enabled -}}
volumes:
{{- if .Values.configmap.enabled }}
  - configMap:
      defaultMode: 493
      name: {{ include "common.fullname" . }}
    name: {{ include "common.fullname" . }}-configmap
{{- end }}
{{- if .Values.storage.enabled }}
  - name: {{ include "common.fullname" . }}-volume
    persistentVolumeClaim:
      claimName: {{ include "common.fullname" . }}
{{- end }}
{{- end -}}
{{- end -}}
