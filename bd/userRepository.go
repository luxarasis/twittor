package bd

import (
	"context"
	"time"

	"github.com/luxarasis/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func SaveUser(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("user")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := col.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	ID, _ := result.InsertedID.(primitive.ObjectID)
	return ID.String(), true, nil
}

func FindUserByEmail(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("user")

	query := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, query).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}

func FindUserByID(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("user")

	var result models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	query := bson.M{"_id": objID}

	err := col.FindOne(ctx, query).Decode(&result)
	result.Password = ""
	if err != nil {
		return result, err
	}

	return result, nil
}

/* Login realiza el chequeo de login en la base de datos */
func Login(email string, pass string) (models.User, bool) {
	user, found, _ := FindUserByEmail(email)

	if !found {
		return user, false
	}

	passBytes := []byte(pass)
	passDBBytes := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passDBBytes, passBytes)
	if err != nil {
		return user, false
	}

	return user, true
}
