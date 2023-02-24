package schemas

import (
	"io/ioutil"
	"testing"
)

func getDataTest(file string) []byte {
	body, err := ioutil.ReadFile("data_tests/" + file + ".yml")
	if err != nil {
		panic(err)
	}
	return body
}
func TestValidator(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want bool
	}{
		{
			name: "invalid data",
			data: getDataTest("invalid"),
			want: false,
		},
		{
			name: "invalid missing parameters data",
			data: getDataTest("invalid_missing_params"),
			want: false,
		},
		{
			name: "valid data",
			data: getDataTest("valid"),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := Validator(tt.data)

			if valid != tt.want {
				t.Errorf("Validator() valid = %t, want = %t, errors = %v", valid, tt.want, err)
			}

		})
	}
}
