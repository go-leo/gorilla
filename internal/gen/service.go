package gen

import (
	"fmt"
	"strings"

	"github.com/gorilla/mux"
	"google.golang.org/protobuf/compiler/protogen"
)

type Service struct {
	ProtoService *protogen.Service
	Endpoints    []*Endpoint
}

func (s *Service) Unexported(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func (s *Service) FullName() string {
	return string(s.ProtoService.Desc.FullName())
}

func (s *Service) Name() string {
	return s.ProtoService.GoName
}

func (s *Service) GorillaName() string {
	return s.Name() + "Gorilla"
}

func (s *Service) ServiceName() string {
	return s.GorillaName() + "Service"
}

func (s *Service) AppendRouteName() string {
	return "Append" + s.GorillaName() + "Route"
}

func (s *Service) HandlerName() string {
	return s.GorillaName() + "Handler"
}

func (s *Service) RequestDecoderName() string {
	return s.GorillaName() + "RequestDecoder"
}

func (s *Service) EncodeResponseName() string {
	return s.GorillaName() + "EncodeResponse"
}

func NewServices(file *protogen.File) ([]*Service, error) {
	var services []*Service
	for _, pbService := range file.Services {
		service := &Service{
			ProtoService: pbService,
		}
		var endpoints []*Endpoint
		router := mux.NewRouter()
		for _, pbMethod := range pbService.Methods {
			endpoint := &Endpoint{
				protoMethod: pbMethod,
			}
			if endpoint.IsStreaming() {
				return nil, fmt.Errorf("gorilla: unsupport stream method, %s", endpoint.FullName())
			}
			endpoint.SetHttpRule()
			route := router.NewRoute()
			route.Methods(endpoint.Method()).Path(endpoint.Path())
			if err := checkRoute(route, endpoint); err != nil {
				return nil, fmt.Errorf("gorilla: %s", err)
			}
			endpoint.SetRoute(route)
			endpoints = append(endpoints, endpoint)
		}
		service.Endpoints = endpoints
		services = append(services, service)
	}
	return services, nil
}

func checkRoute(route *mux.Route, endpoint *Endpoint) error {
	route.Methods(endpoint.Method()).Path(endpoint.Path())
	return route.GetError()
}
