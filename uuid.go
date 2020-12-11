package typex

import (
	"bytes"
	"encoding/hex"
	"strings"

	"github.com/google/uuid"
	"github.com/shestakovda/errx"
)

const hyphen = "-"

// ZeroUUID - пустой UUID, 16 нулей
var ZeroUUID = make(UUID, 16)

var (
	ErrUUIDInvalid = errx.New("invalid uuid")
	ErrUUIDEmpty   = errx.New("empty uuid")
)

var (
	MsgUUIDInvalid = "Некорректный UUIDv4: `%s`"
	MsgUUIDEmpty   = "Пустой UUIDv4"
)

// UUID - unique identifier according RFC4122
type UUID []byte

func NewUUID() UUID {
	uid := uuid.New()
	return UUID(uid[:])
}

func ParseUUID(id string) (_ UUID, err error) {
	var uid uuid.UUID

	if id == "" {
		return nil, ErrUUIDInvalid.WithReason(ErrUUIDEmpty.WithDetail(MsgUUIDEmpty))
	}

	if uid, err = uuid.Parse(id); err != nil {
		return nil, ErrUUIDInvalid.WithReason(err).WithDetail(MsgUUIDInvalid, id)
	}

	if uid == uuid.Nil {
		return nil, ErrUUIDInvalid.WithReason(ErrUUIDEmpty.WithDetail(MsgUUIDEmpty))
	}

	return uid[:], nil
}

func (u UUID) String() string {
	if u.IsEmpty() {
		return ""
	}
	return strings.Join([]string{
		hex.EncodeToString(u[:4]),
		hex.EncodeToString(u[4:6]),
		hex.EncodeToString(u[6:8]),
		hex.EncodeToString(u[8:10]),
		hex.EncodeToString(u[10:16]),
	}, hyphen)
}

func (u UUID) Hex() string {
	if u.IsEmpty() {
		return ""
	}
	return hex.EncodeToString(u)
}

func (u UUID) IsEmpty() bool {
	if len(u) < 16 {
		return true
	}

	if bytes.Equal(u, ZeroUUID) {
		return true
	}

	return false
}
