package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
		err   error
	}{
		"empty header":     {input: http.Header{}, want: "", err: ErrNoAuthHeaderIncluded},
		"no authorization": {input: http.Header{"Authorization": []string{""}}, want: "", err: ErrNoAuthHeaderIncluded},
		"no api key":       {input: http.Header{"Authorization": []string{"Bearer"}}, want: "", err: ErrMalformedAuthHeader},
		"api key":          {input: http.Header{"Authorization": []string{"ApiKey 123"}}, want: "123", err: nil},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tt.input)
			if !errors.Is(err, tt.err) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.err)

			}

			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}

		})
	}
}
