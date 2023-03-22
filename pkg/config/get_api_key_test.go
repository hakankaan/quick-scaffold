package config

import (
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name     string
		setValue string
		want     string
		wantErr  bool
	}{
		{
			name:     "Get the stored API key",
			setValue: "example_api_key",
			want:     "example_api_key",
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := SetAPIKey(tc.setValue)
			if err != nil {
				t.Fatalf("SetAPIKey() error = %v", err)
			}

			got, err := GetAPIKey()
			if (err != nil) != tc.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tc.want)
			}
		})
	}
}
