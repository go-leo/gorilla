package path

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// ---- Mock Services ----

type MockBoolPathService struct{}

func (m *MockBoolPathService) BoolPath(ctx context.Context, req *BoolPathRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockInt32PathService struct{}

func (m *MockInt32PathService) Int32Path(ctx context.Context, req *Int32PathRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockInt64PathService struct{}

func (m *MockInt64PathService) Int64Path(ctx context.Context, req *Int64PathRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockUint32PathService struct{}

func (m *MockUint32PathService) Uint32Path(ctx context.Context, req *Uint32PathRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockUint64PathService struct{}

func (m *MockUint64PathService) Uint64Path(ctx context.Context, req *Uint64PathRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockFloatPathService struct{}

func (m *MockFloatPathService) FloatPath(ctx context.Context, req *FloatPathRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockDoublePathService struct{}

func (m *MockDoublePathService) DoublePath(ctx context.Context, req *DoublePathRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockStringPathService struct{}

func (m *MockStringPathService) StringPath(ctx context.Context, req *StringPathRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockEnumPathService struct{}

func (m *MockEnumPathService) EnumPath(ctx context.Context, req *EnumPathRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

// ---- Test Cases ----

func TestBoolPath(t *testing.T) {
	router := mux.NewRouter()
	router = AppendBoolPathGorillaRoute(router, &MockBoolPathService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/true/false/true"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"bool":true,"optBool":false,"wrapBool":true}`
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestInt32Path(t *testing.T) {
	router := mux.NewRouter()
	router = AppendInt32PathGorillaRoute(router, &MockInt32PathService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/1/2/3/4/5/6/7"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"int32":1,"sint32":2,"sfixed32":3,"optInt32":4,"optSint32":5,"optSfixed32":6,"wrapInt32":7}`
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestInt64Path(t *testing.T) {
	router := mux.NewRouter()
	router = AppendInt64PathGorillaRoute(router, &MockInt64PathService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/10/20/30/40/50/60/70"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"int64":"10","sint64":"20","sfixed64":"30","optInt64":"40","optSint64":"50","optSfixed64":"60","wrapInt64":"70"}`
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestUint32Path(t *testing.T) {
	router := mux.NewRouter()
	router = AppendUint32PathGorillaRoute(router, &MockUint32PathService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/1/2/3/4/5"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"uint32":1,"fixed32":2,"optUint32":3,"optFixed32":4,"wrapUint32":5}`
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestUint64Path(t *testing.T) {
	router := mux.NewRouter()
	router = AppendUint64PathGorillaRoute(router, &MockUint64PathService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/10/20/30/40/50"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"uint64":"10","fixed64":"20","optUint64":"30","optFixed64":"40","wrapUint64":"50"}`
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestFloatPath(t *testing.T) {
	router := mux.NewRouter()
	router = AppendFloatPathGorillaRoute(router, &MockFloatPathService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/1.23/4.56/7.89"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"float":1.23,"optFloat":4.56,"wrapFloat":7.89}`
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestDoublePath(t *testing.T) {
	router := mux.NewRouter()
	router = AppendDoublePathGorillaRoute(router, &MockDoublePathService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/1.23/4.56/7.89"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"double":1.23,"optDouble":4.56,"wrapDouble":7.89}`
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestStringPath(t *testing.T) {
	router := mux.NewRouter()
	router = AppendStringPathGorillaRoute(router, &MockStringPathService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/abc/def/ghi"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"string":"abc","optString":"def","wrapString":"ghi"}`
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestEnumPath(t *testing.T) {
	router := mux.NewRouter()
	router = AppendEnumPathGorillaRoute(router, &MockEnumPathService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/1/2"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"status":"OK","optStatus":"CANCELLED"}`
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}
