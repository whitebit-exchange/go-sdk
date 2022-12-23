package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/whitebit-exchange/go-sdk"
)

type Client struct {
	mock.Mock
}

func (c *Client) SendRequest(endpoint whitebit.Endpoint) ([]byte, error) {
	ret := c.Called(endpoint)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(whitebit.Endpoint) []byte); ok {
		r0 = rf(endpoint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(whitebit.Endpoint) error); ok {
		r1 = rf(endpoint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
