package mongo_client

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

type Int64Decoder struct{}

func (dt Int64Decoder) DecodeValue(ctx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	var err error
	var dtValue int64

	switch vr.Type() {
	case bson.TypeDateTime:
		dtValue, err = vr.ReadDateTime()
		if err != nil {
			return err
		}
		dtValue = dtValue / 1000 // MongoDB stores time in milliseconds, so no conversion is needed
	case bson.TypeInt64:
		dtValue, err = vr.ReadInt64()
		if err != nil {
			return err
		}
	case bson.TypeInt32:
		var i32 int32
		i32, err = vr.ReadInt32()
		if err != nil {
			return err
		}
		dtValue = int64(i32)
	default:
		return bsoncodec.ValueDecoderError{Name: "Int64Decoder", Types: []reflect.Type{reflect.TypeOf(int64(0))}, Received: val}
	}

	if val.Kind() == reflect.Int64 {
		val.SetInt(dtValue)
		return nil
	}

	return bsoncodec.ValueDecoderError{Name: "Int64Decoder", Types: []reflect.Type{reflect.TypeOf(int64(0))}, Received: val}
}
