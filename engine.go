package gojst

import (
	"bytes"
	"html/template"
	"io"

	"github.com/robertkrimen/otto"
	//this will inject underscore to all javascript functions
	_ "github.com/robertkrimen/otto/underscore"
)

type Engine struct {
	env *environment
}

func NewEngine(script io.Reader, data interface{}) (*Engine, error) {
	vm := otto.New()
	err := vm.Set("data", data)
	if err != nil {
		return nil, err
	}
	_, err = vm.Run(script)
	if err != nil {
		return nil, err
	}

	return &Engine{&environment{vm: vm, D: data}}, nil
}

func (e *Engine) SetData(data interface{}) error {
	if err := e.Set("data", data); err != nil {
		return err
	}
	e.env.D = data
	return nil
}

func (e *Engine) Set(variable string, body interface{}) error {
	return e.env.vm.Set(variable, body)
}

func (e *Engine) Check(expr string) (bool, error) {
	return e.EvalBool(expr)
}

func (e *Engine) Eval(expr string) error {
	_, err := e.env.vm.Eval(expr)
	return err
}

func (e *Engine) EvalBool(expr string) (bool, error) {
	ottoVal, err := e.env.vm.Eval(expr)
	if err != nil {
		return false, err
	}
	goVal, err := ottoVal.ToBoolean()
	if err != nil {
		return false, err
	}
	return goVal, nil
}

func (e *Engine) EvalInt(expr string) (int64, error) {
	ottoVal, err := e.env.vm.Eval(expr)
	if err != nil {
		return 0, err
	}
	goVal, err := ottoVal.ToInteger()
	if err != nil {
		return 0, err
	}
	return goVal, nil
}

func (e *Engine) EvalString(expr string) (string, error) {
	ottoVal, err := e.env.vm.Eval(expr)
	if err != nil {
		return "", err
	}
	goVal, err := ottoVal.ToString()
	if err != nil {
		return "", err
	}
	return goVal, nil
}

func (e *Engine) Render(str string) (string, error) {
	tpl, err := template.New("").Parse(str)
	if err != nil {
		return "", err
	}
	var res bytes.Buffer
	if err := tpl.Execute(&res, e.env); err != nil {
		return "", err
	}
	return res.String(), nil
}
