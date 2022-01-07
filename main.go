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
		Logger: nil,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(db)

	var api = new(modules.ApiInfo)
	db.First(api)
	fmt.Println(api)
	fmt.Println("---- api ----------")

	var DataControl = new(modules.DataControl)
	db.First(DataControl)
	fmt.Println(DataControl)
	fmt.Println("---- DataControl ----------")

	var DataSet = new(modules.DataSet)
	db.First(DataSet)
	fmt.Println(DataSet)
	fmt.Println("---- DataSet ----------")

	var DataSource = new(modules.DataSource)
	db.First(DataSource)
	fmt.Println(DataSource)
	fmt.Println("---- DataSource ----------")

	var ApplicationDetail = new(modules.ApplicationDetail)
	db.First(ApplicationDetail)
	fmt.Println(ApplicationDetail)
	fmt.Println("---- ApplicationDetail ----------")

	var ApplicationSummary = new(modules.ApplicationSummary)
	db.First(ApplicationSummary)
	fmt.Println(ApplicationSummary)
	fmt.Println("---- ApplicationSummary ----------")

	var ApplicationApiinfo = new(modules.ApplicationApiinfo)
	db.First(ApplicationApiinfo)
	fmt.Println(ApplicationApiinfo)
	fmt.Println("---- ApplicationApiinfo ----------")

	var ApplicationDatasource = new(modules.ApplicationDatasource)
	db.First(ApplicationDatasource)
	fmt.Println(ApplicationDatasource)
	fmt.Println("---- ApplicationDatasource ----------")

	var AssociatedGroup = new(modules.AssociatedGroup)
	db.First(AssociatedGroup)
	fmt.Println(AssociatedGroup)
	fmt.Println("---- AssociatedGroup ----------")

	var BussinessVal = new(modules.BussinessVal)
	db.First(BussinessVal)
	fmt.Println(BussinessVal)
	fmt.Println("---- BussinessVal ----------")

	var Group = new(modules.Group)
	db.First(Group)
	fmt.Println(Group)
	fmt.Println("---- Group ----------")

	var PageDetail = new(modules.PageDetail)
	db.First(PageDetail)
	fmt.Println(PageDetail)
	fmt.Println("---- PageDetail ----------")

	var PageSummary = new(modules.PageSummary)
	db.First(PageSummary)
	fmt.Println(PageSummary)
	fmt.Println("---- PageSummary ----------")

	var UserTest = new(modules.UserTest)
	db.First(UserTest)
	fmt.Println(UserTest)
	fmt.Println("---- UserTest ----------")

}
