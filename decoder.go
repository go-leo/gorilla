package gorilla

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func RequestDecoder(ctx context.Context, r *http.Request, req proto.Message, unmarshalOptions protojson.UnmarshalOptions) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := unmarshalOptions.Unmarshal(data, req); err != nil {
		return err
	}
	return nil
}

func HttpBodyDecoder(ctx context.Context, r *http.Request, body *httpbody.HttpBody) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	body.Data = data
	body.ContentType = r.Header.Get(ContentTypeKey)
	return nil
}

func HttpRequestDecoder(ctx context.Context, r *http.Request, request *rpchttp.HttpRequest) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	request.Method = r.Method
	request.Uri = r.URL.String()
	for key, values := range r.Header {
		for _, value := range values {
			request.Headers = append(request.Headers, &rpchttp.HttpHeader{Key: key, Value: value})
		}
	}
	request.Body = data
	return nil
}

type FormGetter[T any] func(form url.Values, key string) (T, error)

func FormDecoder[T any](pre error, form url.Values, key string, f FormGetter[T]) (T, error) {
	return Break[T](pre)(func() (T, error) { return f(form, key) })
}

// Break 函数是一个高阶函数，用于处理错误并决定是否继续执行另一个函数。
// 如果pre参数不为nil，它会返回一个零值和pre错误，从而中断后续的执行；
// 如果pre为nil，则会调用并返回f()的结果。
func Break[T any](pre error) func(f func() (T, error)) (T, error) {
	return func(f func() (T, error)) (T, error) {
		if pre != nil {
			var v T       // Declare a zero value of type T
			return v, pre // If pre is not nil, return the zero value of T and the error
		}
		return f() // If pre is nil, execute the wrapped function and return its result
	}
}
