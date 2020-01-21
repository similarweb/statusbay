package notifiers

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"statusbay/notifiers/common"
)

type NotifierMaker func(defaultConfigReader io.Reader, config common.NotifierConfig, urlBase string) (common.Notifier, error)

const notRegisteredTemplate = "notifier by the name %s was not registered"

var registeredNotifiers = map[common.NotifierName]NotifierMaker{}

func Register(name common.NotifierName, newFunc NotifierMaker) {
	registeredNotifiers[name] = newFunc
}

func GetNotifierMaker(name common.NotifierName) (notifierMaker NotifierMaker, err error) {
	var implemented bool

	if notifierMaker, implemented = registeredNotifiers[name]; !implemented {
		err = errors.New(fmt.Sprintf(notRegisteredTemplate, name))
		return
	}

	return
}
