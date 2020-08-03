package uuid

import (
	"crypto/rand"
	"fmt"
	"regexp"
)

// validates both upper and lower case uuids
const patternUUID4 string = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}"

type UUID string

func (u *UUID) String() string {
	return string(*u)
}

// FromStringV4
func FromStringV4(uuid string) (UUID, error) {
	if IsUUIDv4(uuid) {
		return UUID(uuid), nil
	}
	return "", &ErrInvalidUUID{uuid}
}

// IsUUIDv4 returns true if a uuid is valid
// this could theoretically panic if the UUID is invalid tests cover this,
//the tradeoff on a "tidier" signature imo is worth the minimal risk of an unhandled panic
func IsUUIDv4(uuid string) bool {
	return regexp.MustCompile(patternUUID4).MatchString(uuid)
}

// NewV4 generates a random UUID
func NewV4() (UUID, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", &ErrFailureToGenerateUUID{}
	}
	return FromStringV4(fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:]))
}
