package shark

import (
	"fmt"
	"testdoubles/pkg/storage"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"
	"github.com/stretchr/testify/assert"
)
/*
func TestWhiteShark_Hunt(t *testing.T) {
	MODIFY DO WITH THE FUNCTIONALITIES OF THE MOCK PACKAGE
}
*/
func TestWhiteSharkHuntFailedSpeed(t *testing.T) {
	// arrange
	values := map[string]interface{}{
		"white_shark_speed": 5.0,
		"white_shark_x":     0.0,
		"white_shark_y":     1.0,
		"tuna_speed":        10.0,
	}
	storage := storage.NewMockStorage(values)
	whiteSharkInstance := whiteShark{
		speed:     5.0,
		position:  [2]float64{1.0, 2.0},
		simulator: simulator.NewCatchSimulator(2.0),
	}
	preyInstance := prey.CreateTuna(storage)

	// act
	result := whiteSharkInstance.Hunt(preyInstance)

	// assert
	assert.Equal(t, fmt.Errorf("could not hunt the prey"), result)
}