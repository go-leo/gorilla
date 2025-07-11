package gorilla

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// --- Mocks ---

type mockProto_Decoder struct {
	proto.Message
	Data string `json:"data"`
}

func (m *mockProto_Decoder) Reset()         {}
func (m *mockProto_Decoder) String() string { return m.Data }
func (m *mockProto_Decoder) ProtoMessage()  {}

func (m *mockProto_Decoder) MarshalJSON() ([]byte, error) {
	return []byte(`{"data":"` + m.Data + `"}`), nil
}

func (m *mockProto_Decoder) UnmarshalJSON(b []byte) error {
	m.Data = string(bytes.Trim(b, `{"data":}`))
	return nil
}

func TestRequestDecoder(t *testing.T) {
	msg := &httpbody.HttpBody{}
	body := `{"content_type":"json"}`
	r := &http.Request{
		Body: io.NopCloser(strings.NewReader(body)),
	}
	opts := protojson.UnmarshalOptions{}
	err := RequestDecoder(context.Background(), r, msg, opts)
	if err != nil {
		t.Fatalf("RequestDecoder error: %v", err)
	}
	if msg.ContentType != `"json"` && msg.ContentType != "json" {
		t.Errorf("msg.Data = %q, want \"hello\"", msg.Data)
	}
}

func TestHttpBodyDecoder(t *testing.T) {
	data := "abc"
	r := &http.Request{
		Body:   io.NopCloser(strings.NewReader(data)),
		Header: http.Header{ContentTypeKey: []string{"application/test"}},
	}
	body := &httpbody.HttpBody{}
	err := HttpBodyDecoder(context.Background(), r, body)
	if err != nil {
		t.Fatalf("HttpBodyDecoder error: %v", err)
	}
	if string(body.Data) != data {
		t.Errorf("body.Data = %q, want %q", body.Data, data)
	}
	if body.ContentType != "application/test" {
		t.Errorf("body.ContentType = %q, want application/test", body.ContentType)
	}
}

func TestHttpRequestDecoder(t *testing.T) {
	r := &http.Request{
		Method: "POST",
		URL:    &url.URL{Scheme: "http", Host: "localhost", Path: "/foo"},
		Body:   io.NopCloser(strings.NewReader("xyz")),
		Header: http.Header{"X-Foo": []string{"bar", "baz"}},
	}
	req := &rpchttp.HttpRequest{}
	err := HttpRequestDecoder(context.Background(), r, req)
	if err != nil {
		t.Fatalf("HttpRequestDecoder error: %v", err)
	}
	if req.Method != "POST" {
		t.Errorf("req.Method = %q, want POST", req.Method)
	}
	if req.Uri != "http://localhost/foo" {
		t.Errorf("req.Uri = %q, want http://localhost/foo", req.Uri)
	}
	if string(req.Body) != "xyz" {
		t.Errorf("req.Body = %q, want xyz", req.Body)
	}
	var foundBar, foundBaz bool
	for _, h := range req.Headers {
		if h.Key == "X-Foo" && h.Value == "bar" {
			foundBar = true
		}
		if h.Key == "X-Foo" && h.Value == "baz" {
			foundBaz = true
		}
	}
	if !foundBar || !foundBaz {
		t.Errorf("req.Headers missing expected values: %v", req.Headers)
	}
}

func TestFormDecoder(t *testing.T) {
	form := url.Values{}
	form.Set("a", "1")
	getter := func(f url.Values, key string) (int, error) {
		return 42, nil
	}
	v, err := FormDecoder(nil, form, "a", getter)
	if err != nil || v != 42 {
		t.Errorf("FormDecoder = %v, %v; want 42, nil", v, err)
	}

	preErr := errors.New("fail")
	v, err = FormDecoder(preErr, form, "a", getter)
	if err != preErr {
		t.Errorf("FormDecoder with pre error = %v, %v; want pre error", v, err)
	}
}

func TestBreak(t *testing.T) {
	// pre error
	pre := errors.New("fail")
	f := Break[int](pre)
	v, err := f(func() (int, error) { return 1, nil })
	if err != pre || v != 0 {
		t.Errorf("Break with pre error = %v, %v; want 0, pre error", v, err)
	}

	// no error
	f = Break[int](nil)
	v, err = f(func() (int, error) { return 42, nil })
	if err != nil || v != 42 {
		t.Errorf("Break = %v, %v; want 42, nil", v, err)
	}

	// function error
	f = Break[int](nil)
	wantErr := errors.New("fail2")
	v, err = f(func() (int, error) { return 0, wantErr })
	if err != wantErr {
		t.Errorf("Break = %v, %v; want 0, wantErr", v, err)
	}
}

func TestFormDecoder_GenericType(t *testing.T) {
	form := url.Values{}
	form.Set("foo", "bar")
	getter := func(f url.Values, key string) (string, error) {
		return f.Get(key), nil
	}
	v, err := FormDecoder(nil, form, "foo", getter)
	if err != nil || v != "bar" {
		t.Errorf("FormDecoder generic = %v, %v; want bar, nil", v, err)
	}
}

func TestFormDecoder_ErrorPropagation(t *testing.T) {
	form := url.Values{}
	getter := func(f url.Values, key string) (string, error) {
		return "", errors.New("fail")
	}
	v, err := FormDecoder(nil, form, "foo", getter)
	if err == nil {
		t.Errorf("FormDecoder should propagate error")
	}
	if v != "" {
		t.Errorf("FormDecoder should return zero value on error")
	}
}

func TestFormDecoder_NilForm(t *testing.T) {
	getter := func(f url.Values, key string) (string, error) {
		if f == nil {
			return "ok", nil
		}
		return "", nil
	}
	v, err := FormDecoder(nil, nil, "foo", getter)
	if err != nil || v != "ok" {
		t.Errorf("FormDecoder nil form = %v, %v; want ok, nil", v, err)
	}
}

func TestFormGetter_Type(t *testing.T) {
	var _ FormGetter[int]
	var _ FormGetter[string]
}

func TestBreak_Type(t *testing.T) {
	_ = Break[int]
	_ = Break[string]
}
