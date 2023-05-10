package storage

// In the mock I abstract all the logic of reading a file find the value, do the unmarshall
// directly create a struct that contains a map key value
type storageMock struct {
	Values map[string]interface{}
}

// GetValue evaluates if the key matches the map key value and returns that key
func (s *storageMock) GetValue(key string) interface{} {
	if v, ok := s.Values[key]; ok {
		return v
	}
	return nil
}


// builder
func NewMockStorage(values map[string]interface{}) Storage {
	return &storageMock{Values: values}
}