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
	id          string
	name        string
	description string
	location    *LatLong
}

func NewShop(
	id string,
	name string,
	description string,
	latitude float64,
	longitude float64,
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

	location, err := NewLatLong(latitude, longitude)
	if err != nil {
		return nil, err
	}

	return HydrateShop(
		id,
		name,
		description,
		location.Latitude(),
		location.Longitude(),
	), nil
}

func HydrateShop(
	id string,
	name string,
	description string,
	latitude float64,
	longitude float64,
) *Shop {
	return &Shop{
		id:          id,
		name:        name,
		description: description,
		location:    HydrateLatLong(latitude, longitude),
	}
}

func (s *Shop) ID() string {
	return s.id
}

func (s *Shop) Name() string {
	return s.name
}

func (s *Shop) Description() string {
	return s.description
}

func (s *Shop) Latitude() float64 {
	return s.location.Latitude()
}

func (s *Shop) Longitude() float64 {
	return s.location.Longitude()
}
