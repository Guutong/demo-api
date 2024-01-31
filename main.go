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

	userRepository := user.NewUserRepository(db)
	userHandler := user.NewUserHandler(userRepository)
	authRoute.POST("/users", user.NewGinHandler(userHandler.NewUser))
	authRoute.GET("/users", user.NewGinHandler(userHandler.GetUser))
	authRoute.DELETE("/users/:id", user.NewGinHandler(userHandler.DeleteUser))
	authRoute.PATCH("/users/:id", user.NewGinHandler(userHandler.UpdateUser))

	r.Run()
}
