package dao

import (
	"CwIndustries/dbConfig"
    "CwIndustries/dto"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_UpdateProjects (object *dto.Projects)  error {

	result, err := dbConfig.DATABASE.Collection("Projectss").UpdateOne(context.Background(), bson.M{"objectid": object.ObjectId}, bson.M{"$set": object})
	if err != nil {
		return err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return errors.New("Specified ID not found!")
	}

	return nil
}