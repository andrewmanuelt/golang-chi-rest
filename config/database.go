package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {
	dsn := dsn()

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database...")

	return conn
}

func dsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		ConfigField("db", "user"),
		ConfigField("db", "pass"),
		ConfigField("db", "host"),
		ConfigField("db", "port"),
		ConfigField("db", "dbname"),
	)

	return dsn
}
