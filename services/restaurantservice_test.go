package services

import (
	"catalogueservice/models"
	pb "catalogueservice/proto"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func setupRestaurantServiceTest() (sqlmock.Sqlmock, *RestaurantService) {
	mockdb, mock, _ := sqlmock.New()
	dialector := postgres.New(
		postgres.Config{
			DriverName: "postgres",
			Conn:       mockdb,
		})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	service := RestaurantService{Database: gormDB}

	return mock, &service
}

func TestCreateRestaurant(t *testing.T) {
	mock, service := setupRestaurantServiceTest()

	req := &pb.RestaurantRequest{
		Name: "Test Restaurant",
		Location: &pb.Location{
			Longitude: 10.0,
			Latitude:  20.0,
		},
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "restaurants"`)).
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			req.GetName(),
			req.Location.GetLatitude(),
			req.Location.GetLongitude(),
		).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	ctx := context.Background()
	resp, err := service.Create(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Test Restaurant", resp.Name)
	assert.Equal(t, float32(10.0), resp.Location.Longitude)
	assert.Equal(t, float32(20.0), resp.Location.Latitude)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetRestaurant(t *testing.T) {
	mock, service := setupRestaurantServiceTest()

	var restaurantID int64 = 1
	expectedRestaurant := models.Restaurant{
		Model: gorm.Model{ID: 1},
		Name:  "Test Restaurant",
		Location: models.Location{
			Latitude:  20.0,
			Longitude: 10.0,
		},
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "restaurants" WHERE "restaurants"."id" = $1 AND "restaurants"."deleted_at" IS NULL ORDER BY "restaurants"."id" LIMIT $2`)).
		WithArgs(restaurantID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "latitude", "longitude"}).
			AddRow(expectedRestaurant.ID, expectedRestaurant.Name, expectedRestaurant.Location.Latitude, expectedRestaurant.Location.Longitude))

	req := &pb.RestaurantIdRequest{Id: restaurantID}

	ctx := context.Background()
	resp, err := service.Get(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expectedRestaurant.Name, resp.Name)
	assert.Equal(t, expectedRestaurant.Location.Latitude, resp.Location.Latitude)
	assert.Equal(t, expectedRestaurant.Location.Longitude, resp.Location.Longitude)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetAllRestaurants(t *testing.T) {
	mock, service := setupRestaurantServiceTest()

	restaurants := []models.Restaurant{
		{
			Model: gorm.Model{ID: 1},
			Name:  "Test Restaurant 1",
			Location: models.Location{
				Latitude:  20.0,
				Longitude: 10.0,
			},
		},
		{
			Model: gorm.Model{ID: 2},
			Name:  "Test Restaurant 2",
			Location: models.Location{
				Latitude:  30.0,
				Longitude: 20.0,
			},
		},
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "restaurants"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "latitude", "longitude"}).
			AddRow(restaurants[0].ID, restaurants[0].Name, restaurants[0].Location.Latitude, restaurants[0].Location.Longitude).
			AddRow(restaurants[1].ID, restaurants[1].Name, restaurants[1].Location.Latitude, restaurants[1].Location.Longitude))

	ctx := context.Background()
	req := &pb.NoParams{}
	resp, err := service.GetAll(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Restaurants, 2)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
