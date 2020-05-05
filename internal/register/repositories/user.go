package repositories

import (
	"context"

	"github.com/NicolasDeveloper/tracker-microservices/internal/register/models"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/database/dbcontext"
	"go.mongodb.org/mongo-driver/bson"
)

type userRepository struct {
	ctx dbcontext.DbContext
}

//NewUserRepository constructor
func NewUserRepository(ctx dbcontext.DbContext) (IUserRepository, error) {
	return &userRepository{
		ctx: ctx,
	}, nil
}

func (repo *userRepository) GetByPINCode(pinCode string) (models.User, error) {
	collection, err := repo.ctx.GetCollection(models.User{})
	user := models.User{}
	err = collection.FindOne(context.TODO(), bson.M{"pin_code": pinCode}).Decode(&user)
	return user, err
}

func (repo *userRepository) GetByID(id string) (models.User, error) {
	collection, err := repo.ctx.GetCollection(models.User{})
	user := models.User{}
	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	return user, err
}

func (repo *userRepository) Create(user models.User) error {
	collection, err := repo.ctx.GetCollection(user)
	_, err = collection.InsertOne(context.TODO(), user)
	return err
}
