package gorilla

import (
	"google.golang.org/protobuf/encoding/protojson"
)

// Options defines the interface for getting configuration options including:
type Options interface {
	// UnmarshalOptions Protobuf JSON unmarshaling options
	UnmarshalOptions() protojson.UnmarshalOptions
	// MarshalOptions Protobuf JSON marshaling options
	MarshalOptions() protojson.MarshalOptions
	// ErrorEncoder Error encoding handler
	ErrorEncoder() ErrorEncoder
	// ResponseTransformer Response transformation handler
	ResponseTransformer() ResponseTransformer
}

// options implements the Options interface with concrete configuration values
type options struct {
	// unmarshalOptions Protobuf JSON unmarshaling options
	unmarshalOptions protojson.UnmarshalOptions
	// marshalOptions Protobuf JSON marshaling options
	marshalOptions protojson.MarshalOptions
	// errorEncoder Error encoding handler
	errorEncoder ErrorEncoder
	// responseTransformer Response transformation handler
	responseTransformer ResponseTransformer
}

// Option defines a function type for modifying options
type Option func(o *options)

// Apply applies multiple Option functions to the options
func (o *options) Apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// UnmarshalOptions returns the configured protobuf JSON unmarshal options
func (o *options) UnmarshalOptions() protojson.UnmarshalOptions {
	return o.unmarshalOptions
}

// MarshalOptions returns the configured protobuf JSON marshal options
func (o *options) MarshalOptions() protojson.MarshalOptions {
	return o.marshalOptions
}

// ErrorEncoder returns the configured error encoder
func (o *options) ErrorEncoder() ErrorEncoder {
	return o.errorEncoder
}

// ResponseTransformer returns the configured response transformer
func (o *options) ResponseTransformer() ResponseTransformer {
	return o.responseTransformer
}

// WithUnmarshalOptions sets the protobuf JSON unmarshal options
func WithUnmarshalOptions(opts protojson.UnmarshalOptions) Option {
	return func(o *options) {
		o.unmarshalOptions = opts
	}
}

// WithMarshalOptions sets the protobuf JSON marshal options
func WithMarshalOptions(opts protojson.MarshalOptions) Option {
	return func(o *options) {
		o.marshalOptions = opts
	}
}

// WithErrorEncoder sets the error encoder
func WithErrorEncoder(encoder ErrorEncoder) Option {
	return func(o *options) {
		o.errorEncoder = encoder
	}
}

// WithResponseTransformer sets the response transformer
func WithResponseTransformer(transformer ResponseTransformer) Option {
	return func(o *options) {
		o.responseTransformer = transformer
	}
}

// NewOptions creates a new Options instance with default values that can be
// customized using the provided Option functions
func NewOptions(opts ...Option) Options {
	o := &options{
		unmarshalOptions:    protojson.UnmarshalOptions{},
		marshalOptions:      protojson.MarshalOptions{},
		errorEncoder:        DefaultErrorEncoder,
		responseTransformer: DefaultResponseTransformer,
	}
	o = o.Apply(opts...)
	return o
}
