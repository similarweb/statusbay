package config_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"statusbay/config"
	"testing"
)

func TestKubernetes(t *testing.T) {

	_, filename, _, _ := runtime.Caller(0)
	currentFolderPath := filepath.Dir(filename)

	t.Run("valid", func(t *testing.T) {
		cfg, err := config.LoadKubernetesConfig(fmt.Sprintf("%s/testutil/mock/test-config.yaml", currentFolderPath))

		if err != nil {
			t.Fatalf("unexpected error ")
		}

		if reflect.TypeOf(cfg).String() != "config.Kubernetes" {
			t.Fatalf("unexpected configuration data")
		}

	})
	t.Run("invalid", func(t *testing.T) {

		_, err := config.LoadKubernetesConfig(fmt.Sprintf("%s/testutil/mock/no-config.yaml", currentFolderPath))

		if err == nil {
			t.Fatalf("unexpected load configuration error")
		}
	})
	t.Run("invalid_schema", func(t *testing.T) {
		data, err := config.LoadKubernetesConfig(fmt.Sprintf("%s/testutil/mock/no-config.yaml", currentFolderPath))

		if err == nil {
			t.Fatalf("unexpected load configuration error")
		}
		t.Log(data)
	})

}
