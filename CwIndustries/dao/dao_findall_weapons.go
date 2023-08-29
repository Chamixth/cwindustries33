package dao

import (
	"CwIndustries/dbConfig"
	"CwIndustries/dto"
	"go.mongodb.org/mongo-driver/bson"
    "context"
    "errors"
)

func DB_FindallWeapons () (*[]dto.Weapons, error) {
	var objects []dto.Weapons
	results, err := dbConfig.DATABASE.Collection("Weaponss").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("Error When Fetching Weapons")
	}
	for results.Next(context.Background()) {
		var object dto.Weapons
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Weapons")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
