// package main

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	app *gin.Engine
)

var DB *gorm.DB

type Article struct {
	gorm.Model
	Title string
	Slug  string `gorm:"unique_index"`
	Desc  string `gorm:"type:text"`
}

func InitDB() {
	// url := os.Getenv("URL")
	var err error
	DB, err = gorm.Open("postgres", "host=147.139.210.135 port=5432 user=zaki dbname=zaki03 password=zaki123")
	if err != nil {
		panic("failed to connect database")
	}
}

func Migrate() {
	DB.AutoMigrate(&Article{})
}

func SelectAll() []Article {
	items := []Article{}
	DB.Raw("SELECT * FROM articles").Scan(&items)
	return items

	// items := []Article{}
	// return DB.Find(&items)
}

func init() {
	app = gin.New()
	InitDB()
	Migrate()
	res := SelectAll()
	app.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Berhasil",
			"data":   res,
		})
	})
	defer DB.Close()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
