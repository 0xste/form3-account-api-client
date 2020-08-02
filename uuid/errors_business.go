package uuid

import "fmt"

const errMsgInvalidUUID string = "invalid uuid provided %s"

// ErrInvalidUUID is returned if a UUID is not valid.
type ErrInvalidUUID struct {
	uuid string
}

func (e *ErrInvalidUUID) Error() string {
	return fmt.Sprintf(errMsgInvalidUUID, e.uuid)
}
