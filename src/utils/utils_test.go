package utils_test

import (
	"testing"

	"github.com/rodrigocardosodev/pismo-challenge/src/utils"
	"github.com/stretchr/testify/require"
)

func TestUtils_IsValidCPF(t *testing.T) {
	err := utils.IsValidCPF("55724203014")
	require.Nil(t, err)

	err = utils.IsValidCPF("55724203015")
	require.NotNil(t, err)
}
