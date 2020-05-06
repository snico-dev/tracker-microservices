package services

import (
	"github.com/NicolasDeveloper/tracker-microservices/internal/trip/models"
	"github.com/stretchr/testify/mock"
)

//TripRepositorySTUB struct type
type TripRepositorySTUB struct {
	mock.Mock
}

//NewTripRepositoryStub constructor
func NewTripRepositoryStub() (TripRepositorySTUB, error) {
	return TripRepositorySTUB{}, nil
}

//Save save
func (repo *TripRepositorySTUB) Save(trip models.Trip) error {
	args := repo.Called()
	return args.Error(0)
}

//Update save
func (repo *TripRepositorySTUB) Update(trip models.Trip) error {
	args := repo.Called()
	return args.Error(0)
}

//GetTripsByUser save
func (repo *TripRepositorySTUB) GetTripsByUser(userID string) ([]models.Trip, error) {
	args := repo.Called(userID)
	return args.Get(0).([]models.Trip), args.Error(1)
}

//GetOpenTrip save
func (repo *TripRepositorySTUB) GetOpenTrip(userID string) (models.Trip, error) {
	args := repo.Called(userID)
	return args.Get(0).(models.Trip), args.Error(1)
}

//UpdateTracks update
func (repo *TripRepositorySTUB) UpdateTracks(trip models.Trip, tracks []models.Track) error {
	return nil
}

//CloseTrip close
func (repo *TripRepositorySTUB) CloseTrip(trip models.Trip, tracks []models.Track) error {
	return nil
}
