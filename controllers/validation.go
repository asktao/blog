package controllers

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

const (
	REQUIRED     = "required"
	MAX          = "int-max"
	MIN          = "int-min"
	TYPE         = "type"
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

		if strings.Index(valid, TYPE) != -1 {
			v := value.Type().Name()
			split := strings.Split(valid, "=")
			t := split[1]
			if v != t {
				return errors.New(fieldName + " must be a" + t)
			}
		}

		if strings.Index(valid, REQUIRED) != -1 {
			str := value.String()
			if str == "" {
				return errors.New(fieldName + " field is required")
			}
		}
		if strings.Index(valid, MIN) != -1 {
			v := value.Int()
			split := strings.Split(valid, "=")
			rule, err := strconv.Atoi(split[1])
			if err != nil {
				return errors.New(fieldName + ":invalid validation method")
			}
			if int(v) < rule {
				return errors.New(fieldName + " field may not be greater" + strconv.Itoa(rule))
			}
		}

		if strings.Index(valid, MAX) != -1 {
			v := value.Int()
			split := strings.Split(valid, "=")
			rule, err := strconv.Atoi(split[1])
			if err != nil {
				return errors.New(fieldName + ":invalid validation method")
			}
			if int(v) > rule {
				return errors.New(fieldName + " field must <= " + strconv.Itoa(rule))
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
					return errors.New(fieldName + " str length  must be " + lenStr)
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
					return errors.New(fieldName + " str length  <= " + lenStr)
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
					return errors.New(fieldName + " field length  >= " + lenStr)
				}
			}
		}
	}
	return nil
}