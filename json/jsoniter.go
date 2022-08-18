package json

import jsoniter "github.com/json-iterator/go"

var (
	Marshal             = jsoniter.ConfigCompatibleWithStandardLibrary.Marshal
	MarshalIndent       = jsoniter.ConfigCompatibleWithStandardLibrary.MarshalIndent
	MarshalToString     = jsoniter.ConfigCompatibleWithStandardLibrary.MarshalToString
	Unmarshal           = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal
	UnmarshalFromString = jsoniter.ConfigCompatibleWithStandardLibrary.UnmarshalFromString
	NewDecoder          = jsoniter.ConfigCompatibleWithStandardLibrary.NewDecoder
	NewEncoder          = jsoniter.ConfigCompatibleWithStandardLibrary.NewEncoder
	Get                 = jsoniter.ConfigCompatibleWithStandardLibrary.Get
	Valid               = jsoniter.ConfigCompatibleWithStandardLibrary.Valid
)
