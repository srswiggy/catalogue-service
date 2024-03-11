package services

import (
	pb "catalogueservice/proto"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func setupMenuItemServiceTest() (sqlmock.Sqlmock, *MenuItemsService) {
	mockdb, mock, _ := sqlmock.New()
	dialector := postgres.New(
		postgres.Config{
			DriverName: "postgres",
			Conn:       mockdb,
		})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	service := MenuItemsService{Database: gormDB}

	return mock, &service
}

func TestMenuItemsCreatedSuccessfully(t *testing.T) {
	mock, service := setupMenuItemServiceTest()

	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "menu item 1")
	mock.ExpectQuery("INSERT INTO \"items\"").WillReturnRows(rows)
	mock.ExpectCommit()
	_, err := service.CreateMultiple(context.Background(), &pb.ItemsListRequest{})
	if err != nil {
		t.Fatalf(": %s", err)
	}
}

func TestMenuItemsCreateMultipleItemsSuccessfully(t *testing.T) {
	mock, service := setupMenuItemServiceTest()
	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "name", "price"}).AddRow(1, "menuitem1", 200).AddRow(2, "menuitem2", 300)
	mock.ExpectQuery("INSERT INTO \"items\"").WillReturnRows(rows)
	mock.ExpectCommit()
	_, err := service.CreateMultiple(context.Background(), &pb.ItemsListRequest{})
	if err != nil {
		t.Fatalf(": %s", err)
	}
}

func TestGetByRestaurtnIdReturnsSuccessfully(t *testing.T) {
	mock, service := setupMenuItemServiceTest()

	var restaurantID int64 = 7
	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(1, "Pizza", 10.99).
		AddRow(2, "Pasta", 8.99)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "menu_items" WHERE "restaurant_id" = $1 AND "menu_items"."deleted_at" IS NULL`)).
		WithArgs(restaurantID).
		WillReturnRows(rows)

	ctx := context.Background()
	req := &pb.RestaurantIdRequest{Id: restaurantID}
	response, err := service.GetByRestaurantId(ctx, req)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(response.Items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(response.Items))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
