package repositories

import "github.com/NicolasDeveloper/tracker-microservices/internal/trip/models"

//ITripRepository repository
type ITripRepository interface {
	Save(trip models.Trip) error
	GetOpenTrip(userID string) (models.Trip, error)
	GetTripsByUser(userID string) ([]models.Trip, error)
	CloseTrip(trip models.Trip, tracks []models.Track) error
	UpdateTracks(trip models.Trip, tracks []models.Track) error
}
