package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/hakankaan/quick-scaffold/pkg/config"
	"github.com/spf13/cobra"
)

type captureWriter struct {
	buf *bytes.Buffer
}

func (cw *captureWriter) Write(p []byte) (n int, err error) {
	return cw.buf.Write(p)
}

func executeCommand(cmd *cobra.Command, args []string, buf *bytes.Buffer) (string, error) {
	cw := &captureWriter{buf: buf}
	cmd.SetOut(cw)
	cmd.SetErr(cw)
	cmd.SetArgs(args)

	if err := cmd.Execute(); err != nil {
		return "", err
	}

	return strings.TrimSpace(buf.String()), nil
}

func TestSetAPIKeyCmd(t *testing.T) {

	testRootCmd := &cobra.Command{}
	testRootCmd.AddCommand(setAPIKeyCmd)

	testCases := []struct {
		name        string
		apiKey      string
		expectedMsg string
	}{
		{
			name:        "Set a valid API key",
			apiKey:      "example_api_key",
			expectedMsg: "API key set successfully",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			// Replace setAPIKeyCmd with testRootCmd in the line below
			output, err := executeCommand(testRootCmd, []string{"setapikey", tc.apiKey}, buf)
			if err != nil {
				t.Errorf("executeCommand() error = %v", err)
				return
			}

			if strings.TrimSpace(output) != tc.expectedMsg {
				t.Errorf("executeCommand() output = %v, expectedMsg %v", output, tc.expectedMsg)
			}

			err = config.SetAPIKey(tc.apiKey)
			if err != nil {
				t.Errorf("config.SetAPIKey() error = %v", err)
				return
			}

			storedAPIKey, err := config.GetAPIKey()
			if err != nil {
				t.Errorf("config.GetAPIKey() error = %v", err)
				return
			}

			if storedAPIKey != tc.apiKey {
				t.Errorf("Expected API key: %v, got: %v", tc.apiKey, storedAPIKey)
			}
		})
	}
}
