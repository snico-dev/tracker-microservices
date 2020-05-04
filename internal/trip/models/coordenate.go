package models

//Coordenate model type
type Coordenate struct {
	Latitude  float64 `bson:"latitude"`
	Longitude float64 `bson:"longitude"`
}
