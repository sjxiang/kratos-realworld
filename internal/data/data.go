package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"

	
	"gorm.io/driver/mysql"

	"kratos-realworld/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewUserRepo, NewProfileRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db:db}, cleanup, nil
}



func NewDB(c *conf.Data) *gorm.DB {
	// sudo cat /etc/mysql/debian.cnf
	// vi ~/my.cnf

	
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{})
  	if err != nil {
    	panic(err)
  	}

	  // 迁移 schema
  if err := db.AutoMigrate(); err != nil {
	  panic(err)
  }

  return db
}