package dao

import (
	"CwIndustries/dbConfig"
    "CwIndustries/dto"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_UpdateEngineers (object *dto.Engineers)  error {

	result, err := dbConfig.DATABASE.Collection("Engineerss").UpdateOne(context.Background(), bson.M{"objectid": object.ObjectId}, bson.M{"$set": object})
	if err != nil {
		return err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return errors.New("Specified ID not found!")
	}

	return nil
}