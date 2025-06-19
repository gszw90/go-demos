package helper

import (
	"fmt"
	"runtime"
	"testing"
)

type TestStruct struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func TestMapToStruct(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]interface{}
		want    TestStruct
		wantErr bool
		errMsg  string // 预期的错误信息
	}{
		{
			name: "Valid map to struct conversion",
			input: map[string]interface{}{
				"name":  "John Doe",
				"age":   30,
				"email": "john.doe@example.com",
			},
			want: TestStruct{
				Name:  "John Doe",
				Age:   30,
				Email: "john.doe@example.com",
			},
			wantErr: false,
			errMsg:  "",
		},
		{
			name: "Invalid map to struct conversion",
			input: map[string]interface{}{
				"name":  "Jane Doe",
				"age":   "thirty", // Invalid type for age
				"email": "jane.doe@example.com",
			},
			want:    TestStruct{},
			wantErr: true,
			errMsg:  "failed to unmarshal into struct",
		},
		{
			name:    "Nil input map",
			input:   nil,
			want:    TestStruct{},
			wantErr: true,
			errMsg:  "input map cannot be nil",
		},
		{
			name: "Nil output struct",
			input: map[string]interface{}{
				"name":  "John Doe",
				"age":   30,
				"email": "john.doe@example.com",
			},
			want:    TestStruct{},
			wantErr: true,
			errMsg:  "output struct cannot be nil",
		},
	}

	fmt.Println("os:", runtime.GOOS)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result TestStruct
			var err error

			if tt.name == "Nil output struct" {
				// 测试 s 为 nil 的情况
				err = MapToStruct(tt.input, nil)
			} else {
				err = MapToStruct(tt.input, &result)
			}

			// 检查是否返回了预期的错误
			if (err != nil) != tt.wantErr {
				t.Errorf("MapToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 检查错误信息是否匹配
			if tt.wantErr && err != nil && err.Error() != tt.errMsg {
				t.Errorf("MapToStruct() error message = %v, want %v", err.Error(), tt.errMsg)
			}

			// 检查转换结果是否正确
			if !tt.wantErr && result != tt.want {
				t.Errorf("MapToStruct() = %v, want %v", result, tt.want)
			}
		})
	}
}
