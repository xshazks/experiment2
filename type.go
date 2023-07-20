package gosaw

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Monitor struct {
	ID         						primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Proker 		string             `bson:"proker,omitempty" json:"proker,omitempty"`
	Status	 	string             `bson:"status,omitempty" json:"status,omitempty"`
	About	 	string             `bson:"about,omitempty" json:"about,omitempty"`
	Karyawan 	string             `bson:"karyawan,omitempty" json:"karyawan,omitempty"`
}