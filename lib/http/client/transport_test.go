package client

import (
	"bytes"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestInitHeader(t *testing.T) {
	r := &http.Request{}
	initHeaders(r)

	if r.Header == nil {
		t.Errorf("expcted non-nil headers")
	}
}

func TestContentLength(t *testing.T) {
	r := &http.Request{
		ContentLength: 100,
		Body:          bytesReaderCloser{bytes.NewReader([]byte{})},
	}
	initHeaders(r)

	if !reflect.DeepEqual(r.Header["Content-Length"], []string{"100"}) {
		t.Errorf("expected content-length header to be [100], got %s", r.Header["Content-Length"])
	}
}

func TestConnectionClose(t *testing.T) {
	r := &http.Request{
		Close: true,
	}
	initHeaders(r)

	if !reflect.DeepEqual(r.Header["Connection"], []string{"close"}) {
		t.Errorf("expected content-length header to be [close], got %s", r.Header["Content-Length"])
	}
}

func TestUserAgent(t *testing.T) {
	r := &http.Request{}
	initHeaders(r)

	if !reflect.DeepEqual(r.Header["User-Agent"], []string{DEFAULT_USER_AGENT}) {
		t.Errorf("expected content-length header to be [%s], got %s", DEFAULT_USER_AGENT, r.Header["User-Agent"])
	}

	agent := "Override"
	r.Header["User-Agent"] = []string{agent}
	if !reflect.DeepEqual(r.Header["User-Agent"], []string{agent}) {
		t.Errorf("expected content-length header to be [%s], got %s", agent, r.Header["User-Agent"])
	}
}

func TestMakeHeaders(t *testing.T) {
	r := &http.Request{
		Header: http.Header{
			"User-Agent": []string{"foo"},
		},
	}

	headers := makeHeaders(r)
	slice := headers.Slice()
	if len(slice) != 1 {
		t.Errorf("Unexpected header length")
		t.FailNow()
	}
	if slice[0].F0 != "User-Agent" {
		t.Errorf("Unexpected field key: %s", slice[0].F0)
	}
	byteSlice := slice[0].F1.Slice()
	if string(byteSlice) != "foo" {
		t.Errorf("Unexpected values: %v", slice[0].F1.Slice())
	}
}

func TestMakeMultipleHeaders(t *testing.T) {
	r := &http.Request{
		Header: http.Header{
			"User-Agent": []string{
				"foo",
				"bar",
			},
		},
	}

	headers := makeHeaders(r)
	slice := headers.Slice()
	if len(slice) != 2 {
		t.Errorf("Unexpected header length")
		t.FailNow()
	}
	for i := 0; i < 2; i++ {
		if slice[i].F0 != "User-Agent" {
			t.Errorf("Unexpected field key: %s", slice[0].F0)
		}
	}
	byteSlice := slice[0].F1.Slice()
	if string(byteSlice) != "foo" {
		t.Errorf("Unexpected values: %v", slice[0].F1.Slice())
	}
	byteSlice = slice[1].F1.Slice()
	if string(byteSlice) != "bar" {
		t.Errorf("Unexpected values: %v", slice[1].F1.Slice())
	}
}

func TestGetAuthority(t *testing.T) {
	u, _ := url.Parse("http://company.com")
	authority := getAuthority(&http.Request{
		Host: "foo",
		URL:  u,
	})
	if authority != "foo" {
		t.Errorf("unexpected authority: %s", authority)
	}
	authority = getAuthority(&http.Request{
		URL: u,
	})
	if authority != "company.com" {
		t.Errorf("unexpected authority: %s", authority)
	}
}

func TestPopulateHeaders(t *testing.T) {
	headers := http.Header{
		"Status":      []string{"200"},
		"Multi-Value": []string{"foo", "bar"},
	}
	r := http.Response{}
	populateResponseHeaders(makeHeaders(&http.Request{
		Header: headers,
	}), &r)

	if !reflect.DeepEqual(r.Header, headers) {
		t.Errorf("unexpected headers: %v", r.Header)
	}
}
