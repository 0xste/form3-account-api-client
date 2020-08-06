package uuid

import "fmt"

const (
	errMsgFailureToGenerateUUID string = "failed to generate uuid"
	errMsgFailureToCompileRegex string = "failed to compile regex"
	errMsgInvalidUUID           string = "invalid uuid provided %s"
)

// ErrFailureToGenerateUUID is returned if a UUID is not able to be generated.
type ErrFailureToGenerateUUID struct{}

func (e *ErrFailureToGenerateUUID) Error() string {
	return errMsgFailureToGenerateUUID
}

// ErrFailureToGenerateUUID is returned if a UUID is not able to be generated.
type ErrFailureToCompileRegex struct{}

func (e *ErrFailureToCompileRegex) Error() string {
	return errMsgFailureToCompileRegex
}

// ErrInvalidUUID is returned if a UUID is not valid.
type ErrInvalidUUID struct {
	uuid string
}

func (e *ErrInvalidUUID) Error() string {
	return fmt.Sprintf(errMsgInvalidUUID, e.uuid)
}
