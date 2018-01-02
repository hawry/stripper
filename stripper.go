// Package stripper provides functions to remove values from specified struct fields, e.g. to reduce the risk of sending protected attributes in a JSON-response in an API.
package stripper

import (
	"encoding/json"
	"errors"
	"reflect"
)

const tagName = "clean"

//Marshal removes (setting to default Zero value) any data on struct fields with the tag `clean:"true"` on them, but leaves the rest, and then calls json.Marshal
//
// Examples of struct field tags and their meanings:
//  // Field will be set to the Zero value (reflect.Zero)
//  Field int `clean:"true"`
//
//  // Nested Struct will also be checked (as specified by those struct field tags)
//  Field Struct `clean:"true"`
// All other values than 'true' will make the cleaner ignore that field
func Marshal(a interface{}) ([]byte, error) {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Ptr {
		return []byte{}, errors.New("must provide a pointer")
	}
	cpy := clean(v)
	return json.Marshal(cpy)
}

//MarshalIndent returns a cleaned json marshalled response
func MarshalIndent(a interface{}, prefix, indent string) ([]byte, error) {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Ptr {
		return []byte{}, errors.New("must provide a pointer")
	}
	cpy := clean(v)
	return json.MarshalIndent(cpy, prefix, indent)
}

//Clean modifies the interface given and returns the cleaned version, the response must be type asserted
func Clean(a interface{}) (interface{}, error) {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Ptr {
		return nil, errors.New("must provide a pointer")
	}
	cpy := clean(v)
	return cpy, nil
}

func clean(v reflect.Value) interface{} {
	in := reflect.Indirect(v)
	for i := 0; i < in.NumField(); i++ {
		tag := in.Type().Field(i).Tag.Get(tagName)
		if tag != "true" {
			continue
		}
		if in.Field(i).Kind() == reflect.Struct {
			clean(in.Field(i))
			continue
		}
		in.Field(i).Set(reflect.Zero(in.Field(i).Type()))
	}
	return v.Interface()
}
