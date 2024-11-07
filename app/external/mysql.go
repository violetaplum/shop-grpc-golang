package external

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shop-grpc-golang/domain"
)

type externalMysql struct {
	shopCore *gorm.DB
}

func (e *externalMysql) ShopCore() *gorm.DB {
	return e.shopCore
}

func NewMysql() domain.ExternalMysql {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "aspyn:1234@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &externalMysql{shopCore: db}
}
