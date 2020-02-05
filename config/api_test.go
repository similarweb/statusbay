package config_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"statusbay/config"
	"testing"
)

func TestAPI(t *testing.T) {

	_, filename, _, _ := runtime.Caller(0)
	currentFolderPath := filepath.Dir(filename)

	t.Run("valid", func(t *testing.T) {
		config, err := config.LoadConfigAPI(fmt.Sprintf("%s/testutil/mock/test-config.yaml", currentFolderPath))

		if err != nil {
			t.Fatalf("unexpected error %s", err.Error())
		}
		fmt.Println(reflect.TypeOf(config).String())
		if reflect.TypeOf(config).String() != "config.API" {
			t.Fatalf("unexpected configuration data")
		}

	})
	t.Run("invalid", func(t *testing.T) {

		_, err := config.LoadConfigAPI(fmt.Sprintf("%s/testutil/mock/no-config.yaml", currentFolderPath))

		if err == nil {
			t.Fatalf("unexpected load configuration error")
		}
	})
	t.Run("invalid_schema", func(t *testing.T) {
		data, err := config.LoadConfigAPI(fmt.Sprintf("%s/testutil/mock/no-config.yaml", currentFolderPath))

		if err == nil {
			t.Fatalf("unexpected load configuration error")
		}
		t.Log(data)
	})

}
