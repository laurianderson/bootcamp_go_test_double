package storage

import "github.com/stretchr/testify/mock"


type StorageMock struct {
	mock.Mock
}

func (s *StorageMock) GetValue(key string) interface{} {
	args := s.Called(key)
    return args.Get(0)
}