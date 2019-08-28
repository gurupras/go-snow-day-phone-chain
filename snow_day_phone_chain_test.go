package snowdayphonechain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKnownValue2(t *testing.T) {
	require := require.New(t)

	minutes := uint32(2)
	callsPerPerson := uint32(2)
	result := CalculateNumPhoneCalls(minutes, callsPerPerson)
	require.Equal(uint64(6), result)
}

func TestKnownValue3(t *testing.T) {
	require := require.New(t)

	minutes := uint32(3)
	callsPerPerson := uint32(3)
	result := CalculateNumPhoneCalls(minutes, callsPerPerson)
	require.Equal(uint64(39), result)
}
