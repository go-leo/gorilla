package gorilla

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func ParseBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func ParseInt[Signed constraints.Signed](s string, base int, bitSize int) (Signed, error) {
	i, err := strconv.ParseInt(s, base, bitSize)
	return Signed(i), err
}

func ParseUint[Unsigned constraints.Unsigned](s string, base int, bitSize int) (Unsigned, error) {
	i, err := strconv.ParseUint(s, base, bitSize)
	return Unsigned(i), err
}

func ParseFloat[Float constraints.Float](s string, bitSize int) (Float, error) {
	f, err := strconv.ParseFloat(s, bitSize)
	return Float(f), err
}

func ParseBoolSlice(s []string) ([]bool, error) {
	if s == nil {
		return nil, nil
	}
	r := make([]bool, 0, len(s))
	for _, str := range s {
		b, err := strconv.ParseBool(str)
		if err != nil {
			return nil, err
		}
		r = append(r, b)
	}
	return r, nil
}

func ParseIntSlice[Signed constraints.Signed](s []string, base int, bitSize int) ([]Signed, error) {
	if s == nil {
		return nil, nil
	}
	r := make([]Signed, 0, len(s))
	for _, str := range s {
		i, err := ParseInt[Signed](str, base, bitSize)
		if err != nil {
			return nil, err
		}
		r = append(r, i)
	}
	return r, nil
}

func ParseUintSlice[Unsigned constraints.Unsigned](s []string, base int, bitSize int) ([]Unsigned, error) {
	if s == nil {
		return nil, nil
	}
	r := make([]Unsigned, 0, len(s))
	for _, str := range s {
		i, err := ParseUint[Unsigned](str, base, bitSize)
		if err != nil {
			return nil, err
		}
		r = append(r, i)
	}
	return r, nil
}

func ParseFloatSlice[Float constraints.Float](s []string, bitSize int) ([]Float, error) {
	if s == nil {
		return nil, nil
	}
	r := make([]Float, 0, len(s))
	for _, str := range s {
		f, err := ParseFloat[Float](str, bitSize)
		if err != nil {
			return nil, err
		}
		r = append(r, f)
	}
	return r, nil
}

func ParseBytesSlice(s []string) [][]byte {
	if s == nil {
		return nil
	}
	r := make([][]byte, 0, len(s))
	for _, str := range s {
		r = append(r, []byte(str))
	}
	return r
}
