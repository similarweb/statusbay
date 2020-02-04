package notifiers

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"statusbay/notifiers/common"
)

type NotifierMaker func(defaultConfigReader io.Reader, urlBase string) (common.Notifier, error)

const notRegisteredTemplate = "notifier by the name %s was not registered"

var registeredNotifiers = map[common.NotifierName]NotifierMaker{}

// Register adds a notifier ctor to registeredNotifiers map
func Register(name common.NotifierName, newFunc NotifierMaker) {
	registeredNotifiers[name] = newFunc
}

// Deregister removes a notifier ctor from the registeredNotifiers map
func Deregister(name common.NotifierName) {
	delete(registeredNotifiers, name)
}

// GetNotifierMaker retrieves a notifier ctor from the registeredNotifiers map by name
func GetNotifierMaker(name common.NotifierName) (notifierMaker NotifierMaker, err error) {
	var implemented bool

	if notifierMaker, implemented = registeredNotifiers[name]; !implemented {
		err = errors.New(fmt.Sprintf(notRegisteredTemplate, name))
		return
	}

	return
}
