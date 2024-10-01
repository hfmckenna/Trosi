package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// MockFileSystem implements FileSystem for testing
type MockFileSystem map[string]string

func (m MockFileSystem) Open(name string) (io.ReadCloser, error) {
	content, ok := m[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return io.NopCloser(strings.NewReader(content)), nil
}

func TestReadFile(t *testing.T) {
	mockFS := MockFileSystem{
		"testfile.txt": "Hello, World!",
	}

	tests := []struct {
		name     string
		fileName string
		want     string
		wantErr  bool
	}{
		{"Existing file", "testfile.txt", "Hello, World!", false},
		{"Non-existent file", "nonexistent.txt", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantErr {
						t.Errorf("readFile() panicked unexpectedly: %v", r)
					}
				}
			}()

			got := readFile(mockFS, tt.fileName)
			if got != tt.want {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
