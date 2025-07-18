// Code generated by protoc-gen-gorilla. DO NOT EDIT.

package query

import (
	context "context"
	gorilla "github.com/go-leo/gorilla"
	mux "github.com/gorilla/mux"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	http "net/http"
)

type BoolQueryGorillaService interface {
	BoolQuery(ctx context.Context, request *BoolQueryRequest) (*httpbody.HttpBody, error)
}

func AppendBoolQueryGorillaRoute(router *mux.Router, service BoolQueryGorillaService, opts ...gorilla.Option) *mux.Router {
	options := gorilla.NewOptions(opts...)
	handler := boolQueryGorillaHandler{
		service: service,
		decoder: boolQueryGorillaRequestDecoder{
			unmarshalOptions:        options.UnmarshalOptions(),
			shouldFailFast:          options.ShouldFailFast(),
			onValidationErrCallback: options.OnValidationErrCallback(),
		},
		encoder: boolQueryGorillaEncodeResponse{
			marshalOptions:      options.MarshalOptions(),
			unmarshalOptions:    options.UnmarshalOptions(),
			responseTransformer: options.ResponseTransformer(),
		},
		errorEncoder: gorilla.DefaultEncodeError,
	}
	router.NewRoute().
		Name("/leo.gorilla.example.query.v1.BoolQuery/BoolQuery").
		Methods("GET").
		Path("/v1/bool").
		Handler(gorilla.Chain(handler.BoolQuery(), options.Middlewares()...))
	return router
}

type boolQueryGorillaHandler struct {
	service      BoolQueryGorillaService
	decoder      boolQueryGorillaRequestDecoder
	encoder      boolQueryGorillaEncodeResponse
	errorEncoder gorilla.ErrorEncoder
}

func (h boolQueryGorillaHandler) BoolQuery() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.BoolQuery(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.BoolQuery(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.BoolQuery(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

type boolQueryGorillaRequestDecoder struct {
	unmarshalOptions        protojson.UnmarshalOptions
	shouldFailFast          bool
	onValidationErrCallback gorilla.OnValidationErrCallback
}

func (decoder boolQueryGorillaRequestDecoder) BoolQuery(ctx context.Context, r *http.Request) (*BoolQueryRequest, error) {
	req := &BoolQueryRequest{}
	if ok, err := gorilla.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	queries := r.URL.Query()
	var queryErr error
	req.Bool, queryErr = gorilla.DecodeForm[bool](queryErr, queries, "bool", gorilla.GetBool)
	req.OptBool, queryErr = gorilla.DecodeForm[*bool](queryErr, queries, "opt_bool", gorilla.GetBoolPtr)
	req.WrapBool, queryErr = gorilla.DecodeForm[*wrapperspb.BoolValue](queryErr, queries, "wrap_bool", gorilla.GetBoolValue)
	req.ListBool, queryErr = gorilla.DecodeForm[[]bool](queryErr, queries, "list_bool", gorilla.GetBoolSlice)
	req.ListWrapBool, queryErr = gorilla.DecodeForm[[]*wrapperspb.BoolValue](queryErr, queries, "list_wrap_bool", gorilla.GetBoolValueSlice)
	if queryErr != nil {
		return nil, queryErr
	}
	return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}

type boolQueryGorillaEncodeResponse struct {
	marshalOptions      protojson.MarshalOptions
	unmarshalOptions    protojson.UnmarshalOptions
	responseTransformer gorilla.ResponseTransformer
}

func (encoder boolQueryGorillaEncodeResponse) BoolQuery(ctx context.Context, w http.ResponseWriter, resp *httpbody.HttpBody) error {
	return gorilla.EncodeHttpBody(ctx, w, resp)
}

type Int32QueryGorillaService interface {
	Int32Query(ctx context.Context, request *Int32QueryRequest) (*httpbody.HttpBody, error)
}

func AppendInt32QueryGorillaRoute(router *mux.Router, service Int32QueryGorillaService, opts ...gorilla.Option) *mux.Router {
	options := gorilla.NewOptions(opts...)
	handler := int32QueryGorillaHandler{
		service: service,
		decoder: int32QueryGorillaRequestDecoder{
			unmarshalOptions:        options.UnmarshalOptions(),
			shouldFailFast:          options.ShouldFailFast(),
			onValidationErrCallback: options.OnValidationErrCallback(),
		},
		encoder: int32QueryGorillaEncodeResponse{
			marshalOptions:      options.MarshalOptions(),
			unmarshalOptions:    options.UnmarshalOptions(),
			responseTransformer: options.ResponseTransformer(),
		},
		errorEncoder: gorilla.DefaultEncodeError,
	}
	router.NewRoute().
		Name("/leo.gorilla.example.query.v1.Int32Query/Int32Query").
		Methods("GET").
		Path("/v1/int32").
		Handler(gorilla.Chain(handler.Int32Query(), options.Middlewares()...))
	return router
}

type int32QueryGorillaHandler struct {
	service      Int32QueryGorillaService
	decoder      int32QueryGorillaRequestDecoder
	encoder      int32QueryGorillaEncodeResponse
	errorEncoder gorilla.ErrorEncoder
}

func (h int32QueryGorillaHandler) Int32Query() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.Int32Query(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.Int32Query(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.Int32Query(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

type int32QueryGorillaRequestDecoder struct {
	unmarshalOptions        protojson.UnmarshalOptions
	shouldFailFast          bool
	onValidationErrCallback gorilla.OnValidationErrCallback
}

func (decoder int32QueryGorillaRequestDecoder) Int32Query(ctx context.Context, r *http.Request) (*Int32QueryRequest, error) {
	req := &Int32QueryRequest{}
	if ok, err := gorilla.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	queries := r.URL.Query()
	var queryErr error
	req.Int32, queryErr = gorilla.DecodeForm[int32](queryErr, queries, "int32", gorilla.GetInt32)
	req.Sint32, queryErr = gorilla.DecodeForm[int32](queryErr, queries, "sint32", gorilla.GetInt32)
	req.Sfixed32, queryErr = gorilla.DecodeForm[int32](queryErr, queries, "sfixed32", gorilla.GetInt32)
	req.OptInt32, queryErr = gorilla.DecodeForm[*int32](queryErr, queries, "opt_int32", gorilla.GetInt32Ptr)
	req.OptSint32, queryErr = gorilla.DecodeForm[*int32](queryErr, queries, "opt_sint32", gorilla.GetInt32Ptr)
	req.OptSfixed32, queryErr = gorilla.DecodeForm[*int32](queryErr, queries, "opt_sfixed32", gorilla.GetInt32Ptr)
	req.WrapInt32, queryErr = gorilla.DecodeForm[*wrapperspb.Int32Value](queryErr, queries, "wrap_int32", gorilla.GetInt32Value)
	req.ListInt32, queryErr = gorilla.DecodeForm[[]int32](queryErr, queries, "list_int32", gorilla.GetInt32Slice)
	req.ListSint32, queryErr = gorilla.DecodeForm[[]int32](queryErr, queries, "list_sint32", gorilla.GetInt32Slice)
	req.ListSfixed32, queryErr = gorilla.DecodeForm[[]int32](queryErr, queries, "list_sfixed32", gorilla.GetInt32Slice)
	req.ListWrapInt32, queryErr = gorilla.DecodeForm[[]*wrapperspb.Int32Value](queryErr, queries, "list_wrap_int32", gorilla.GetInt32ValueSlice)
	if queryErr != nil {
		return nil, queryErr
	}
	return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}

type int32QueryGorillaEncodeResponse struct {
	marshalOptions      protojson.MarshalOptions
	unmarshalOptions    protojson.UnmarshalOptions
	responseTransformer gorilla.ResponseTransformer
}

func (encoder int32QueryGorillaEncodeResponse) Int32Query(ctx context.Context, w http.ResponseWriter, resp *httpbody.HttpBody) error {
	return gorilla.EncodeHttpBody(ctx, w, resp)
}

type Int64QueryGorillaService interface {
	Int64Query(ctx context.Context, request *Int64QueryRequest) (*httpbody.HttpBody, error)
}

func AppendInt64QueryGorillaRoute(router *mux.Router, service Int64QueryGorillaService, opts ...gorilla.Option) *mux.Router {
	options := gorilla.NewOptions(opts...)
	handler := int64QueryGorillaHandler{
		service: service,
		decoder: int64QueryGorillaRequestDecoder{
			unmarshalOptions:        options.UnmarshalOptions(),
			shouldFailFast:          options.ShouldFailFast(),
			onValidationErrCallback: options.OnValidationErrCallback(),
		},
		encoder: int64QueryGorillaEncodeResponse{
			marshalOptions:      options.MarshalOptions(),
			unmarshalOptions:    options.UnmarshalOptions(),
			responseTransformer: options.ResponseTransformer(),
		},
		errorEncoder: gorilla.DefaultEncodeError,
	}
	router.NewRoute().
		Name("/leo.gorilla.example.query.v1.Int64Query/Int64Query").
		Methods("GET").
		Path("/v1/int64").
		Handler(gorilla.Chain(handler.Int64Query(), options.Middlewares()...))
	return router
}

type int64QueryGorillaHandler struct {
	service      Int64QueryGorillaService
	decoder      int64QueryGorillaRequestDecoder
	encoder      int64QueryGorillaEncodeResponse
	errorEncoder gorilla.ErrorEncoder
}

func (h int64QueryGorillaHandler) Int64Query() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.Int64Query(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.Int64Query(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.Int64Query(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

type int64QueryGorillaRequestDecoder struct {
	unmarshalOptions        protojson.UnmarshalOptions
	shouldFailFast          bool
	onValidationErrCallback gorilla.OnValidationErrCallback
}

func (decoder int64QueryGorillaRequestDecoder) Int64Query(ctx context.Context, r *http.Request) (*Int64QueryRequest, error) {
	req := &Int64QueryRequest{}
	if ok, err := gorilla.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	queries := r.URL.Query()
	var queryErr error
	req.Int64, queryErr = gorilla.DecodeForm[int64](queryErr, queries, "int64", gorilla.GetInt64)
	req.Sint64, queryErr = gorilla.DecodeForm[int64](queryErr, queries, "sint64", gorilla.GetInt64)
	req.Sfixed64, queryErr = gorilla.DecodeForm[int64](queryErr, queries, "sfixed64", gorilla.GetInt64)
	req.OptInt64, queryErr = gorilla.DecodeForm[*int64](queryErr, queries, "opt_int64", gorilla.GetInt64Ptr)
	req.OptSint64, queryErr = gorilla.DecodeForm[*int64](queryErr, queries, "opt_sint64", gorilla.GetInt64Ptr)
	req.OptSfixed64, queryErr = gorilla.DecodeForm[*int64](queryErr, queries, "opt_sfixed64", gorilla.GetInt64Ptr)
	req.WrapInt64, queryErr = gorilla.DecodeForm[*wrapperspb.Int64Value](queryErr, queries, "wrap_int64", gorilla.GetInt64Value)
	req.ListInt64, queryErr = gorilla.DecodeForm[[]int64](queryErr, queries, "list_int64", gorilla.GetInt64Slice)
	req.ListSint64, queryErr = gorilla.DecodeForm[[]int64](queryErr, queries, "list_sint64", gorilla.GetInt64Slice)
	req.ListSfixed64, queryErr = gorilla.DecodeForm[[]int64](queryErr, queries, "list_sfixed64", gorilla.GetInt64Slice)
	req.ListWrapInt64, queryErr = gorilla.DecodeForm[[]*wrapperspb.Int64Value](queryErr, queries, "list_wrap_int64", gorilla.GetInt64ValueSlice)
	if queryErr != nil {
		return nil, queryErr
	}
	return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}

type int64QueryGorillaEncodeResponse struct {
	marshalOptions      protojson.MarshalOptions
	unmarshalOptions    protojson.UnmarshalOptions
	responseTransformer gorilla.ResponseTransformer
}

func (encoder int64QueryGorillaEncodeResponse) Int64Query(ctx context.Context, w http.ResponseWriter, resp *httpbody.HttpBody) error {
	return gorilla.EncodeHttpBody(ctx, w, resp)
}

type Uint32QueryGorillaService interface {
	Uint32Query(ctx context.Context, request *Uint32QueryRequest) (*httpbody.HttpBody, error)
}

func AppendUint32QueryGorillaRoute(router *mux.Router, service Uint32QueryGorillaService, opts ...gorilla.Option) *mux.Router {
	options := gorilla.NewOptions(opts...)
	handler := uint32QueryGorillaHandler{
		service: service,
		decoder: uint32QueryGorillaRequestDecoder{
			unmarshalOptions:        options.UnmarshalOptions(),
			shouldFailFast:          options.ShouldFailFast(),
			onValidationErrCallback: options.OnValidationErrCallback(),
		},
		encoder: uint32QueryGorillaEncodeResponse{
			marshalOptions:      options.MarshalOptions(),
			unmarshalOptions:    options.UnmarshalOptions(),
			responseTransformer: options.ResponseTransformer(),
		},
		errorEncoder: gorilla.DefaultEncodeError,
	}
	router.NewRoute().
		Name("/leo.gorilla.example.query.v1.Uint32Query/Uint32Query").
		Methods("GET").
		Path("/v1/uint32").
		Handler(gorilla.Chain(handler.Uint32Query(), options.Middlewares()...))
	return router
}

type uint32QueryGorillaHandler struct {
	service      Uint32QueryGorillaService
	decoder      uint32QueryGorillaRequestDecoder
	encoder      uint32QueryGorillaEncodeResponse
	errorEncoder gorilla.ErrorEncoder
}

func (h uint32QueryGorillaHandler) Uint32Query() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.Uint32Query(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.Uint32Query(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.Uint32Query(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

type uint32QueryGorillaRequestDecoder struct {
	unmarshalOptions        protojson.UnmarshalOptions
	shouldFailFast          bool
	onValidationErrCallback gorilla.OnValidationErrCallback
}

func (decoder uint32QueryGorillaRequestDecoder) Uint32Query(ctx context.Context, r *http.Request) (*Uint32QueryRequest, error) {
	req := &Uint32QueryRequest{}
	if ok, err := gorilla.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	queries := r.URL.Query()
	var queryErr error
	req.Uint32, queryErr = gorilla.DecodeForm[uint32](queryErr, queries, "uint32", gorilla.GetUint32)
	req.Fixed32, queryErr = gorilla.DecodeForm[uint32](queryErr, queries, "fixed32", gorilla.GetUint32)
	req.OptUint32, queryErr = gorilla.DecodeForm[*uint32](queryErr, queries, "opt_uint32", gorilla.GetUint32Ptr)
	req.OptFixed32, queryErr = gorilla.DecodeForm[*uint32](queryErr, queries, "opt_fixed32", gorilla.GetUint32Ptr)
	req.WrapUint32, queryErr = gorilla.DecodeForm[*wrapperspb.UInt32Value](queryErr, queries, "wrap_uint32", gorilla.GetUint32Value)
	req.ListUint32, queryErr = gorilla.DecodeForm[[]uint32](queryErr, queries, "list_uint32", gorilla.GetUint32Slice)
	req.ListFixed32, queryErr = gorilla.DecodeForm[[]uint32](queryErr, queries, "list_fixed32", gorilla.GetUint32Slice)
	req.ListWrapUint32, queryErr = gorilla.DecodeForm[[]*wrapperspb.UInt32Value](queryErr, queries, "list_wrap_uint32", gorilla.GetUint32ValueSlice)
	if queryErr != nil {
		return nil, queryErr
	}
	return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}

type uint32QueryGorillaEncodeResponse struct {
	marshalOptions      protojson.MarshalOptions
	unmarshalOptions    protojson.UnmarshalOptions
	responseTransformer gorilla.ResponseTransformer
}

func (encoder uint32QueryGorillaEncodeResponse) Uint32Query(ctx context.Context, w http.ResponseWriter, resp *httpbody.HttpBody) error {
	return gorilla.EncodeHttpBody(ctx, w, resp)
}

type Uint64QueryGorillaService interface {
	Uint64Query(ctx context.Context, request *Uint64QueryRequest) (*httpbody.HttpBody, error)
}

func AppendUint64QueryGorillaRoute(router *mux.Router, service Uint64QueryGorillaService, opts ...gorilla.Option) *mux.Router {
	options := gorilla.NewOptions(opts...)
	handler := uint64QueryGorillaHandler{
		service: service,
		decoder: uint64QueryGorillaRequestDecoder{
			unmarshalOptions:        options.UnmarshalOptions(),
			shouldFailFast:          options.ShouldFailFast(),
			onValidationErrCallback: options.OnValidationErrCallback(),
		},
		encoder: uint64QueryGorillaEncodeResponse{
			marshalOptions:      options.MarshalOptions(),
			unmarshalOptions:    options.UnmarshalOptions(),
			responseTransformer: options.ResponseTransformer(),
		},
		errorEncoder: gorilla.DefaultEncodeError,
	}
	router.NewRoute().
		Name("/leo.gorilla.example.query.v1.Uint64Query/Uint64Query").
		Methods("GET").
		Path("/v1/uint64").
		Handler(gorilla.Chain(handler.Uint64Query(), options.Middlewares()...))
	return router
}

type uint64QueryGorillaHandler struct {
	service      Uint64QueryGorillaService
	decoder      uint64QueryGorillaRequestDecoder
	encoder      uint64QueryGorillaEncodeResponse
	errorEncoder gorilla.ErrorEncoder
}

func (h uint64QueryGorillaHandler) Uint64Query() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.Uint64Query(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.Uint64Query(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.Uint64Query(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

type uint64QueryGorillaRequestDecoder struct {
	unmarshalOptions        protojson.UnmarshalOptions
	shouldFailFast          bool
	onValidationErrCallback gorilla.OnValidationErrCallback
}

func (decoder uint64QueryGorillaRequestDecoder) Uint64Query(ctx context.Context, r *http.Request) (*Uint64QueryRequest, error) {
	req := &Uint64QueryRequest{}
	if ok, err := gorilla.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	queries := r.URL.Query()
	var queryErr error
	req.Uint64, queryErr = gorilla.DecodeForm[uint64](queryErr, queries, "uint64", gorilla.GetUint64)
	req.Fixed64, queryErr = gorilla.DecodeForm[uint64](queryErr, queries, "fixed64", gorilla.GetUint64)
	req.OptUint64, queryErr = gorilla.DecodeForm[*uint64](queryErr, queries, "opt_uint64", gorilla.GetUint64Ptr)
	req.OptFixed64, queryErr = gorilla.DecodeForm[*uint64](queryErr, queries, "opt_fixed64", gorilla.GetUint64Ptr)
	req.WrapUint64, queryErr = gorilla.DecodeForm[*wrapperspb.UInt64Value](queryErr, queries, "wrap_uint64", gorilla.GetUint64Value)
	req.ListUint64, queryErr = gorilla.DecodeForm[[]uint64](queryErr, queries, "list_uint64", gorilla.GetUint64Slice)
	req.ListFixed64, queryErr = gorilla.DecodeForm[[]uint64](queryErr, queries, "list_fixed64", gorilla.GetUint64Slice)
	req.ListWrapUint64, queryErr = gorilla.DecodeForm[[]*wrapperspb.UInt64Value](queryErr, queries, "list_wrap_uint64", gorilla.GetUint64ValueSlice)
	if queryErr != nil {
		return nil, queryErr
	}
	return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}

type uint64QueryGorillaEncodeResponse struct {
	marshalOptions      protojson.MarshalOptions
	unmarshalOptions    protojson.UnmarshalOptions
	responseTransformer gorilla.ResponseTransformer
}

func (encoder uint64QueryGorillaEncodeResponse) Uint64Query(ctx context.Context, w http.ResponseWriter, resp *httpbody.HttpBody) error {
	return gorilla.EncodeHttpBody(ctx, w, resp)
}

type FloatQueryGorillaService interface {
	FloatQuery(ctx context.Context, request *FloatQueryRequest) (*httpbody.HttpBody, error)
}

func AppendFloatQueryGorillaRoute(router *mux.Router, service FloatQueryGorillaService, opts ...gorilla.Option) *mux.Router {
	options := gorilla.NewOptions(opts...)
	handler := floatQueryGorillaHandler{
		service: service,
		decoder: floatQueryGorillaRequestDecoder{
			unmarshalOptions:        options.UnmarshalOptions(),
			shouldFailFast:          options.ShouldFailFast(),
			onValidationErrCallback: options.OnValidationErrCallback(),
		},
		encoder: floatQueryGorillaEncodeResponse{
			marshalOptions:      options.MarshalOptions(),
			unmarshalOptions:    options.UnmarshalOptions(),
			responseTransformer: options.ResponseTransformer(),
		},
		errorEncoder: gorilla.DefaultEncodeError,
	}
	router.NewRoute().
		Name("/leo.gorilla.example.query.v1.FloatQuery/FloatQuery").
		Methods("GET").
		Path("/v1/float").
		Handler(gorilla.Chain(handler.FloatQuery(), options.Middlewares()...))
	return router
}

type floatQueryGorillaHandler struct {
	service      FloatQueryGorillaService
	decoder      floatQueryGorillaRequestDecoder
	encoder      floatQueryGorillaEncodeResponse
	errorEncoder gorilla.ErrorEncoder
}

func (h floatQueryGorillaHandler) FloatQuery() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.FloatQuery(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.FloatQuery(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.FloatQuery(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

type floatQueryGorillaRequestDecoder struct {
	unmarshalOptions        protojson.UnmarshalOptions
	shouldFailFast          bool
	onValidationErrCallback gorilla.OnValidationErrCallback
}

func (decoder floatQueryGorillaRequestDecoder) FloatQuery(ctx context.Context, r *http.Request) (*FloatQueryRequest, error) {
	req := &FloatQueryRequest{}
	if ok, err := gorilla.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	queries := r.URL.Query()
	var queryErr error
	req.Float, queryErr = gorilla.DecodeForm[float32](queryErr, queries, "float", gorilla.GetFloat32)
	req.OptFloat, queryErr = gorilla.DecodeForm[*float32](queryErr, queries, "opt_float", gorilla.GetFloat32Ptr)
	req.WrapFloat, queryErr = gorilla.DecodeForm[*wrapperspb.FloatValue](queryErr, queries, "wrap_float", gorilla.GetFloat32Value)
	req.ListFloat, queryErr = gorilla.DecodeForm[[]float32](queryErr, queries, "list_float", gorilla.GetFloat32Slice)
	req.ListWrapFloat, queryErr = gorilla.DecodeForm[[]*wrapperspb.FloatValue](queryErr, queries, "list_wrap_float", gorilla.GetFloat32ValueSlice)
	if queryErr != nil {
		return nil, queryErr
	}
	return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}

type floatQueryGorillaEncodeResponse struct {
	marshalOptions      protojson.MarshalOptions
	unmarshalOptions    protojson.UnmarshalOptions
	responseTransformer gorilla.ResponseTransformer
}

func (encoder floatQueryGorillaEncodeResponse) FloatQuery(ctx context.Context, w http.ResponseWriter, resp *httpbody.HttpBody) error {
	return gorilla.EncodeHttpBody(ctx, w, resp)
}

type DoubleQueryGorillaService interface {
	DoubleQuery(ctx context.Context, request *DoubleQueryRequest) (*httpbody.HttpBody, error)
}

func AppendDoubleQueryGorillaRoute(router *mux.Router, service DoubleQueryGorillaService, opts ...gorilla.Option) *mux.Router {
	options := gorilla.NewOptions(opts...)
	handler := doubleQueryGorillaHandler{
		service: service,
		decoder: doubleQueryGorillaRequestDecoder{
			unmarshalOptions:        options.UnmarshalOptions(),
			shouldFailFast:          options.ShouldFailFast(),
			onValidationErrCallback: options.OnValidationErrCallback(),
		},
		encoder: doubleQueryGorillaEncodeResponse{
			marshalOptions:      options.MarshalOptions(),
			unmarshalOptions:    options.UnmarshalOptions(),
			responseTransformer: options.ResponseTransformer(),
		},
		errorEncoder: gorilla.DefaultEncodeError,
	}
	router.NewRoute().
		Name("/leo.gorilla.example.query.v1.DoubleQuery/DoubleQuery").
		Methods("GET").
		Path("/v1/double").
		Handler(gorilla.Chain(handler.DoubleQuery(), options.Middlewares()...))
	return router
}

type doubleQueryGorillaHandler struct {
	service      DoubleQueryGorillaService
	decoder      doubleQueryGorillaRequestDecoder
	encoder      doubleQueryGorillaEncodeResponse
	errorEncoder gorilla.ErrorEncoder
}

func (h doubleQueryGorillaHandler) DoubleQuery() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.DoubleQuery(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.DoubleQuery(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.DoubleQuery(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

type doubleQueryGorillaRequestDecoder struct {
	unmarshalOptions        protojson.UnmarshalOptions
	shouldFailFast          bool
	onValidationErrCallback gorilla.OnValidationErrCallback
}

func (decoder doubleQueryGorillaRequestDecoder) DoubleQuery(ctx context.Context, r *http.Request) (*DoubleQueryRequest, error) {
	req := &DoubleQueryRequest{}
	if ok, err := gorilla.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	queries := r.URL.Query()
	var queryErr error
	req.Double, queryErr = gorilla.DecodeForm[float64](queryErr, queries, "double", gorilla.GetFloat64)
	req.OptDouble, queryErr = gorilla.DecodeForm[*float64](queryErr, queries, "opt_double", gorilla.GetFloat64Ptr)
	req.WrapDouble, queryErr = gorilla.DecodeForm[*wrapperspb.DoubleValue](queryErr, queries, "wrap_double", gorilla.GetFloat64Value)
	req.ListDouble, queryErr = gorilla.DecodeForm[[]float64](queryErr, queries, "list_double", gorilla.GetFloat64Slice)
	req.ListWrapDouble, queryErr = gorilla.DecodeForm[[]*wrapperspb.DoubleValue](queryErr, queries, "list_wrap_double", gorilla.GetFloat64ValueSlice)
	if queryErr != nil {
		return nil, queryErr
	}
	return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}

type doubleQueryGorillaEncodeResponse struct {
	marshalOptions      protojson.MarshalOptions
	unmarshalOptions    protojson.UnmarshalOptions
	responseTransformer gorilla.ResponseTransformer
}

func (encoder doubleQueryGorillaEncodeResponse) DoubleQuery(ctx context.Context, w http.ResponseWriter, resp *httpbody.HttpBody) error {
	return gorilla.EncodeHttpBody(ctx, w, resp)
}

type StringQueryGorillaService interface {
	StringQuery(ctx context.Context, request *StringQueryRequest) (*httpbody.HttpBody, error)
}

func AppendStringQueryGorillaRoute(router *mux.Router, service StringQueryGorillaService, opts ...gorilla.Option) *mux.Router {
	options := gorilla.NewOptions(opts...)
	handler := stringQueryGorillaHandler{
		service: service,
		decoder: stringQueryGorillaRequestDecoder{
			unmarshalOptions:        options.UnmarshalOptions(),
			shouldFailFast:          options.ShouldFailFast(),
			onValidationErrCallback: options.OnValidationErrCallback(),
		},
		encoder: stringQueryGorillaEncodeResponse{
			marshalOptions:      options.MarshalOptions(),
			unmarshalOptions:    options.UnmarshalOptions(),
			responseTransformer: options.ResponseTransformer(),
		},
		errorEncoder: gorilla.DefaultEncodeError,
	}
	router.NewRoute().
		Name("/leo.gorilla.example.query.v1.StringQuery/StringQuery").
		Methods("GET").
		Path("/v1/string").
		Handler(gorilla.Chain(handler.StringQuery(), options.Middlewares()...))
	return router
}

type stringQueryGorillaHandler struct {
	service      StringQueryGorillaService
	decoder      stringQueryGorillaRequestDecoder
	encoder      stringQueryGorillaEncodeResponse
	errorEncoder gorilla.ErrorEncoder
}

func (h stringQueryGorillaHandler) StringQuery() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.StringQuery(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.StringQuery(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.StringQuery(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

type stringQueryGorillaRequestDecoder struct {
	unmarshalOptions        protojson.UnmarshalOptions
	shouldFailFast          bool
	onValidationErrCallback gorilla.OnValidationErrCallback
}

func (decoder stringQueryGorillaRequestDecoder) StringQuery(ctx context.Context, r *http.Request) (*StringQueryRequest, error) {
	req := &StringQueryRequest{}
	if ok, err := gorilla.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	queries := r.URL.Query()
	var queryErr error
	req.String_ = queries.Get("string")
	req.OptString = proto.String(queries.Get("opt_string"))
	req.WrapString = wrapperspb.String(queries.Get("wrap_string"))
	req.ListString = queries["list_string"]
	req.ListWrapString = gorilla.WrapStringSlice(queries["list_wrap_string"])
	if queryErr != nil {
		return nil, queryErr
	}
	return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}

type stringQueryGorillaEncodeResponse struct {
	marshalOptions      protojson.MarshalOptions
	unmarshalOptions    protojson.UnmarshalOptions
	responseTransformer gorilla.ResponseTransformer
}

func (encoder stringQueryGorillaEncodeResponse) StringQuery(ctx context.Context, w http.ResponseWriter, resp *httpbody.HttpBody) error {
	return gorilla.EncodeHttpBody(ctx, w, resp)
}

type EnumQueryGorillaService interface {
	EnumQuery(ctx context.Context, request *EnumQueryRequest) (*httpbody.HttpBody, error)
}

func AppendEnumQueryGorillaRoute(router *mux.Router, service EnumQueryGorillaService, opts ...gorilla.Option) *mux.Router {
	options := gorilla.NewOptions(opts...)
	handler := enumQueryGorillaHandler{
		service: service,
		decoder: enumQueryGorillaRequestDecoder{
			unmarshalOptions:        options.UnmarshalOptions(),
			shouldFailFast:          options.ShouldFailFast(),
			onValidationErrCallback: options.OnValidationErrCallback(),
		},
		encoder: enumQueryGorillaEncodeResponse{
			marshalOptions:      options.MarshalOptions(),
			unmarshalOptions:    options.UnmarshalOptions(),
			responseTransformer: options.ResponseTransformer(),
		},
		errorEncoder: gorilla.DefaultEncodeError,
	}
	router.NewRoute().
		Name("/leo.gorilla.example.query.v1.EnumQuery/EnumQuery").
		Methods("GET").
		Path("/v1/enum").
		Handler(gorilla.Chain(handler.EnumQuery(), options.Middlewares()...))
	return router
}

type enumQueryGorillaHandler struct {
	service      EnumQueryGorillaService
	decoder      enumQueryGorillaRequestDecoder
	encoder      enumQueryGorillaEncodeResponse
	errorEncoder gorilla.ErrorEncoder
}

func (h enumQueryGorillaHandler) EnumQuery() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		in, err := h.decoder.EnumQuery(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.EnumQuery(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.EnumQuery(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

type enumQueryGorillaRequestDecoder struct {
	unmarshalOptions        protojson.UnmarshalOptions
	shouldFailFast          bool
	onValidationErrCallback gorilla.OnValidationErrCallback
}

func (decoder enumQueryGorillaRequestDecoder) EnumQuery(ctx context.Context, r *http.Request) (*EnumQueryRequest, error) {
	req := &EnumQueryRequest{}
	if ok, err := gorilla.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	queries := r.URL.Query()
	var queryErr error
	req.Status, queryErr = gorilla.DecodeForm[EnumQueryRequest_Status](queryErr, queries, "status", gorilla.GetInt[EnumQueryRequest_Status])
	req.OptStatus, queryErr = gorilla.DecodeForm[*EnumQueryRequest_Status](queryErr, queries, "opt_status", gorilla.GetIntPtr[EnumQueryRequest_Status])
	req.ListStatus, queryErr = gorilla.DecodeForm[[]EnumQueryRequest_Status](queryErr, queries, "list_status", gorilla.GetIntSlice[EnumQueryRequest_Status])
	if queryErr != nil {
		return nil, queryErr
	}
	return req, gorilla.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}

type enumQueryGorillaEncodeResponse struct {
	marshalOptions      protojson.MarshalOptions
	unmarshalOptions    protojson.UnmarshalOptions
	responseTransformer gorilla.ResponseTransformer
}

func (encoder enumQueryGorillaEncodeResponse) EnumQuery(ctx context.Context, w http.ResponseWriter, resp *httpbody.HttpBody) error {
	return gorilla.EncodeHttpBody(ctx, w, resp)
}
