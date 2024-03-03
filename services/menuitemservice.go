package services

import (
	"catalogueservice/models"
	pb "catalogueservice/proto"
	"context"
	"gorm.io/gorm"
)

type MenuItemsService struct {
	pb.UnimplementedMenuItemsServiceServer
	Database *gorm.DB
}

func (server *MenuItemsService) CreateMultiple(ctx context.Context, req *pb.ItemsListRequest) (*pb.ItemsListResponse, error) {
	var menuItems []*models.MenuItem

	for _, itemReq := range req.Items {
		menuItems = append(menuItems, &models.MenuItem{Name: itemReq.Name, RestaurantID: req.RestaurantId})
	}

	server.Database.Create(menuItems)

	var itemsResponseList []*pb.ItemResponse

	for _, item := range menuItems {
		itemsResponseList = append(itemsResponseList, &pb.ItemResponse{Id: item.ID, Name: item.Name})
	}

	return &pb.ItemsListResponse{
		Items: itemsResponseList,
	}, nil
}

func (server *MenuItemsService) GetByRestaurantId(ctx context.Context, req *pb.RestaurantIdRequest) (*pb.ItemsListResponse, error) {
	var menuItems []*models.MenuItem
	server.Database.Where("restaurant_id", req.Id).Find(&menuItems)

	var itemsResponseList []*pb.ItemResponse

	for _, item := range menuItems {
		itemsResponseList = append(itemsResponseList, &pb.ItemResponse{Id: item.ID, Name: item.Name})
	}

	return &pb.ItemsListResponse{
		Items: itemsResponseList,
	}, nil
}
