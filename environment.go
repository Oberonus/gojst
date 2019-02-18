package gojst

import (
	"github.com/robertkrimen/otto"
)

type environment struct {
	vm *otto.Otto
	D  interface{}
}

func (e *environment) C(name string, params ...interface{}) (interface{}, error) {
	ottoVal, err := e.vm.Call(name, nil, params...)
	if err != nil {
		return nil, err
	}
	goVal, err := ottoVal.Export()
	if err != nil {
		return nil, err
	}
	return goVal, nil
}

func (e *environment) V(name string) (interface{}, error) {
	ottoVal, err := e.vm.Get(name)
	if err != nil {
		return nil, err
	}
	goVal, err := ottoVal.Export()
	if err != nil {
		return nil, err
	}
	return goVal, nil
}
