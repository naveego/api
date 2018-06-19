package utils

import "github.com/satori/go.uuid"

func NewGUID() (uuid.UUID, error) {
	return uuid.NewV4()
}

func ParseGUID(id string) (uuid.UUID, error) {
	u, err := uuid.FromString(id)
	return u, err
}
