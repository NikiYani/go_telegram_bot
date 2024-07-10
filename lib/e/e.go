package e

import "fmt"

func Wrap(errMsg string, err error) error {
	return fmt.Errorf("%s: %w", errMsg, err)
}

func WrapIfErr(errMsg string, err error) error {
	if err == nil {
		return nil
	}

	return Wrap(errMsg, err)
}
