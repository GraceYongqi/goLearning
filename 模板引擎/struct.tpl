package {{.PackageName}}

type {{.Data.Name}} struct {
	{{- range $index, $value := .Data.Fields }}
    {{$value.Name}} {{$value.Type}} {{$value.Tag}}
	{{- end }}
}