package uuid

const (
	errMsgFailureToGenerateUUID string = "failed to generate uuid"
	errMsgFailureToCompileRegex string = "failued to compile regex"
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
