package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type Motorbike interface {
	GetMotorbikeType() int
}

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}