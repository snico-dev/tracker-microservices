package repositories

import (
	"context"
	"log"

	"github.com/NicolasDeveloper/tracker-microservices/pkg/database/dbcontext"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

type deviceRepository struct {
	ctx dbcontext.DbContext
}

//NewDeviceRepository constructor
func NewDeviceRepository(ctx dbcontext.DbContext) (IDeviceRepository, error) {
	return &deviceRepository{
		ctx: ctx,
	}, nil
}

//GetLoggedDevice get logged device
func (repo *deviceRepository) GetLoggedDevice(deviceID string) (models.Device, error) {
	collection, err := repo.ctx.GetCollection(models.Device{})
	device := models.Device{}
	err = collection.FindOne(context.TODO(), bson.M{"device_id": deviceID, "logged": true}).Decode(&device)
	return device, err
}

//GetActiveDevice search for device in database
func (repo *deviceRepository) GetActiveDevice(deviceID string) (models.Device, error) {
	collection, err := repo.ctx.GetCollection(models.Device{})
	device := models.Device{}
	err = collection.FindOne(context.TODO(), bson.M{"device_id": deviceID, "active": true}).Decode(&device)
	return device, err
}

//CreateDevice search for device in database
func (repo *deviceRepository) CreateDevice(device models.Device) error {
	collection, error := repo.ctx.GetCollection(device)

	_, err := collection.InsertOne(context.TODO(), device)
	if err != nil {
		log.Fatal(err)
	}

	return error
}

//Update update method
func (repo *deviceRepository) Update(device models.Device) error {
	collection, error := repo.ctx.GetCollection(device)
	filter := bson.M{"_id": bson.M{"$eq": device.ID}}
	_, err := collection.UpdateOne(context.TODO(), filter, device)
	if err != nil {
		log.Fatal(err)
	}
	return error
}
