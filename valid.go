package algorithm

import (
	"reflect"
	"testing"
)

type Valid interface {
	Valid() bool
}

type Validator struct {
	Want interface{}
	Rlt  interface{}
	T    *testing.T
}

func New(rlt, want interface{}, t *testing.T) *Validator {
	return &Validator{
		Want: want,
		Rlt:  rlt,
		T:    t,
	}
}

func (v *Validator) Valid() bool {
	if !reflect.DeepEqual(v.Want, v.Rlt) {
		v.T.Errorf("rlt: %v not want: %v", v.Rlt, v.Want)
		return false
	}
	return true
}
