package code_test

import (
	"code"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize(t *testing.T) {
	stTest := struct {
		name string
		size int64
	}{
		name: "testdata/test.txt",
		size: 55,
	}
	stCode := &code.FileInf{}
	err := code.GetSize("testdata/test.txt", stCode)

	require.NoError(t, err)
	require.Equal(t, stTest.name, stCode.Name)
	require.Equal(t, stTest.size, stCode.Size)
}

func TestCalculateBytes(t *testing.T) {
	res, err := code.Calculate("testdata/test.txt", "b")
	require.NoError(t, err)
	require.Contains(t, res, "55B")
}

func TestCalculateAuto(t *testing.T) {
	res, err := code.Calculate("testdata", "auto")
	require.NoError(t, err)
	require.NotEmpty(t, res)
}
