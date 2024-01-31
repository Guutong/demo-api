package main

import (
	"log"
	"net"
	"os"

	"github.com/guutong/demo-gin/user"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	pb "github.com/guutong/demo-gin/proto"
)

// func main() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("Can not Load env!")
// 	}

// 	dsn := os.Getenv("DATABASE")
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Printf("Can not connect DB!")
// 	}

// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong!",
// 		})
// 	})

// 	db.AutoMigrate(
// 		&user.User{},
// 	)

// 	r.GET("/tokenz", auth.GetToken(os.Getenv("WIFI_SECRET")))
// 	authRoute := r.Group("", auth.AuthMiddleware(os.Getenv("WIFI_SECRET")))

// 	userRepository := user.NewUserRepository(db)
// 	userHandler := user.NewUserHandler(userRepository)
// 	authRoute.POST("/users", router.NewGinHandler(userHandler.NewUser))
// 	authRoute.GET("/users", router.NewGinHandler(userHandler.GetUser))
// 	authRoute.DELETE("/users/:id", router.NewGinHandler(userHandler.DeleteUser))
// 	authRoute.PATCH("/users/:id", router.NewGinHandler(userHandler.UpdateUser))

// 	r.Run()
// }

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

	lis, err := net.Listen("tcp", "[::1]:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userRepository := user.NewUserRepository(db)
	grpcServer := grpc.NewServer()
	service := user.NewUserServiceServer(userRepository)

	pb.RegisterUserServiceServer(grpcServer, service)
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("Error strating server: %v", err)
	}
}
