package client

import "fmt"

const errMsgBadAccountRequest string = "field %s is invalid with value %v"
type ErrBadAccountRequest struct {
	FieldKey string
	FieldValue interface{}
}
func (e ErrBadAccountRequest) Error() string {
	return fmt.Sprintf(errMsgBadAccountRequest, e.FieldKey, e.FieldValue)
}