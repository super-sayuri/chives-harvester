早上/中午/晚上好，现在是{{ .Datetime }}
{{ range .Items }}
当前 {{ .Name }} 的价格是 {{ .Price }} {{ .Currency }}
24H内涨跌幅度为 {{ .ChangeVal }}，折合百分比 {{ .ChangePercent }}
最高价 {{ .High24h }}，最低价 {{ .Low24h }}
{{ end }}
