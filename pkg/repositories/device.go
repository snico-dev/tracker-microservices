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
	err = collection.FindOne(context.TODO(), bson.M{"device_id": deviceID, "plugged": true}).Decode(&device)
	return device, err
}

//GetActiveDevice search for device in database
func (repo *deviceRepository) GetActiveDevice(id string) (models.Device, error) {
	collection, err := repo.ctx.GetCollection(models.Device{})
	device := models.Device{}
	err = collection.FindOne(context.TODO(), bson.M{"_id": id, "active": true}).Decode(&device)
	return device, err
}

//GetActiveDeviceByCode search for device in database
func (repo *deviceRepository) GetActiveDeviceByCode(id string) (models.Device, error) {
	collection, err := repo.ctx.GetCollection(models.Device{})
	device := models.Device{}
	err = collection.FindOne(context.TODO(), bson.M{"device_id": id, "active": true}).Decode(&device)
	return device, err
}

//GetActiveDeviceByPINCode search for device in database
func (repo *deviceRepository) GetActiveDeviceByPINCode(pinCode string) (models.Device, error) {
	collection, err := repo.ctx.GetCollection(models.Device{})
	device := models.Device{}
	err = collection.FindOne(context.TODO(), bson.M{"activation_pin_code": pinCode, "active": true}).Decode(&device)
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
	collection, err := repo.ctx.GetCollection(device)
	filter := bson.M{"_id": bson.M{"$eq": device.ID}}
	update := bson.D{{
		"$set",
		bson.D{
			{"update_at", device.UpdateAt},
			{"plugged", device.Plugged},
			{"active", device.Active},
			{"activation_pin_code", device.ActivationPINCode},
		},
	}}

	_, err = collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
