package shark

import (
	"fmt"
	"testdoubles/pkg/storage"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
)
/*
func TestWhiteShark_Hunt(t *testing.T) {
	// arrange
	storage := storage.NewStorageMock()
	pr := prey.NewPreyStub()
	pr.GetSpeedFn = func() float64 {
 		return 0
 	}
	sm := simulator.NewCatchSimulatorMockTestify()
	sm.On("CanCatch", 0.0, 0.0, 0.0).Return(true)
 	sm.On("GetLinearDistance", [2]float64{1,2}).Return(0.0)

 	sh := CreateWhiteShark(sm)

 	// act
	sh.Hunt(pr)

	// assert
	//sm.AssertCalled(t, "CanCatch", 0.0, 0.0, 0.0)
	//sm.AssertCalled(t, "GetLinearDistance", [2]float64{1,2})
	sm.AssertCalled(t, "Hunt")
	// expected called
	//assert.True(t, sm.CanCatch())
 
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