package util

import "google.golang.org/protobuf/encoding/protojson"

var Unmarshaler = protojson.UnmarshalOptions{
	DiscardUnknown: true,
}
