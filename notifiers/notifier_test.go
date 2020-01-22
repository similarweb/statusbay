package notifiers_test

import (
	"io"
	"statusbay/notifiers"
	"statusbay/notifiers/common"
	"testing"
)

func TestNotifierRegistration(t *testing.T) {
	var notifierName common.NotifierName = "test_notifier"
	var notifierMaker notifiers.NotifierMaker = func(defaultConfigReader io.Reader, config common.NotifierConfig, urlBase string) (notifier common.Notifier, err error) {
		return
	}

	notifiers.Register(notifierName, notifierMaker)

	t.Run("Checking if the notifier was registered", func(t *testing.T) {
		if testMakerFunc, err := notifiers.GetNotifierMaker(notifierName); err != nil {
			t.Errorf("Unexpected error %v", err.Error())
		} else if testMakerFunc == nil {
			t.Error("Unexpected nil maker func")
		}
	})

	notifiers.Deregister(notifierName)

	t.Run("Checking if the notifier was de-registered", func(t *testing.T) {
		if _, err := notifiers.GetNotifierMaker(notifierName); err == nil {
			t.Error("Error expected")
		}
	})

	t.Run("Checking if an unregistered notifier is returned", func(t *testing.T) {
		if _, err := notifiers.GetNotifierMaker("unregistered_notifier"); err == nil {
			t.Error("Error expected")
		}
	})
}

func TestUnregisteredNotifierDoesNotReturn(t *testing.T) {
	var notifierName common.NotifierName = "unregistered_notifier"

	t.Run("Checking if the an unregistered notifier is returned", func(t *testing.T) {
		if _, err := notifiers.GetNotifierMaker(notifierName); err == nil {
			t.Error("notifier returned, expected error")
		}
	})
}

func TestMultipleNotifierRegistration(t *testing.T) {
	notifiersToRegister := map[common.NotifierName]notifiers.NotifierMaker{
		"test_notifier1": func(_ io.Reader, _ common.NotifierConfig, _ string) (_ common.Notifier, _ error) { return },
		"test_notifier2": func(_ io.Reader, _ common.NotifierConfig, _ string) (_ common.Notifier, _ error) { return },
		"test_notifier3": func(_ io.Reader, _ common.NotifierConfig, _ string) (_ common.Notifier, _ error) { return },
	}

	for notifierName, notifierMaker := range notifiersToRegister {
		notifiers.Register(notifierName, notifierMaker)
	}

	t.Run("Checking if all notifiers were registered", func(t *testing.T) {
		for notifierName := range notifiersToRegister {
			if testMakerFunc, err := notifiers.GetNotifierMaker(notifierName); err != nil {
				t.Errorf("Unexpected error %v", err.Error())
			} else if testMakerFunc == nil {
				t.Error("Unexpected nil maker func")
			}
		}
	})

	t.Run("Checking if an unregistered notifier is returned", func(t *testing.T) {
		if _, err := notifiers.GetNotifierMaker("test_notifier4"); err == nil {
			t.Error("Error expected")
		}
	})

	for notifierName := range notifiersToRegister {
		notifiers.Deregister(notifierName)
	}

	t.Run("Checking if the notifier was de-registered", func(t *testing.T) {
		for notifierName := range notifiersToRegister {
			if _, err := notifiers.GetNotifierMaker(notifierName); err == nil {
				t.Error("Error expected")
			}
		}
	})
}
