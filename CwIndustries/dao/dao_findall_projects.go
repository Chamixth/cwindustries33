package dao

import (
	"CwIndustries/dbConfig"
	"CwIndustries/dto"
	"go.mongodb.org/mongo-driver/bson"
    "context"
    "errors"
)

func DB_FindallProjects () (*[]dto.Projects, error) {
	var objects []dto.Projects
	results, err := dbConfig.DATABASE.Collection("Projectss").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("Error When Fetching Projects")
	}
	for results.Next(context.Background()) {
		var object dto.Projects
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Projects")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
