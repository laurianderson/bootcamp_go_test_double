package prey

// Stub o Dummy: is created to temporarily replace a part of the code that is not yet implemented

// builder
func NewPreyStub() *PreyStub {
	return &PreyStub{}
}

// PreyStub is a struc that contains the method GetSpeedFn
type PreyStub struct {
	GetSpeedFn func() float64
}

// the function calls the method of the interface, but returns the method of the struct stub
func (ps *PreyStub) GetSpeed() float64 {
	return ps.GetSpeedFn()
}