package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpRule_RegularizePath(t *testing.T) {
	tests := []struct {
		name               string
		path               string
		expectedPath       string
		expectedName       string
		expectedTemplate   string
		expectedParameters []string
	}{
		{
			name:               "Successfully regularize path with named parameters",
			path:               "{name=things/*/otherthings/*}",
			expectedPath:       "things/{thingsId}/otherthings/{otherthingsId}",
			expectedName:       "name",
			expectedTemplate:   "things/%s/otherthings/%s",
			expectedParameters: []string{"thingsId", "otherthingsId"},
		},
		{
			name:               "No named parameters in path",
			path:               "things/123/otherthings/456",
			expectedPath:       "things/123/otherthings/456",
			expectedName:       "",
			expectedTemplate:   "",
			expectedParameters: nil,
		},
		{
			name:               "Empty path",
			path:               "",
			expectedPath:       "",
			expectedName:       "",
			expectedTemplate:   "",
			expectedParameters: nil,
		},
		{
			name:               "Path with named parameter but no wildcards",
			path:               "{name=things}",
			expectedPath:       "things",
			expectedName:       "name",
			expectedTemplate:   "things",
			expectedParameters: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &HttpRule{}
			path, name, template, params := r.RegularizePath(tt.path)

			assert.Equal(t, tt.expectedPath, path)
			assert.Equal(t, tt.expectedName, name)
			assert.Equal(t, tt.expectedTemplate, template)
			assert.Equal(t, tt.expectedParameters, params)
		})
	}
}

func TestRegularizePath(t *testing.T) {
	// Setup a dummy HttpRule (the method does not use the struct fields)
	r := &HttpRule{}

	type want struct {
		path                string
		namedPathName       string
		template            string
		namedPathParameters []string
	}

	tests := []struct {
		name  string
		input string
		want  want
	}{
		{
			name:  "no named path parameter",
			input: "/v1/things/{id}",
			want: want{
				path:                "/v1/things/{id}",
				namedPathName:       "",
				template:            "",
				namedPathParameters: nil,
			},
		},
		{
			name:  "single named path parameter",
			input: "/v1/{name=things/*}",
			want: want{
				path:                "/v1/things/{thing}",
				namedPathName:       "name",
				template:            "things/%s",
				namedPathParameters: []string{"{thing}"},
			},
		},
		{
			name:  "multiple named path parameters",
			input: "/v1/{name=things/*/otherthings/*}",
			want: want{
				path:                "/v1/things/{thing}/otherthing/{otherthing}",
				namedPathName:       "name",
				template:            "things/%s/otherthings/%s",
				namedPathParameters: []string{"thingId", "otherthingId"},
			},
		},
		{
			name:  "named path parameter with prefix and suffix",
			input: "/v1/projects/{resource=things/*/otherthings/*}/details",
			want: want{
				path:                "/v1/projects/things/{thingId}/otherthings/{otherthingId}/details",
				namedPathName:       "resource",
				template:            "things/%s/otherthings/%s",
				namedPathParameters: []string{"thingId", "otherthingId"},
			},
		},
		{
			name:  "empty path",
			input: "",
			want: want{
				path:                "",
				namedPathName:       "",
				template:            "",
				namedPathParameters: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPath, gotName, gotTemplate, gotParams := r.RegularizePath(tt.input)
			if gotPath != tt.want.path {
				t.Errorf("path = %q, want %q", gotPath, tt.want.path)
			}
			if gotName != tt.want.namedPathName {
				t.Errorf("namedPathName = %q, want %q", gotName, tt.want.namedPathName)
			}
			if gotTemplate != tt.want.template {
				t.Errorf("template = %q, want %q", gotTemplate, tt.want.template)
			}
			if len(gotParams) != len(tt.want.namedPathParameters) {
				t.Errorf("namedPathParameters = %v, want %v", gotParams, tt.want.namedPathParameters)
			} else {
				for i := range gotParams {
					if gotParams[i] != tt.want.namedPathParameters[i] {
						t.Errorf("namedPathParameters[%d] = %q, want %q", i, gotParams[i], tt.want.namedPathParameters[i])
					}
				}
			}
		})
	}
}
