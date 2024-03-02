package mongo_client

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DateTimeDecoder struct{}

func (dt DateTimeDecoder) DecodeValue(ctx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if vr.Type() != bson.TypeDateTime {
		return bsoncodec.ValueDecoderError{Name: "DateTimeDecoder", Types: []reflect.Type{reflect.TypeOf(primitive.DateTime(0))}, Received: val}
	}

	dtValue, err := vr.ReadDateTime()
	if err != nil {
		return err
	}

	// MongoDB stores time in milliseconds, so no conversion is needed
	if val.Kind() == reflect.Int64 {
		val.SetInt(dtValue)
		return nil
	}

	return bsoncodec.ValueDecoderError{Name: "DateTimeDecoder", Types: []reflect.Type{reflect.TypeOf(int64(0))}, Received: val}
}
