package domain

import (
	"errors"
	"strings"
)

var (
	ErrEmptyID          = errors.New("empty id")
	ErrEmptyName        = errors.New("empty name")
	ErrEmptyDescription = errors.New("empty description")
)

type Shop struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Location    LatLong `json:"location"`
}

func NewShop(
	id string,
	name string,
	description string,
	latitude int,
	longitude int,
) (*Shop, error) {
	if strings.TrimSpace(id) == "" {
		return nil, ErrEmptyID
	}

	if strings.TrimSpace(name) == "" {
		return nil, ErrEmptyName
	}

	if strings.TrimSpace(description) == "" {
		return nil, ErrEmptyDescription
	}

	return &Shop{
		ID:          id,
		Name:        name,
		Description: description,
		Location: LatLong{
			Latitude:  latitude,
			Longitude: longitude,
		},
	}, nil
}
