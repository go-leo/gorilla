package gen

import (
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	ProtoJsonPackage               = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	ProtoJsonMarshalOptionsIdent   = ProtoJsonPackage.Ident("MarshalOptions")
	ProtoJsonUnmarshalOptionsIdent = ProtoJsonPackage.Ident("UnmarshalOptions")
)

var (
	ContextPackage = protogen.GoImportPath("context")
	ContextIdent   = ContextPackage.Ident("Context")
)

var (
	HttpPackage          = protogen.GoImportPath("net/http")
	ClientIdent          = HttpPackage.Ident("Client")
	HttpHandlerIdent     = HttpPackage.Ident("Handler")
	HttpHandlerFuncIdent = HttpPackage.Ident("HandlerFunc")
	ResponseWriterIdent  = HttpPackage.Ident("ResponseWriter")
	RequestIdent         = HttpPackage.Ident("Request")
	ResponseIdent        = HttpPackage.Ident("Response")
)

var (
	FmtPackage   = protogen.GoImportPath("fmt")
	SprintfIdent = FmtPackage.Ident("Sprintf")
)

var (
	MuxPackage  = protogen.GoImportPath("github.com/gorilla/mux")
	RouterIdent = MuxPackage.Ident("Router")
	VarsIdent   = MuxPackage.Ident("Vars")
)

var (
	ProtoPackage     = protogen.GoImportPath("google.golang.org/protobuf/proto")
	ProtoStringIdent = ProtoPackage.Ident("String")
)

var (
	GorillaPackage               = protogen.GoImportPath("github.com/go-leo/gorilla")
	ErrorEncoderIdent            = GorillaPackage.Ident("ErrorEncoder")
	ResponseTransformerIdent     = GorillaPackage.Ident("ResponseTransformer")
	DefaultEncodeErrorIdent      = GorillaPackage.Ident("DefaultEncodeError")
	EncodeResponseIdent          = GorillaPackage.Ident("EncodeResponse")
	EncodeHttpBodyIdent          = GorillaPackage.Ident("EncodeHttpBody")
	EncodeHttpResponseIdent      = GorillaPackage.Ident("EncodeHttpResponse")
	DecodeRequestIdent           = GorillaPackage.Ident("DecodeRequest")
	DecodeHttpBodyIdent          = GorillaPackage.Ident("DecodeHttpBody")
	DecodeHttpRequestIdent       = GorillaPackage.Ident("DecodeHttpRequest")
	DecodeFormIdent              = GorillaPackage.Ident("DecodeForm")
	OptionIdent                  = GorillaPackage.Ident("Option")
	NewOptionsIdent              = GorillaPackage.Ident("NewOptions")
	ChainIdent                   = GorillaPackage.Ident("Chain")
	CustomDecodeRequestIdent     = GorillaPackage.Ident("CustomDecodeRequest")
	OnValidationErrCallbackIdent = GorillaPackage.Ident("OnValidationErrCallback")
	ValidateRequestIdent         = GorillaPackage.Ident("ValidateRequest")
)

var (
	WrapperspbPackage     = protogen.GoImportPath("google.golang.org/protobuf/types/known/wrapperspb")
	WrapperspbStringIdent = WrapperspbPackage.Ident("String")
)
