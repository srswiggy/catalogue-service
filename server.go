package main

import (
	"catalogueservice/models"
	pb "catalogueservice/proto"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net"
)

const (
	port = ":9000"
)

func databaseConn() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=catalogueservice port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("error connecting to database: %s", err)
	}

	err = db.AutoMigrate(&models.Restaurant{})
	if err != nil {
		return nil
	}

	db.Logger.LogMode(logger.Info)

	return db
}

func main() {
	db := databaseConn()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error listening to port: %s", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterRestaurantServiceServer(grpcServer, &restaurantService{database: db})

	log.Println("Listening to port", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to Server: %s", err)
	}
}
