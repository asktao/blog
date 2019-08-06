package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestString struct {
	Blank			string	`valid:""`
	Required		string	`valid:"required"`
	StrMaxlength	string 	`valid:"str-max-len=10"`
	StrMinLength	string  `valid:"str-min-len=1"`
	StrLength		string	`valid:"str-len=3"`
}

type TestInt struct {
	Max		int64	`valid:"int-max=1000"`
	Min		int64	`valid:"int-min=-20"`
}

type TestValidation struct {
	Validation
}

func TestStringValidate(t *testing.T) {

	assert := assert.New(t)

	tv := TestValidation{}

	success01 := &TestString{
		Blank:			"",
		Required:		"required",
		StrMaxlength:	"123456790",
		StrMinLength:	"123",
		StrLength:		"123",
	}
	successResponse01 := tv.Validator(success01)
	assert.Nil(successResponse01)

	fail01 := &TestString{
		Blank:			"",
		Required:		"",
	}
	failResponse01 := tv.Validator(fail01)
	assert.EqualError(failResponse01, "The Required field is required.")

	fail02 := &TestString{
		Required:		"required",
		StrMaxlength:	"1234567890A",
	}
	failResponse02 := tv.Validator(fail02)
	assert.EqualError(failResponse02, "The StrMaxlength may not be greater than 10 characters.")

	fail03 := &TestString{
		Required:		"required",
		StrMaxlength:	"1234567890",
		StrMinLength:	"",
	}
	failResponse03 := tv.Validator(fail03)
	assert.EqualError(failResponse03, "The StrMinLength must be at least 1 characters.")

	fail04 := &TestString{
		Required:		"required",
		StrMaxlength:	"1234567890",
		StrMinLength:	"12",
		StrLength:		"1234",
	}
	failResponse04 := tv.Validator(fail04)
	assert.EqualError(failResponse04, "The StrLength must be 3 characters.")
}

func TestIntValidate(t *testing.T) {

	assert := assert.New(t)

	tv := TestValidation{}

	success01 := &TestInt{
		Max:	100,
		Min:	-10,
	}
	successResponse01 := tv.Validator(success01)
	assert.Nil(successResponse01)

	fail01 := &TestInt{
		Max:	1000000,
	}
	failResponse01 := tv.Validator(fail01)
	assert.EqualError(failResponse01, "The Max may not be greater than 1000.")

	fail02 := &TestInt{
		Max:	100,
		Min:	-100,
	}
	failResponse02 := tv.Validator(fail02)
	assert.EqualError(failResponse02, "The Min must be at least -20.")
}