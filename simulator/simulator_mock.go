package simulator

import "github.com/stretchr/testify/mock"

// builder
func NewCatchSimulatorMockTestify() *CatchSimulatorMockTestify {
	return &CatchSimulatorMockTestify{}
}

type CatchSimulatorMockTestify struct {
	//  provides a system by which it is possible to mock your objects and verify calls are happening as expected
	// advantage of the library: it has a method you don't have to do the spy as in the manual way
	mock.Mock
}

// the "Testify" mocking library is being used to create a CatchSimulator mock object (CatchSimulatorMockTestify) 
//that implements the same interface as the actual CatchSimulator class
func (cm *CatchSimulatorMockTestify) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	args := cm.Called(distance, speed, catchSpeed)

	// indicates that the return value of the first mock method call at position 0 of the registered return values is being retrieved
	r0 := args.Bool(0)

	return r0
}

func (cm *CatchSimulatorMockTestify) GetLinearDistance(position [2]float64) float64 {
	args := cm.Called(position)

	r0 := args.Get(0).(float64)

	return r0
}

