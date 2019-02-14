package goutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"regexp"

	"github.com/leesper/holmes"
	"gopkg.in/mgo.v2/bson"
	validator "gopkg.in/validator.v2"
)

// DecodeAndValidate JSON解码并验证参数是否合法
func DecodeAndValidate(body io.Reader, input interface{}) error {
	err := json.NewDecoder(body).Decode(input)
	if err != nil {
		holmes.Errorln(err, input)
		return err
	}
	holmes.Debugf("input: %+v", input)
	if err = validator.Validate(input); err != nil {
		holmes.Errorln(err)
		return err
	}
	return err
}

func init() {
	validator.SetValidationFunc("phone", validPhone)
	validator.SetValidationFunc("email", validEmail)
	validator.SetValidationFunc("price", validPrice)
	validator.SetValidationFunc("pricture", validPicture)
	validator.SetValidationFunc("longitude", validLongitude)
	validator.SetValidationFunc("latitude", validLatitude)
	validator.SetValidationFunc("objectid", validObjectID)
}

func validPhone(v interface{}, param string) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.String {
		return validator.ErrUnsupported
	}
	phone := val.String()
	isValid := regexp.MustCompile(`^1[3|4|5|6|7|8|9][0-9]{9}$`).MatchString(phone)
	if !isValid {
		return errors.New("not a valid phone number")
	}
	return nil
}

func validEmail(v interface{}, param string) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.String {
		return validator.ErrUnsupported
	}
	addr := val.String()
	isValid := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(addr)
	if !isValid {
		return errors.New("not a valid email address")
	}
	return nil
}

func validObjectID(v interface{}, param string) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.String {
		return validator.ErrUnsupported
	}
	objectID := val.String()
	if !bson.IsObjectIdHex(objectID) {
		return errors.New("not a valid ObjectID")
	}
	return nil
}

func validPrice(v interface{}, param string) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Int {
		holmes.Errorf("%T\n", v)
		return validator.ErrUnsupported
	}
	price := val.Int()
	if price <= 0 {
		return errors.New("price should be larger than 0")
	}
	return nil
}

func validPicture(v interface{}, param string) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Slice {
		return validator.ErrUnsupported
	}
	if reflect.TypeOf(v).Elem().Kind() != reflect.String {
		return validator.ErrUnsupported
	}
	pictures := v.([]string)
	size := 0
	for _, p := range pictures {
		size += len(p)
	}
	if size >= 10*(1<<20) {
		return errors.New("picture size is more than 10M")
	}
	return nil
}

func validLongitude(v interface{}, param string) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Float64 {
		return validator.ErrUnsupported
	}
	if val.Float() >= -180.0 && val.Float() <= 180.0 {
		return nil
	}
	return fmt.Errorf("longitude value %.2f", val.Float())
}

func validLatitude(v interface{}, param string) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Float64 {
		return validator.ErrUnsupported
	}
	if val.Float() >= -180.0 && val.Float() <= 180.0 {
		return nil
	}
	return fmt.Errorf("latitude value %.2f", val.Float())
}
