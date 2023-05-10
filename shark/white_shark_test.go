package shark

import (
	"testdoubles/pkg/storage"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"
	// "github.com/stretchr/testify/assert"
)

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