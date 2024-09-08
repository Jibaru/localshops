package domain

import "errors"

var (
	ErrLatitudeOutOfRange  = errors.New("latitude out of range")
	ErrLongitudeOutOfRange = errors.New("longitude out of range")
)

type LatLong struct {
	latitude  float64
	longitude float64
}

func NewLatLong(
	latitude float64,
	longitude float64,
) (*LatLong, error) {
	if !(latitude >= -90 && latitude <= 90) {
		return nil, ErrLatitudeOutOfRange
	}

	if !(longitude >= -180 && longitude <= 180) {
		return nil, ErrLongitudeOutOfRange
	}

	return HydrateLatLong(latitude, longitude), nil
}

func HydrateLatLong(
	latitude float64,
	longitude float64,
) *LatLong {
	return &LatLong{
		latitude:  latitude,
		longitude: longitude,
	}
}

func (l *LatLong) Latitude() float64 {
	return l.latitude
}

func (l *LatLong) Longitude() float64 {
	return l.longitude
}
