package services

import "github.com/stretchr/testify/mock"

//MapBoxACLSub type
type MapBoxACLSub struct {
	mock.Mock
}

//NewMapBoxStub constructor
func NewMapBoxStub() (MapBoxACLSub, error) {
	return MapBoxACLSub{}, nil
}

//GetAddressName return address name
func (mp *MapBoxACLSub) GetAddressName(latitude float64, longitude float64) (string, error) {
	args := mp.Called(latitude, longitude)
	return args.String(0), args.Error(1)
}
