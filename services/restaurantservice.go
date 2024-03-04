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
	location := models.Location{
		Longitude: req.Location.Longitude,
		Latitude:  req.Location.Latitude,
	}
	restaurant := models.Restaurant{
		Name:     req.GetName(),
		Location: location,
	}

	server.Database.Create(&restaurant)

	return &pb.RestaurantResponse{
		Id:   restaurant.ID,
		Name: restaurant.Name,
		Location: &pb.Location{
			Longitude: restaurant.Location.Longitude,
			Latitude:  restaurant.Location.Latitude,
		},
	}, nil
}

func (server *RestaurantService) Get(ctx context.Context, req *pb.RestaurantIdRequest) (*pb.RestaurantResponse, error) {

	restaurant := models.Restaurant{}

	server.Database.First(&restaurant, req.GetId())

	return &pb.RestaurantResponse{
		Id:   restaurant.ID,
		Name: restaurant.Name,
		Location: &pb.Location{
			Latitude:  restaurant.Location.Latitude,
			Longitude: restaurant.Location.Longitude,
		},
	}, nil
}

func (server *RestaurantService) GetAll(ctx context.Context, req *pb.NoParams) (*pb.RestaurantListResponse, error) {
	var restaurants []models.Restaurant

	server.Database.Find(&restaurants)

	var restaurantResponseList []*pb.RestaurantResponse

	for _, restaurant := range restaurants {
		location := &pb.Location{
			Latitude:  restaurant.Latitude,
			Longitude: restaurant.Longitude,
		}
		restaurantResponseList = append(restaurantResponseList, &pb.RestaurantResponse{
			Id:       restaurant.ID,
			Name:     restaurant.Name,
			Location: location,
		})
	}

	return &pb.RestaurantListResponse{Restaurants: restaurantResponseList}, nil
}
