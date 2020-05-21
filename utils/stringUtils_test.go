package utils

import (
	"testing"
)

func TestPrintStruct(t *testing.T) {
	type TaskConfig struct {
		DbURL      string
		DbPassword string
	}
	type Config struct {
		Name       string
		Enabled    bool
		TaskConfig TaskConfig
	}

	testCases := []Config{
		{"Aditya", true, TaskConfig{}},
	}
	for _, test := range testCases {
		s := PrintStruct(test)
		t.Log(s)
	}
}
