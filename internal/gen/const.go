package gen

import (
	"regexp"

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

	HttpPackage          = protogen.GoImportPath("net/http")
	HttpHandlerIdent     = HttpPackage.Ident("Handler")
	HttpHandlerFuncIdent = HttpPackage.Ident("HandlerFunc")
	ResponseWriterIdent  = HttpPackage.Ident("ResponseWriter")
	RequestIdent         = HttpPackage.Ident("Request")

	FmtPackage   = protogen.GoImportPath("fmt")
	SprintfIdent = FmtPackage.Ident("Sprintf")

	MuxPackage  = protogen.GoImportPath("github.com/gorilla/mux")
	RouterIdent = MuxPackage.Ident("Router")
	VarsIdent   = MuxPackage.Ident("Vars")

	ProtoxPackage        = protogen.GoImportPath("github.com/go-leo/gox/protox")
	WrapStringSliceIdent = ProtoxPackage.Ident("WrapStringSlice")

	ProtoPackage     = protogen.GoImportPath("google.golang.org/protobuf/proto")
	ProtoStringIdent = ProtoPackage.Ident("String")

	GorillaPackage           = protogen.GoImportPath("github.com/go-leo/gorilla")
	ErrorEncoderIdent        = GorillaPackage.Ident("ErrorEncoder")
	ResponseTransformerIdent = GorillaPackage.Ident("ResponseTransformer")
	DefaultErrorEncoderIdent = GorillaPackage.Ident("DefaultErrorEncoder")
	ResponseEncoderIdent     = GorillaPackage.Ident("ResponseEncoder")
	HttpBodyEncoderIdent     = GorillaPackage.Ident("HttpBodyEncoder")
	HttpResponseEncoderIdent = GorillaPackage.Ident("HttpResponseEncoder")
	RequestDecoderIdent      = GorillaPackage.Ident("RequestDecoder")
	HttpBodyDecoderIdent     = GorillaPackage.Ident("HttpBodyDecoder")
	HttpRequestDecoderIdent  = GorillaPackage.Ident("HttpRequestDecoder")
	FormDecoderIdent         = GorillaPackage.Ident("FormDecoder")
	OptionIdent              = GorillaPackage.Ident("Option")
	NewOptionsIdent          = GorillaPackage.Ident("NewOptions")

	WrapperspbPackage     = protogen.GoImportPath("google.golang.org/protobuf/types/known/wrapperspb")
	WrapperspbStringIdent = WrapperspbPackage.Ident("String")
)

var (
	namedPathPattern = regexp.MustCompile("{([^{}]+)=([^{}]+)}")
	pathPattern      = regexp.MustCompile("{([^=}]+)}")
)
