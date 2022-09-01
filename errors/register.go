package errors

import (
	"fmt"
	"net/http"
	"sync"
)

// ErrCode implements errors.Coder interface.
type ErrCode struct {
	// C refers to the code of the ErrCode.
	C int
	// HTTP status that should be used for the associated error code.
	HTTP int
	// External (user) facing error text.
	Ext string
	// Ref specify the reference document.
	Ref string
}

var _ Coder = &ErrCode{}

// Code returns the integer code of ErrCode.
func (coder ErrCode) Code() int {
	return coder.C
}

// String implements stringer. String returns the external error message,
// if any.
func (coder ErrCode) String() string {
	return coder.Ext
}

// Reference returns the reference document.
func (coder ErrCode) Reference() string {
	return coder.Ref
}

// HTTPStatus returns the associated HTTP status code, if any. Otherwise,
// returns 200.
func (coder ErrCode) HTTPStatus() int {
	if coder.HTTP == 0 {
		return http.StatusInternalServerError
	}

	return coder.HTTP
}

var codemap = map[int]bool{
	200: true, 400: true, 401: true, 403: true, 404: true, 500: true,
}

// codes contains a map of error codes to metadata.
var codes = map[int]Coder{}
var codeMux = &sync.Mutex{}

// Register register a user define error code.
// It will overrid the exist code.
//func Register(coder Coder) {
//	if coder.Code() == 0 {
//		panic("code `0` is unknownCode error code")
//	}
//
//	codeMux.Lock()
//	defer codeMux.Unlock()
//
//	codes[coder.Code()] = coder
//}

func Register(httpCode, errCode int, message string, refs ...string) {
	if _, ok := codemap[httpCode]; !ok {
		panic("http status code is not supported")
	}

	var reference string
	if len(refs) > 0 {
		reference = refs[0]
	}

	coder := &ErrCode{
		C:    errCode,
		HTTP: httpCode,
		Ext:  message,
		Ref:  reference,
	}

	MustRegister(coder)
}

// MustRegister register a user define error code.
// It will panic when the same Code already exist.
func MustRegister(coder Coder) {
	if coder.Code() == 0 {
		panic("code '0' is ErrUnknown error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	if _, ok := codes[coder.Code()]; ok {
		panic(fmt.Sprintf("code: %d already exist", coder.Code()))
	}

	codes[coder.Code()] = coder
}

// ParseCoder parse any error into *withCode.
// nil error will return nil direct.
// None withStack error will be parsed as ErrUnknown.
func ParseCoder(err error) Coder {
	if err == nil {
		return nil
	}

	if v, ok := err.(*withCode); ok {
		if coder, ok := codes[v.code]; ok {
			return coder
		}
	}

	return unknownCoder
}
