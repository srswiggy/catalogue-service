package services

import (
	"catalogueservice/models"
	pb "catalogueservice/proto"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type RestaurantService struct {
	pb.UnimplementedRestaurantServiceServer
	Database *gorm.DB
}

func (server *RestaurantService) Create(ctx context.Context, req *pb.RestaurantRequest) (*pb.RestaurantResponse, error) {
	restaurant := models.Restaurant{
		Name: req.GetName(),
	}

	server.Database.Create(&restaurant)

	return &pb.RestaurantResponse{
		Id:   restaurant.ID,
		Name: restaurant.Name,
	}, nil
}

func (server *RestaurantService) Get(ctx context.Context, req *pb.RestaurantIdRequest) (*pb.RestaurantResponse, error) {

	restaurant := models.Restaurant{}

	server.Database.First(&restaurant, req.GetId())

	return &pb.RestaurantResponse{
		Id:   restaurant.ID,
		Name: restaurant.Name,
	}, nil
}

func (server *RestaurantService) GetAll(ctx context.Context, req *pb.NoParams) (*pb.RestaurantListResponse, error) {
	var restaurants []models.Restaurant

	server.Database.Find(&restaurants)

	var restaurantResponseList []*pb.RestaurantResponse

	for _, restaurant := range restaurants {
		restaurantResponseList = append(restaurantResponseList, &pb.RestaurantResponse{
			Id:   restaurant.ID,
			Name: restaurant.Name,
		})
	}

	return &pb.RestaurantListResponse{Restaurants: restaurantResponseList}, nil
}
