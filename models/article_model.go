package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Title       string             `json:"title,omitempty" validate:"required"`
	Content     string             `json:"content,omitempty" validate:"required"`
	Author      string             `json:"author,omitempty" validate:"requiired"`
	CreatedDate time.Time          `json:"createDate"`
}
