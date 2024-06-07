package repositories 
import (
	"context"
	"time"
	 "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "golang-chi-mongodb/models"
)
type UserRepository struct {
	collection *mongo.Collection 

}
func NewUserRepository(db *mongo.Database)* UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}

}

func (u *UserRepository) CreateUser(ctx context.Context,user *models.User) error {
	ctx,cancel :=context.WithTimeout(ctx,10*time.Second)
    defer cancel()

	_,err := u.collection.InsertOne(ctx,user); 
	return err; 

 }

func (u *UserRepository) GetUser(ctx context.Context, id string) (*models.User, error) {
	
}
func (u *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {}
{}
func (u *UserRepository) DeleteUser(ctx context.Context, id string) error {}
{}
