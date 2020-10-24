// Copyright 2020 Vasyl Naumenko. All rights reserved.

package main

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Миграция схем
	db.AutoMigrate(&Product{})

	// Создание
	db.Create(&Product{Code: "D42", Price: 100})

	// Чтение
	var product Product
	tx := db.First(&product, 1) // find product with integer primary key
	if tx.Error != nil {
		tx.Logger.Error(context.Background(), tx.Error.Error())
	}
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Обновление - обновить цену товара в 200
	db.Model(&product).Update("Price", 200)
	// Обновление - обновить несколько полей
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Удаление - удаление товара
	db.Delete(&product, []string{"code = ?"}, "D42")
}
