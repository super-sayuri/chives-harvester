package model

type GoodsItem struct {
	Id    string            `json:"id"`
	Type  string            `json:"type"`
	Alias map[string]string `json:"alias"`
}

func (g *GoodsItem) GetAlias(key string) string {
	res, ok := g.Alias[key]
	if !ok {
		return ""
	}
	return res
}
