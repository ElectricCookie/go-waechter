package waechter

import "sync"
import "fmt"

var blocked sync.Map

//EnableThrottle determines whether there is a check for blocking. This is only used for testing. Disabling throttles in production is NOT recommended!
var EnableThrottle = true

func try(identifier string) error {
	_, ok := blocked.Load(identifier)
	if !ok && EnableThrottle {
		return fmt.Errorf("Ressource blocked")
	}

	blocked.Store(identifier, true)
	return nil

}

func release(identifier string) {

	blocked.Delete(identifier)

}
