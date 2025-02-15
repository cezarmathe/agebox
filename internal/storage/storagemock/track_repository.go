// Code generated by mockery v2.5.1. DO NOT EDIT.

package storagemock

import (
	context "context"

	model "github.com/slok/agebox/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// TrackRepository is an autogenerated mock type for the TrackRepository type
type TrackRepository struct {
	mock.Mock
}

// GetSecretRegistry provides a mock function with given fields: ctx
func (_m *TrackRepository) GetSecretRegistry(ctx context.Context) (*model.SecretRegistry, error) {
	ret := _m.Called(ctx)

	var r0 *model.SecretRegistry
	if rf, ok := ret.Get(0).(func(context.Context) *model.SecretRegistry); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SecretRegistry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveSecretRegistry provides a mock function with given fields: ctx, reg
func (_m *TrackRepository) SaveSecretRegistry(ctx context.Context, reg model.SecretRegistry) error {
	ret := _m.Called(ctx, reg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SecretRegistry) error); ok {
		r0 = rf(ctx, reg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
