package db

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sayuri_crypto_bot/conf"
)

type GormModel struct {
	ID        string `json:"id" gorm:"primaryKey"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `json:"updatedAt,omitempty" gorm:"autoUploadTime:milli"`
	DeletedAt int64  `json:"-"`
}

var _g *gorm.DB

func gormInit(db *conf.DbConfig) (err error) {
	if db.Driver == "sqlite3" {
		_g, err = gorm.Open(sqlite.Open(db.Url), &gorm.Config{})
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("database driver not support")
}

func Gorm() *gorm.DB {
	return _g
}
