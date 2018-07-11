package goutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"

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
	if err = validator.Validate(input); err != nil {
		holmes.Errorln(err)
		return err
	}
	return err
}

func init() {
	validator.SetValidationFunc("validprice", validPrice)
	validator.SetValidationFunc("validpictures", validPictures)
	validator.SetValidationFunc("validlongitude", validLongitude)
	validator.SetValidationFunc("validlatitude", validLatitude)
	validator.SetValidationFunc("validobjectid", validObjectID)
}

func validObjectID(v interface{}, param string) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.String {
		return validator.ErrUnsupported
	}
	objectID := val.String()
	if !bson.IsObjectIdHex(objectID) {
		return errors.New("invalid param: not a valid ObjectID")
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
		return errors.New("invalid param: price should be larger than 0")
	}
	return nil
}

func validPictures(v interface{}, param string) error {
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
		return errors.New("invalid param: picture size is more than 10M")
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
	return fmt.Errorf("invalid param: longitude value %.2f", val.Float())
}

func validLatitude(v interface{}, param string) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Float64 {
		return validator.ErrUnsupported
	}
	if val.Float() >= -180.0 && val.Float() <= 180.0 {
		return nil
	}
	return fmt.Errorf("invalid param: latitude value %.2f", val.Float())
}
