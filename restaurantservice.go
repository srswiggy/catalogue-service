package main

import (
	"catalogueservice/models"
	pb "catalogueservice/proto"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type restaurantService struct {
	pb.UnimplementedRestaurantServiceServer
	database *gorm.DB
}

func (server *restaurantService) Create(ctx context.Context, req *pb.RestaurantRequest) (*pb.RestaurantResponse, error) {
	restaurant := models.Restaurant{
		Name: req.GetName(),
	}

	server.database.Create(&restaurant)

	return &pb.RestaurantResponse{
		Id:   restaurant.ID,
		Name: restaurant.Name,
	}, nil
}

func (server *restaurantService) Get(ctx context.Context, req *pb.RestaurantIdRequest) (*pb.RestaurantResponse, error) {

	restaurant := models.Restaurant{}

	server.database.First(&restaurant, req.GetId())

	return &pb.RestaurantResponse{
		Id:   restaurant.ID,
		Name: restaurant.Name,
	}, nil
}

func (server *restaurantService) GetAll(ctx context.Context, req *pb.NoParams) (*pb.RestaurantListResponse, error) {
	var restaurants []models.Restaurant

	server.database.Find(&restaurants)

	var restaurantResponseList []*pb.RestaurantResponse

	for _, restaurant := range restaurants {
		restaurantResponseList = append(restaurantResponseList, &pb.RestaurantResponse{
			Id:   restaurant.ID,
			Name: restaurant.Name,
		})
	}

	return &pb.RestaurantListResponse{Restaurants: restaurantResponseList}, nil
}
