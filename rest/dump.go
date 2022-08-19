package rest

import (
	"go.uber.org/zap/zapcore"
)

type DumpContext struct {
	dumpRequest  *dumpRequest
	dumpResponse *dumpResponse
}

func (o *DumpContext) MarshalLogObject(e zapcore.ObjectEncoder) error {
	if o.dumpRequest == nil {
		return nil
	}
	if err := e.AddObject("request", o.dumpRequest); err != nil {
		return err
	}

	if o.dumpResponse == nil {
		return nil
	}
	if err := e.AddObject("response", o.dumpResponse); err != nil {
		return err
	}

	return nil
}

type dumpRequest struct {
	Method string
	Host   string
	Path   string
	Proto  string
	Header map[string][]string
	Form   map[string][]string
	Body   any
}

func (o *dumpRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {
	if err := e.AddReflected("form", o.Form); err != nil {
		return err
	}
	if err := e.AddReflected("body", o.Body); err != nil {
		return err
	}
	if err := e.AddReflected("header", o.Header); err != nil {
		return err
	}

	e.AddString("method", o.Method)
	e.AddString("host", o.Host)
	e.AddString("path", o.Path)
	e.AddString("proto", o.Proto)

	return nil
}

type dumpResponse struct {
	Status string
	Proto  string
	Header map[string][]string
	Body   string
}

func (o *dumpResponse) MarshalLogObject(e zapcore.ObjectEncoder) error {
	if err := e.AddReflected("header", o.Header); err != nil {
		return err
	}

	e.AddString("status", o.Status)
	e.AddString("proto", o.Proto)
	e.AddString("body", o.Body)

	return nil
}
