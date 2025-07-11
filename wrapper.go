package gorilla

import "google.golang.org/protobuf/types/known/wrapperspb"

func WrapInt32Slice(s []int32) []*wrapperspb.Int32Value {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.Int32Value, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Int32(v))
	}
	return r
}

func WrapInt64Slice(s []int64) []*wrapperspb.Int64Value {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.Int64Value, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Int64(v))
	}
	return r
}

func WrapBoolSlice(s []bool) []*wrapperspb.BoolValue {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.BoolValue, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Bool(v))
	}
	return r
}

func WrapUint32Slice(s []uint32) []*wrapperspb.UInt32Value {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.UInt32Value, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.UInt32(v))
	}
	return r
}

func WrapUint64Slice(s []uint64) []*wrapperspb.UInt64Value {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.UInt64Value, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.UInt64(v))
	}
	return r
}

func WrapFloat32Slice(s []float32) []*wrapperspb.FloatValue {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.FloatValue, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Float(v))
	}
	return r
}

func WrapFloat64Slice(s []float64) []*wrapperspb.DoubleValue {
	if s == nil {
		return nil
	}
	r := make([]*wrapperspb.DoubleValue, 0, len(s))
	for _, v := range s {
		r = append(r, wrapperspb.Double(v))
	}
	return r
}
