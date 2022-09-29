package withinterfaces

import "math"

const (
	MAX_DELIVERY_DISTANCE = 100
)

type Calculator interface {
	Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64
}

type calculator struct{}

func NewCalculator() Calculator {
	return &calculator{}
}

type Deliverable interface {
	WithinDeliveryDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, c Calculator) bool
}

type deliverabler struct{}

func NewDeliverable() Deliverable {
	return &deliverabler{}
}

// fonte: https://gist.github.com/hotdang-ca/6c1ee75c48e515aec5bc6db6e3265e49
func (c *calculator) Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

func (d *deliverabler) WithinDeliveryDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, c Calculator) bool {
	distance := c.Distance(lat1, lng1, lat2, lng2, "N")

	return distance < MAX_DELIVERY_DISTANCE
}
