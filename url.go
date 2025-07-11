package gorilla

import (
	"net/url"
	"strconv"

	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func GetInt32(form url.Values, key string) (int32, error) {
	return GetInt[int32](form, key)
}

func GetInt32Ptr(form url.Values, key string) (*int32, error) {
	return GetIntPtr[int32](form, key)
}

func GetInt64(form url.Values, key string) (int64, error) {
	return GetInt[int64](form, key)
}

func GetInt64Ptr(form url.Values, key string) (*int64, error) {
	return GetIntPtr[int64](form, key)
}

func GetUint32(form url.Values, key string) (uint32, error) {
	return GetUint[uint32](form, key)
}

func GetUint32Ptr(form url.Values, key string) (*uint32, error) {
	return GetUintPtr[uint32](form, key)
}

func GetUint64(form url.Values, key string) (uint64, error) {
	return GetUint[uint64](form, key)
}

func GetUint64Ptr(form url.Values, key string) (*uint64, error) {
	return GetUintPtr[uint64](form, key)
}

func GetFloat32(form url.Values, key string) (float32, error) {
	return GetFloat[float32](form, key)
}

func GetFloat32Ptr(form url.Values, key string) (*float32, error) {
	return GetFloatPtr[float32](form, key)
}

func GetFloat32Slice(form url.Values, key string) ([]float32, error) {
	return GetFloatSlice[float32](form, key)
}

func GetFloat64(form url.Values, key string) (float64, error) {
	return GetFloat[float64](form, key)
}

func GetFloat64Ptr(form url.Values, key string) (*float64, error) {
	return GetFloatPtr[float64](form, key)
}

func GetFloat64Slice(form url.Values, key string) ([]float64, error) {
	return GetFloatSlice[float64](form, key)
}

func GetBool(form url.Values, key string) (bool, error) {
	if _, ok := form[key]; !ok {
		return false, nil
	}
	return strconv.ParseBool(form.Get(key))
}

func GetBoolPtr(form url.Values, key string) (*bool, error) {
	v, err := GetBool(form, key)
	return &v, err
}

func GetBoolSlice(form url.Values, key string) ([]bool, error) {
	if _, ok := form[key]; !ok {
		return nil, nil
	}
	return ParseBoolSlice(form[key])
}

func GetBoolValue(form url.Values, key string) (*wrapperspb.BoolValue, error) {
	v, err := strconv.ParseBool(form.Get(key))
	return wrapperspb.Bool(v), err
}

func GetBoolValueSlice(form url.Values, key string) ([]*wrapperspb.BoolValue, error) {
	v, err := ParseBoolSlice(form[key])
	return WrapBoolSlice(v), err
}

func GetInt[Signed constraints.Signed](form url.Values, key string) (Signed, error) {
	if _, ok := form[key]; !ok {
		var v Signed
		return v, nil
	}
	return ParseInt[Signed](form.Get(key), 10, 64)
}

func GetIntPtr[Signed constraints.Signed](form url.Values, key string) (*Signed, error) {
	v, err := GetInt[Signed](form, key)
	return &v, err
}

func GetIntSlice[Signed constraints.Signed](form url.Values, key string) ([]Signed, error) {
	if _, ok := form[key]; !ok {
		var v []Signed
		return v, nil
	}
	return ParseIntSlice[Signed](form[key], 10, 64)
}

func GetInt32Value(form url.Values, key string) (*wrapperspb.Int32Value, error) {
	v, err := GetInt[int32](form, key)
	return wrapperspb.Int32(v), err
}

func GetInt32ValueSlice(form url.Values, key string) ([]*wrapperspb.Int32Value, error) {
	v, err := GetIntSlice[int32](form, key)
	return WrapInt32Slice(v), err
}

func GetInt64Value(form url.Values, key string) (*wrapperspb.Int64Value, error) {
	v, err := GetInt[int64](form, key)
	return wrapperspb.Int64(v), err
}

func GetInt64ValueSlice(form url.Values, key string) ([]*wrapperspb.Int64Value, error) {
	v, err := GetIntSlice[int64](form, key)
	return WrapInt64Slice(v), err
}

func GetUint[Unsigned constraints.Unsigned](form url.Values, key string) (Unsigned, error) {
	if _, ok := form[key]; !ok {
		var v Unsigned
		return v, nil
	}
	return ParseUint[Unsigned](form.Get(key), 10, 64)
}

func GetUintPtr[Unsigned constraints.Unsigned](form url.Values, key string) (*Unsigned, error) {
	v, err := GetUint[Unsigned](form, key)
	return &v, err
}

func GetUintSlice[Unsigned constraints.Unsigned](form url.Values, key string) ([]Unsigned, error) {
	if _, ok := form[key]; !ok {
		var v []Unsigned
		return v, nil
	}
	return ParseUintSlice[Unsigned](form[key], 10, 64)
}

func GetUint32Value(form url.Values, key string) (*wrapperspb.UInt32Value, error) {
	v, err := GetUint[uint32](form, key)
	return wrapperspb.UInt32(v), err
}

func GetUint32ValueSlice(form url.Values, key string) ([]*wrapperspb.UInt32Value, error) {
	v, err := GetUintSlice[uint32](form, key)
	return WrapUint32Slice(v), err
}

func GetUint64Value(form url.Values, key string) (*wrapperspb.UInt64Value, error) {
	v, err := GetUint[uint64](form, key)
	return wrapperspb.UInt64(v), err
}

func GetUint64ValueSlice(form url.Values, key string) ([]*wrapperspb.UInt64Value, error) {
	v, err := GetUintSlice[uint64](form, key)
	return WrapUint64Slice(v), err
}

func GetFloat[Float constraints.Float](form url.Values, key string) (Float, error) {
	if _, ok := form[key]; !ok {
		var v Float
		return v, nil
	}
	return ParseFloat[Float](form.Get(key), 64)
}

func GetFloatPtr[Float constraints.Float](form url.Values, key string) (*Float, error) {
	v, err := GetFloat[Float](form, key)
	return &v, err
}

func GetFloatSlice[Float constraints.Float](form url.Values, key string) ([]Float, error) {
	if _, ok := form[key]; !ok {
		var v []Float
		return v, nil
	}
	return ParseFloatSlice[Float](form[key], 64)
}

func GetFloat32Value(form url.Values, key string) (*wrapperspb.FloatValue, error) {
	v, err := GetFloat[float32](form, key)
	return wrapperspb.Float(v), err
}

func GetFloat32ValueSlice(form url.Values, key string) ([]*wrapperspb.FloatValue, error) {
	v, err := GetFloatSlice[float32](form, key)
	return WrapFloat32Slice(v), err
}

func GetFloat64Value(form url.Values, key string) (*wrapperspb.DoubleValue, error) {
	v, err := GetFloat[float64](form, key)
	return wrapperspb.Double(v), err
}

func GetFloat64ValueSlice(form url.Values, key string) ([]*wrapperspb.DoubleValue, error) {
	v, err := GetFloatSlice[float64](form, key)
	return WrapFloat64Slice(v), err
}

func FormFromMap(m map[string]string) url.Values {
	if m == nil {
		return nil
	}
	form := url.Values{}
	for key, value := range m {
		form.Add(key, value)
	}
	return form
}
