package snowdayphonechain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZero(t *testing.T) {
	require := require.New(t)

	minutes := uint32(0)
	callsPerPerson := uint32(2)
	result := CalculateNumPhoneCalls(minutes, callsPerPerson)
	require.Equal(uint64(0), result)
}
func TestKnownValue2(t *testing.T) {
	require := require.New(t)

	minutes := uint32(2)
	callsPerPerson := uint32(2)
	result := CalculateNumPhoneCalls(minutes, callsPerPerson)
	require.Equal(uint64(6), result)

	minutes = 3
	callsPerPerson = 2
	result = CalculateNumPhoneCalls(minutes, callsPerPerson)
	require.Equal(uint64(14), result)
}

func TestKnownValue3(t *testing.T) {
	require := require.New(t)

	minutes := uint32(2)
	callsPerPerson := uint32(3)
	result := CalculateNumPhoneCalls(minutes, callsPerPerson)
	require.Equal(uint64(12), result)

	minutes = 3
	callsPerPerson = 3
	result = CalculateNumPhoneCalls(minutes, callsPerPerson)
	require.Equal(uint64(39), result)
}
