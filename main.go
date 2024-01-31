package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/guutong/demo-gin/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/guutong/demo-gin/auth"
	"github.com/guutong/demo-gin/user"
)

// go build \
// -ldflags "-X main.buildcommit=`git rev-parse --short HEAD` -X main.buildtime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`" \
// -o main
var (
	buildcommit = "dev"
	buildtime   = time.Now().String()
)

// @title           Example API
// @version         1.0
// @description     This is a sample server celler server.
func main() {
	_, err := os.Create("/tmp/healthy")
	if err != nil {
		log.Fatal("can not create")
	}
	defer os.Remove("/tmp/healthy")

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Can not Load env!")
	}

	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Can not connect DB!")
	}

	r := gin.Default()

	r.GET("/health", Health())

	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"buildcommit": buildcommit,
			"buildtime":   buildtime,
		})
	})

	db.AutoMigrate(
		&user.User{},
	)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/tokenz", auth.GetToken(os.Getenv("WIFI_SECRET")))
	authRoute := r.Group("", auth.AuthMiddleware(os.Getenv("WIFI_SECRET")))

	userHandler := user.NewUserHandler(db)
	authRoute.POST("/users", userHandler.NewUser)
	authRoute.GET("/users", userHandler.GetUser)
	authRoute.DELETE("/users/:id", userHandler.DeleteUser)
	authRoute.PATCH("/users/:id", userHandler.UpdateUser)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGTERM)
	defer stop()

	server := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// go routine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("listen err %s", err.Error())
		}
	}()

	// - >
	// < - channel
	<-ctx.Done()
	stop()
	fmt.Println("Shuting down...")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server force shutdown...", err)
	}
}

// Health godoc
// @Summary      Health check
// @Description  Health check
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object} map[string]interface{}
// @Router       /health [get]
func Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok!",
		})
	}
}
