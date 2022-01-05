package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"wmform/modules"
)

const dsn = "root:ol0Ya29n8mxLVdWC@(144.34.204.227:27306)/wmform?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(db)

	var api = new(modules.ApiInfo)

	db.First(api)

	fmt.Println(api)
}
