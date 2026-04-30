package retries

import "fmt"

type Action func() error

func WithRetries(action Action, attempts int) Action {
	var err error
	return func() error {
		for i := 0; i < attempts; i++ {
			err = action()
			if err == nil {
				return nil
			}
		}
		return fmt.Errorf("Action failed after %d attempts with error: %v", attempts, err)
	}
}
