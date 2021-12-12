package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tetsucceed/byrdy/pkg/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func ShowEditPost(rc *gin.Context) {
	rc.JSON(http.StatusOK, gin.H{"title": "Hello"})
}

func main() {

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsnUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", dbHost,
		dbUser, dbPassword, dbName)

	pdb, err := gorm.Open(postgres.Open(dsnUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	_ = pdb.AutoMigrate(&db.Structure{})

	env := db.Env{
		Links: db.LinkDbModel{DB: pdb},
	}

	// connect to postgresql and apply migrations
	log.Print("Starting")

	r := gin.Default()
	r.GET("/", ShowEditPost)
	r.GET("/api/get", env.GetLink)
	r.POST("/api/create", env.SaveShort)

	_ = r.Run(":8080")
}
