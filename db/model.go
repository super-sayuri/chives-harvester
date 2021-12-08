package db

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"sayuri_crypto_bot/model"
	"strconv"
)

func GetGroupIds() (groupIds []int64, err error) {
	groups, err := GetRedisDb().LRange(context.Background(), "syr_groups_to_send", 0, -1).Result()
	if err != nil {
		return
	}
	groupIds = make([]int64, len(groups))
	for _, group := range groups {
		id, err := strconv.Atoi(group)
		if err != nil {
			log.Error("cannot get group id: ", err)
			continue
		}
		groupIds = append(groupIds, int64(id))
	}
	return
}

func GetCryptoItems() ([]*model.GoodsItem, error) {
	rdb := GetRedisDb()
	cryptoDb, err := rdb.HGetAll(context.Background(), "syr_crypto").Result()
	if err != nil {
		return nil, err
	}
	var cryptoItems []*model.GoodsItem
	for _, itemStr := range cryptoDb {
		var item *model.GoodsItem
		err = json.Unmarshal([]byte(itemStr), item)
		if err != nil {
			log.Error("cannot marshal item json ", err)
			continue
		}
		cryptoItems = append(cryptoItems, item)
	}
	return cryptoItems, nil
}
