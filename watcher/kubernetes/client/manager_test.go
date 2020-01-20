package client

import (
	"os"
	"testing"
)

func TestNewClientManager(t *testing.T) {

	t.Run("valid client", func(t *testing.T) {
		os.Setenv("KUBERNETES_SERVICE_HOST", "127.0.0.1")
		os.Setenv("KUBERNETES_SERVICE_PORT", "443")
		kubernetesClientManager, err := NewClientManager("", "")

		if err != nil {
			t.Fatalf("unexpected error message, error should be empty")
		}

		kubernetesClient := kubernetesClientManager.GetInsecureClient()
		if kubernetesClient == nil {
			t.Fatalf("unexpected kubernetes client value. should be not empty")
		}

		os.Unsetenv("KUBERNETES_SERVICE_HOST")
		os.Unsetenv("KUBERNETES_SERVICE_PORT")
	})

	t.Run("Invalid client configuration", func(t *testing.T) {

		_, err := NewClientManager("", "")

		if err == nil {
			t.Fatalf("Error should be not empty")
		}

	})

}
