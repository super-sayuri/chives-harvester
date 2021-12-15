早上/中午/晚上好，现在是{{ .Datetime }}
{{ range .Items }}
现在 {{ .Name }} 的价格是 {{ .Price }} {{ .Currency }}
比昨天的成交价变化 {{ .ChangeVal }}，折合百分比 {{ .ChangePercent }}
24H内最高价 {{ .High24h }}，最低价 {{ .Low24h }}
{{ end }}
这里应该有一句鸡汤，下个版本再更新吧。