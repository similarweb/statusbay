package config_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"statusbay/config"
	"testing"
)

func TestWebserver(t *testing.T) {

	_, filename, _, _ := runtime.Caller(0)
	currentFolderPath := filepath.Dir(filename)

	t.Run("valid", func(t *testing.T) {
		config, err := config.LoadConfigWebserver(fmt.Sprintf("%s/testutil/mock/test-config.yaml", currentFolderPath))

		if err != nil {
			t.Fatalf("unexpected not error")
		}

		if reflect.TypeOf(config).String() != "config.Webserver" {
			t.Fatalf("unexpected configuration data")
		}

	})
	t.Run("invalid", func(t *testing.T) {

		_, err := config.LoadConfigWebserver(fmt.Sprintf("%s/testutil/mock/no-config.yaml", currentFolderPath))

		if err == nil {
			t.Fatalf("unexpected load configuration error")
		}
	})
	t.Run("invalid_schema", func(t *testing.T) {
		data, err := config.LoadConfigWebserver(fmt.Sprintf("%s/testutil/mock/no-config.yaml", currentFolderPath))

		if err == nil {
			t.Fatalf("unexpected load configuration error")
		}
		t.Log(data)
	})

}
