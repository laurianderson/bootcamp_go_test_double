package prey

import (
	"math/rand"
	"testdoubles/pkg/storage"
)

type tuna struct {
	// max speed in m/s
	maxSpeed  float64
}

func (t *tuna) GetSpeed() float64 {
	return t.maxSpeed * rand.Float64()
}

func CreateTuna(s storage.Storage) Prey {
	return &tuna{
		maxSpeed:  252,
	}
}
