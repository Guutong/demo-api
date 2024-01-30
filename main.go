package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/guutong/demo-gin/auth"
	"github.com/guutong/demo-gin/user"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Can not Load env!")
	}

	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Can not connect DB!")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	db.AutoMigrate(
		&user.User{},
	)

	r.GET("/tokenz", auth.GetToken(os.Getenv("WIFI_SECRET")))
	authRoute := r.Group("", auth.AuthMiddleware(os.Getenv("WIFI_SECRET")))

	userHandler := user.NewUserHandler(db)
	authRoute.POST("/users", userHandler.NewUser)
	authRoute.GET("/users", userHandler.GetUser)
	authRoute.DELETE("/users/:id", userHandler.DeleteUser)
	authRoute.PATCH("/users/:id", userHandler.UpdateUser)

	r.Run()
}
