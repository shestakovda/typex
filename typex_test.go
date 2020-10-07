package typex_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestTypes(t *testing.T) {
	suite.Run(t, new(TypesSuite))
}

type TypesSuite struct {
	suite.Suite
}
