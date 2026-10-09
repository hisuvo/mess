package servers

import "errors"

func ServerTwo() error {
	return errors.New("server two connected failed")
}