package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMerchant(t *testing.T) {
	merchant, err := NewMerchant()

	require.NoError(t, err)
	require.NotNil(t, merchant)
}
