package controllers

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

const (
	Required     = "required"
	Max          = "int-max"
	Min          = "int-min"
//	Type         = "type"
	StrMaxLength = "str-max-len"
	StrMinLength = "str-min-len"
	StrLength    = "str-len"
)

type Validation struct {

}

type ValidationInterface interface {
	Validator(bean interface{}) error
}

func (v *Validation) Validator(bean interface{}) error {

	fields := reflect.ValueOf(bean).Elem()
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Type().Field(i)
		valid := field.Tag.Get("valid")
		if valid == "" {
			continue
		}
		value := fields.FieldByName(field.Name)
		err := fieldValidate(field.Name, valid, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func fieldValidate(fieldName, valid string, value reflect.Value) error {
	valids := strings.Split(valid, " ")
	for _, valid := range valids {
		//if strings.Index(valid, Type) != -1 {
		//	v := value.Type().Name()
		//	split := strings.Split(valid, "=")
		//	t := split[1]
		//	if v != t {
		//		return errors.New("The " + fieldName + " must be a" + t)
		//	}
		//}

		if strings.Index(valid, Required) != -1 {
			str := value.String()
			if str == "" {
				return errors.New("The " + fieldName + " field is required.")
			}
		}
		if strings.Index(valid, Min) != -1 {
			v := value.Int()
			split := strings.Split(valid, "=")
			rule, err := strconv.Atoi(split[1])
			if err != nil {
				return errors.New(fieldName + ":invalid validation method")
			}
			if int(v) < rule {
				return errors.New("The " + fieldName + " must be at least " + strconv.Itoa(rule) + ".")

			}
		}

		if strings.Index(valid, Max) != -1 {
			v := value.Int()
			split := strings.Split(valid, "=")
			rule, err := strconv.Atoi(split[1])
			if err != nil {
				return errors.New(fieldName + ":invalid validation method")
			}
			if int(v) > rule {
				return errors.New("The " + fieldName + " may not be greater than " + strconv.Itoa(rule) + ".")
			}
		}
		//字符串特殊处理
		if value.Type().Name() == "string" {
			if strings.Index(valid, StrLength) != -1 {
				v := value.String()
				split := strings.Split(valid, "=")
				lenStr := split[1]
				length, err := strconv.Atoi(lenStr)
				if err != nil {
					return errors.New(fieldName + " " + StrLength + " rule is error")
				}
				if len(v) != length {
					return errors.New("The " + fieldName + " must be " + lenStr + " characters.")
				}
			}
			if strings.Index(valid, StrMaxLength) != -1 {
				v := value.String()
				split := strings.Split(valid, "=")
				lenStr := split[1]
				length, err := strconv.Atoi(lenStr)
				if err != nil {
					return errors.New(fieldName + " " + StrLength + " rule is error")
				}
				if len(v) > length {
					return errors.New("The " + fieldName + " may not be greater than " + lenStr + " characters.")
				}
			}

			if strings.Index(valid, StrMinLength) != -1 {
				v := value.String()
				split := strings.Split(valid, "=")
				lenStr := split[1]
				length, err := strconv.Atoi(lenStr)
				if err != nil {
					return errors.New(fieldName + " " + StrLength + " rule is error")
				}
				if len(v) < length {
					return errors.New("The " + fieldName + " must be at least " + lenStr + " characters.")
				}
			}
		}
	}
	return nil
}