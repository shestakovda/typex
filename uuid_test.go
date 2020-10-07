package typex_test

import (
	"github.com/shestakovda/errx"
	"github.com/shestakovda/typex"
)

func (s *TypesSuite) TestUUID() {
	s.Nil(typex.UUID(nil))
	s.Len(typex.NewUUID().Hex(), 32)
	s.Len(typex.NewUUID().String(), 36)

	if _, err := typex.ParseUUID(""); s.Error(err) {
		s.True(errx.Is(err, typex.ErrUUIDInvalid, typex.ErrUUIDEmpty))
	}

	if _, err := typex.ParseUUID("unknown"); s.Error(err) {
		s.True(errx.Is(err, typex.ErrUUIDInvalid))
		s.EqualError(errx.Unwrap(err), "invalid UUID length: 7")
	}

	if _, err := typex.ParseUUID("00000000000000000000000000000000"); s.Error(err) {
		s.True(errx.Is(err, typex.ErrUUIDInvalid, typex.ErrUUIDEmpty))
	}

	uid := typex.NewUUID()

	if uid2, err := typex.ParseUUID(uid.String()); s.NoError(err) {
		s.Equal(uid, uid2)
		s.Equal(uid.Hex(), uid2.Hex())
		s.Equal(uid.String(), uid2.String())
	}
}
