package types

import (
	"github.com/oklog/ulid"
)

type Keyword struct {
	Id       ulid.ULID
	Keyword  string
	UserId   ulid.ULID
	MaxPrice *float64
}
