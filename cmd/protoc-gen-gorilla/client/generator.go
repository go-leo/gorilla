package client

import (
	"github.com/go-leo/gorilla/internal/gen"
	"google.golang.org/protobuf/compiler/protogen"
)

type Generator struct{}

func (f *Generator) GenerateClient(service *gen.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.Unexported(service.ClientName()), " struct {")
	g.P("client *", gen.ClientIdent)
	g.P("encoder ", service.Unexported(service.RequestEncoderName()))
	g.P("decoder ", service.Unexported(service.ResponseDecoderName()))
	g.P("shouldFailFast bool")
	g.P("onValidationErrCallback ", gen.OnValidationErrCallbackIdent)
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (c *", service.Unexported(service.ClientName()), ") ", endpoint.Name(), "(ctx ", gen.ContextIdent, ", in *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
		g.P("if err := ", gen.ValidateRequestIdent, "(ctx, in, c.shouldFailFast, c.onValidationErrCallback); err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("req, err := c.encoder.", endpoint.Name(), "(ctx, in)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("resp, err := c.client.Do(req)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("out, err := c.decoder.", endpoint.Name(), "(ctx, resp)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("return out, nil")
		g.P("}")
		g.P()
	}
	return nil
}

// func (f *Generator) GenerateNewClientFunc() {
// 	f.g.P("func New", f.service.HttpClientName(), "(target string, opts ...", internal.HttpxTransportxPackage.Ident("ClientOption"), ") ", f.service.ServiceName(), " {")
// 	f.g.P("options := ", internal.HttpxTransportxPackage.Ident("NewClientOptions"), "(opts...)")

// 	f.g.P("requestEncoder := &", f.service.Unexported(f.service.HttpClientRequestEncoderName()), "{")
// 	f.g.P("marshalOptions: options.MarshalOptions(),")
// 	f.g.P("router: append", f.service.HttpRoutesName(), "(", internal.MuxPackage.Ident("NewRouter"), "()),")
// 	f.g.P("scheme: options.Scheme(),")
// 	f.g.P("}")
// 	f.g.P("responseDecoder := &", f.service.Unexported(f.service.HttpClientResponseDecoderName()), "{")
// 	f.g.P("unmarshalOptions: options.UnmarshalOptions(),")
// 	f.g.P("}")
// 	f.g.P("transports :=  &", f.service.Unexported(f.service.HttpClientTransportsName()), "{")
// 	f.g.P("clientOptions: options.ClientTransportOptions(),")
// 	f.g.P("middlewares:   options.Middlewares(),")
// 	f.g.P("requestEncoder:  requestEncoder,")
// 	f.g.P("responseDecoder: responseDecoder,")
// 	f.g.P("}")
// 	f.g.P("factories := &", f.service.Unexported(f.service.FactoriesName()), "{")
// 	f.g.P("transports: transports,")
// 	f.g.P("}")
// 	f.g.P("endpointer := &", f.service.Unexported(f.service.EndpointersName()), "{")
// 	f.g.P("target:    target,")
// 	f.g.P("builder:   options.Builder(),")
// 	f.g.P("factories: factories,")
// 	f.g.P("logger:    options.Logger(),")
// 	f.g.P("options:   options.EndpointerOptions(),")
// 	f.g.P("}")
// 	f.g.P("balancers := &", f.service.Unexported(f.service.BalancersName()), "{")
// 	f.g.P("factory:    options.BalancerFactory(),")
// 	f.g.P("endpointer: endpointer,")
// 	f.g.P("}")
// 	f.g.P("endpoints := &", f.service.Unexported(f.service.ClientEndpointsName()), "{")
// 	f.g.P("balancers: balancers,")
// 	f.g.P("}")

// 	f.g.P("return &", f.service.Unexported(f.service.ClientServiceName()), "{")
// 	f.g.P("endpoints:     endpoints,")
// 	f.g.P("transportName: ", internal.HttpxTransportxPackage.Ident("HttpClient"), ",")
// 	f.g.P("}")
// 	f.g.P("}")
// 	f.g.P()
// }
