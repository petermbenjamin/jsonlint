package jsonlint

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestLint(t *testing.T) {
	var d map[string]interface{}
	sb := bytes.NewBufferString("{\"one\":\"hello world!\", \"two\":2}")
	i := sb.Bytes()
	json.Unmarshal(i, &d)
	o, _ := json.MarshalIndent(d, "", "    ")

	type args struct {
		jsonData []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Lint",
			args{i},
			o,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Lint(tt.args.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lint() = %v, want %v", got, tt.want)
			}
		})
	}
}
