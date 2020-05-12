package repositories

import (
	"context"
	"log"

	"github.com/NicolasDeveloper/tracker-microservices/internal/trip/models"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/database/dbcontext"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//tripReposiory struct type
type tripReposiory struct {
	ctx dbcontext.DbContext
}

//NewTripRepository constructor
func NewTripRepository(ctx dbcontext.DbContext) (ITripRepository, error) {
	return &tripReposiory{
		ctx: ctx,
	}, nil
}

//Save method to insert trip in database
func (repo *tripReposiory) Save(trip models.Trip) error {
	collection, error := repo.ctx.GetCollection(trip)

	_, err := collection.InsertOne(context.TODO(), trip)
	if err != nil {
		log.Fatal(err)
	}

	return error
}

//Update save
func (repo *tripReposiory) CloseTrip(trip models.Trip, tracks []models.Track) error {
	collection, err := repo.ctx.GetCollection(trip)

	filter := bson.M{"_id": bson.M{"$eq": trip.ID}}
	update := bson.D{{
		"$set",
		bson.D{
			{"total_gps_mileage", trip.TotalGpsMileage},
			{"total_mileage", trip.TotalMileage},
			{"current_coordinate", trip.CurrentCoordinate},
			{"update_at", trip.UpdateAt},
			{"max_speed", trip.MaxSpeed},
			{"max_rpm", trip.MaxRPM},
			{"closed_at", trip.ClosedAt},
			{"end_track", trip.EndTrack},
			{"open", trip.Open},
		},
	}, {
		"$push",
		bson.M{"tracks": bson.M{"$each": tracks}},
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

//UpdateTracks save
func (repo *tripReposiory) UpdateTracks(trip models.Trip, tracks []models.Track) error {
	collection, err := repo.ctx.GetCollection(trip)

	filter := bson.M{"_id": bson.M{"$eq": trip.ID}}
	update := bson.D{{
		"$set",
		bson.D{
			{"total_gps_mileage", trip.TotalGpsMileage},
			{"total_mileage", trip.TotalMileage},
			{"current_coordinate", trip.CurrentCoordinate},
			{"update_at", trip.UpdateAt},
			{"max_speed", trip.MaxSpeed},
			{"max_rpm", trip.MaxRPM},
		},
	}, {
		"$push",
		bson.M{"tracks": bson.M{"$each": tracks}},
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

//GetOpenTrip save
func (repo *tripReposiory) GetOpenTrip(userID string) (models.Trip, error) {
	collection, err := repo.ctx.GetCollection(&models.Trip{})

	matchStage := bson.D{{"$match", bson.D{{"user_id", userID}, {"open", true}}}}
	setStage := bson.D{{"$set", bson.M{"tracks": bson.M{"$slice": []interface{}{"$tracks", -1}}}}}

	ctx := context.TODO()
	cursor, err := collection.Aggregate(
		ctx,
		mongo.Pipeline{matchStage, setStage},
	)
	defer cursor.Close(ctx)

	var trip models.Trip

	if cursor.Next(ctx) {
		err = cursor.Decode(&trip)
	}

	return trip, err
}

//GetTripsByUser save
func (repo *tripReposiory) GetTripsByUser(userID string) ([]models.Trip, error) {
	collection, err := repo.ctx.GetCollection(&models.Trip{})

	matchStage := bson.D{{"$match", bson.D{{"user_id", userID}}}}
	projectStage := bson.D{{"$project", bson.D{{"tracks", 0}}}}

	ctx := context.TODO()
	cursor, err := collection.Aggregate(
		ctx,
		mongo.Pipeline{matchStage, projectStage},
	)
	defer cursor.Close(ctx)

	var results []models.Trip
	err = cursor.All(ctx, &results)

	return results, err
}
