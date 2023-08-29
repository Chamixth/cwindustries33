package dao

import (
	"CwIndustries/dbConfig"
	"CwIndustries/dto"
	"go.mongodb.org/mongo-driver/bson"
    "context"
    "errors"
)

func DB_FindallEngineers () (*[]dto.Engineers, error) {
	var objects []dto.Engineers
	results, err := dbConfig.DATABASE.Collection("Engineerss").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("Error When Fetching Engineers")
	}
	for results.Next(context.Background()) {
		var object dto.Engineers
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Engineers")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
