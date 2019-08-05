package controllers

import (
	"testing"
)

type User struct {
	Name	string	`valid:"Required"		json:"name"`
	Age		int		`valid:"int-max=100"	json"age"`
}

var Users = []User{
	{
		"AT",
		29,
	},
	{
		"Zhao",
		20,
	},
}

type TestValidation struct {
	Validation
}

func TestValidate(t *testing.T) {
	for _, user := range Users {
		tv := TestValidation{}
		err := tv.Validator(&user)
		if err != nil {
			t.Errorf("validation error")
		}
	}
}