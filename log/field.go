package log

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"go.uber.org/zap/zapcore"
)

type context struct {
	request  *request
	response *response
}

func (c *context) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	if c.request != nil {
		if err := encoder.AddObject("request", c.request); err != nil {
			return err
		}
	}
	if c.response != nil {
		if err := encoder.AddObject("response", c.response); err != nil {
			return err
		}
	}

	return nil
}

type request struct {
	Method string
	Host   string
	Path   string
	Proto  string
	Header http.Header
	Form   url.Values
	Body   any
}

func (r *request) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("method", r.Method)
	encoder.AddString("host", r.Host)
	encoder.AddString("path", r.Path)
	encoder.AddString("proto", r.Proto)

	err := encoder.AddReflected("header", r.Header)
	if err != nil {
		return err
	}
	err = encoder.AddReflected("form", r.Form)
	if err != nil {
		return err
	}
	err = encoder.AddReflected("body", r.Body)
	if err != nil {
		return err
	}

	return nil
}

type response struct {
	Status     string
	StatusCode int
	Proto      string
	Header     http.Header
	Body       string
}

func (r *response) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("status", r.Status)
	encoder.AddInt("status_code", r.StatusCode)
	encoder.AddString("proto", r.Proto)
	encoder.AddString("body", r.Body)

	err := encoder.AddReflected("header", r.Header)
	if err != nil {
		return err
	}

	return nil
}

func Context(body any, raw *http.Response) *context {
	if raw == nil || raw.Body == nil {
		return &context{}
	}

	bytes, err := io.ReadAll(raw.Body)
	if err != nil {
		log.Printf("error reading response body: %s", err)
		return &context{}
	}

	return &context{
		request: &request{
			Method: raw.Request.Method,
			Host:   raw.Request.URL.Host,
			Path:   raw.Request.URL.Path,
			Proto:  raw.Request.Proto,
			Header: raw.Request.Header,
			Form:   raw.Request.Form,
			Body:   body,
		},
		response: &response{
			Status:     raw.Status,
			StatusCode: raw.StatusCode,
			Proto:      raw.Proto,
			Header:     raw.Header,
			Body:       string(bytes),
		},
	}
}
