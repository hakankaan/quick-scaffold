package conversation

import (
	"reflect"
	"testing"
)

func TestParseResponse(t *testing.T) {
	testCases := []struct {
		name     string
		response string
		want     []FileOperation
		wantErr  bool
	}{
		{
			name: "Example response with create and modify file operations",
			response: `I understand your requirement here is what you should do.
 create_file: folder_name/file_name.extension
` + "```" + `
multiline
file content
another line
` + "```" + `
And
modify_file: folder_name/file_name.extension
` + "```" + `
new file content
` + "```" + `
delete_file: folder_name/file_name.extension`,
			want: []FileOperation{
				{
					Type:     "create_file",
					FileName: "folder_name/file_name.extension",
					Content: `multiline
file content
another line`,
				},
				{
					Type:     "modify_file",
					FileName: "folder_name/file_name.extension",
					Content:  "new file content",
				},
				{
					Type:     "delete_file",
					FileName: "folder_name/file_name.extension",
				},
			},
			wantErr: false,
		},
		{
			name: "Example response for asking file content",
			response: `
Need_content_of_file: folder_name/file_name.extension
Need_content_of_file: folder_name2/file_name2.extension
`,
			want: []FileOperation{
				{
					Type:     "Need_content_of_file",
					FileName: "folder_name/file_name.extension",
				},
				{
					Type:     "Need_content_of_file",
					FileName: "folder_name2/file_name2.extension",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseResponse(tc.response)
			if (err != nil) != tc.wantErr {
				t.Errorf("ParseResponse() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ParseResponse() = %v, want %v", got, tc.want)
			}
		})
	}
}
