package router

import (
	"errors"
	"gorm.io/gorm"
	"sayuri_crypto_bot/db"
)

func listALL() ([]*CRouter, error) {
	x := db.Gorm()
	res := make([]*CRouter, 0)
	err := x.Find(&res).Error
	return res, err
}

func listOne(key string) (*CRouter, error) {
	x := db.Gorm()
	res := &CRouter{}
	err := x.First(&res, "name=?", key).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return res, nil
}

func update(key string, obj *CRouter) error {
	x := db.Gorm()
	err :=
		x.Model(obj).Where("name=?", key).Update(
			"value", obj.Value).Error
	return err
}

func insert(obj *CRouter) error {
	x := db.Gorm()
	err := x.Create(obj).Error
	return err
}

func deleteOne(key string) error {
	x := db.Gorm()
	err := x.Delete(&CRouter{}, "name=?", key).Error
	return err
}
