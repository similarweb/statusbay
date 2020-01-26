package load_test

import (
	"errors"
	"os"
	"statusbay/notifiers"
	"statusbay/notifiers/common"
	"statusbay/notifiers/load"
	"statusbay/notifiers/testutil"
	"testing"
)

func TestNotifierRegistration(t *testing.T) {
	var registeredNotifierName common.NotifierName = "registered_notifier"
	var unRegisteredNotifierName common.NotifierName = "unregistered_notifier"

	notifierConfigs := common.ConfigByName{}

	t.Run("Trying to access an notifier that was not implemented", func(t *testing.T) {
		notifierConfigs[unRegisteredNotifierName] = nil
		defer delete(notifierConfigs, unRegisteredNotifierName)

		if _, err := load.Load(notifierConfigs, "", ""); err == nil {
			t.Error("Error expected")
		}
	})

	t.Run("Failing to read default config file", func(t *testing.T) {
		defer delete(notifierConfigs, registeredNotifierName)
		notifierConfigs[registeredNotifierName] = nil

		defer notifiers.Deregister(registeredNotifierName)
		notifiers.Register(registeredNotifierName, testutil.GetNotifierMakerMock("", ""))

		defer func(preTestFunc func(basePath string, notifierName common.NotifierName) (*os.File, error)) {
			load.GetDefaultConfigReaderFunc = preTestFunc
		}(load.GetDefaultConfigReaderFunc)

		load.GetDefaultConfigReaderFunc = func(basePath string, notifierName common.NotifierName) (file *os.File, err error) {
			err = errors.New("failed to read default config file")
			return
		}

		if _, err := load.Load(notifierConfigs, "", ""); err == nil {
			t.Error("Error expected")
		}
	})

	t.Run("Failing to initialize a notifier", func(t *testing.T) {
		errorMessage := "init fail"

		defer delete(notifierConfigs, registeredNotifierName)
		notifierConfigs[registeredNotifierName] = nil

		defer notifiers.Deregister(registeredNotifierName)
		notifiers.Register(registeredNotifierName, testutil.GetNotifierMakerMock("error", errorMessage))

		defer func(preTestFunc func(basePath string, notifierName common.NotifierName) (*os.File, error)) {
			load.GetDefaultConfigReaderFunc = preTestFunc
		}(load.GetDefaultConfigReaderFunc)

		load.GetDefaultConfigReaderFunc = func(basePath string, notifierName common.NotifierName) (file *os.File, err error) {
			file = &os.File{}
			return
		}

		if _, err := load.Load(notifierConfigs, "", ""); err == nil {
			t.Error("Error expected")
		} else if err.Error() != errorMessage {
			t.Errorf("Unexpected error message %s != %s", err.Error(), errorMessage)
		}
	})

	t.Run("Failing to load notifier config", func(t *testing.T) {
		errorMessage := "load config failed"

		defer delete(notifierConfigs, registeredNotifierName)
		notifierConfigs[registeredNotifierName] = nil

		defer notifiers.Deregister(registeredNotifierName)
		notifiers.Register(registeredNotifierName, testutil.GetNotifierMakerMock("mock", errorMessage))

		defer func(preTestFunc func(basePath string, notifierName common.NotifierName) (*os.File, error)) {
			load.GetDefaultConfigReaderFunc = preTestFunc
		}(load.GetDefaultConfigReaderFunc)

		load.GetDefaultConfigReaderFunc = func(basePath string, notifierName common.NotifierName) (file *os.File, err error) {
			file = &os.File{}
			return
		}

		if _, err := load.Load(notifierConfigs, "", ""); err == nil {
			t.Error("Error expected")
		} else if err.Error() != errorMessage {
			t.Errorf("Unexpected error message %s != %s", err.Error(), errorMessage)
		}
	})

	t.Run("Successfully initialized at least one notifier", func(t *testing.T) {
		expectedNumberOfNotifiers := 1

		defer delete(notifierConfigs, registeredNotifierName)
		notifierConfigs[registeredNotifierName] = nil

		defer notifiers.Deregister(registeredNotifierName)
		notifiers.Register(registeredNotifierName, testutil.GetNotifierMakerMock("mock", ""))

		defer func(preTestFunc func(basePath string, notifierName common.NotifierName) (*os.File, error)) {
			load.GetDefaultConfigReaderFunc = preTestFunc
		}(load.GetDefaultConfigReaderFunc)

		load.GetDefaultConfigReaderFunc = func(basePath string, notifierName common.NotifierName) (file *os.File, err error) {
			file = &os.File{}
			return
		}

		if notifierInstances, err := load.Load(notifierConfigs, "", ""); err != nil {
			t.Errorf("Unexpected error %s", err.Error())
		} else if len(notifierInstances) != expectedNumberOfNotifiers {
			t.Errorf("Unexpected number of notifier instances %d!=%d. %#v", expectedNumberOfNotifiers, len(notifierInstances), notifierInstances)
		}
	})
}
