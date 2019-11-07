package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	defaultLimit = 100
)

type Options struct {
	Limit    int64
	CursorID primitive.ObjectID
}
