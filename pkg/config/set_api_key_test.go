package config

import (
	"testing"
)

func TestSetAPIKey(t *testing.T) {
	testCases := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "Set a valid API key",
			value: "example_api_key",
			want:  "example_api_key",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := SetAPIKey(tc.value)
			if err != nil {
				t.Errorf("SetAPIKey() error = %v", err)
			}

			got, err := GetAPIKey()
			if err != nil {
				t.Errorf("GetAPIKey() error = %v", err)
			}

			if got != tc.want {
				t.Errorf("Expected API key: %v, got: %v", tc.want, got)
			}
		})
	}
}
