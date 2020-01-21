package client

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

const (
	tokenFile  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	rootCAFile = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
)

func GenSecrets() {
	// write token
	token := []byte("token")
	err2 := ioutil.WriteFile(tokenFile, token, 0644)
	fmt.Println("@@@@@@@@@@@@")
	if err2 != nil {
		fmt.Println(err2)
	}
	cmd := exec.Command("ls ~")
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))
}
func TestNewClientManager(t *testing.T) {
	GenSecrets()
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
