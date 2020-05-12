package dtos

import "time"

//DeviceDTO device dto
type DeviceDTO struct {
	ID       string    `bson:"_id" json:"id"`
	DeviceID string    `bson:"device_id"`
	CreateAt time.Time `bson:"create_at"`
	Plugged  bool      `bson:"plugged"`
	Active   bool      `bson:"active"`
	UserID   string    `bson:"user_id"`
}
