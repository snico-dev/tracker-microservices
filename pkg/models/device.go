package models

import (
	"crypto/rand"
	"io"
	"time"

	"github.com/NicolasDeveloper/tracker-microservices/pkg/timeconvert"
	"github.com/beevik/guid"
)

//Device model
type Device struct {
	ID                string    `bson:"_id" json:"id"`
	DeviceID          string    `bson:"device_id"`
	UpdateAt          time.Time `bson:"update_at"`
	CreateAt          time.Time `bson:"create_at"`
	Plugged           bool      `bson:"plugged"`
	Active            bool      `bson:"active"`
	UserID            string    `bson:"user_id"`
	ActivationPINCode string    `bson:"activation_pin_code"`
}

//NewDevice constructor
func NewDevice(deviceID string, userID string) (Device, error) {
	guidid := guid.New()
	return Device{
		ID:                guidid.String(),
		CreateAt:          time.Now().In(timeconvert.GetLocation()),
		UpdateAt:          time.Now().In(timeconvert.GetLocation()),
		DeviceID:          deviceID,
		Plugged:           false,
		Active:            true,
		UserID:            userID,
		ActivationPINCode: encodeToString(6),
	}, nil
}

//Plug do login
func (device *Device) Plug() error {
	device.UpdateAt = time.Now().In(timeconvert.GetLocation())
	device.Plugged = true
	return nil
}

//UnPlug do logout
func (device *Device) UnPlug() error {
	device.UpdateAt = time.Now().In(timeconvert.GetLocation())
	device.Plugged = false
	return nil
}

//Inactive inactive device in database
func (device *Device) Inactive() {
	device.UpdateAt = time.Now().In(timeconvert.GetLocation())
	device.Active = false
}

func encodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
