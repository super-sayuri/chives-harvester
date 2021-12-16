package db

import (
	"context"
	"encoding/json"
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/model"
	"strconv"
	"strings"
)

func GetGroupIds(ctx context.Context) (groupIds []int64, err error) {
	log := conf.GetLog(ctx)
	groups, err := GetRedisDb().LRange(ctx, DB_KEY_GROUPS, 0, -1).Result()
	if err != nil {
		return
	}
	groupIds = make([]int64, 0, len(groups))
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

func GetCryptoItems(ctx context.Context) ([]*model.GoodsItem, error) {
	log := conf.GetLog(ctx)
	rdb := GetRedisDb()
	cryptoDb, err := rdb.HGetAll(ctx, DB_KEY_CRYPTO_ITEMS).Result()
	if err != nil {
		return nil, err
	}
	var cryptoItems []*model.GoodsItem
	for _, itemStr := range cryptoDb {
		item := &model.GoodsItem{}
		err = json.Unmarshal([]byte(itemStr), item)
		if err != nil {
			log.Error("cannot marshal item json ", err)
			continue
		}
		cryptoItems = append(cryptoItems, item)
	}
	return cryptoItems, nil
}

func GetCryptoItemById(ctx context.Context, id string) (*model.GoodsItem, error) {
	rdb := GetRedisDb()
	id = strings.ToLower(id)
	cryptoDb, err := rdb.HGet(ctx, DB_KEY_CRYPTO_ITEMS, id).Result()
	if err != nil {
		return nil, err
	}
	item := &model.GoodsItem{}
	err = json.Unmarshal([]byte(cryptoDb), item)
	if err != nil {
		return nil, err
	}
	return item, nil
}
