package code_test

<<<<<<< HEAD
import(
	"code"
	"testing"
	"github.com/stretchr/testify/require"
	"math"
)

type fileInfTest struct{
	name string
    size  int64
	formatsize float64
}

func TestGetSize(t *testing.T) {
	stTest := fileInfTest{
		name: "testdata/test.txt", 
		size: 55,
	}
	stCode := &code.FileInf{}
	code.GetSize("testdata/test.txt",stCode)
	require.Equal(t,stTest.name,stCode.Name)
	require.Equal(t,stTest.size,stCode.Size)
}

func TestFormatSize(t *testing.T) {
	stTest := fileInfTest{
		name: "testdata", 
		formatsize: 0.1,
	}
	stCode := &code.FileInf{}
	code.GetSize("testdata",stCode)
	code.FormatSize(stCode)
	require.Equal(t,stTest.name,stCode.Name)
	stCode.FormatSize = math.Round(stCode.FormatSize*10) / 10
	require.Equal(t,stCode.FormatSize,stTest.formatsize)
=======
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
>>>>>>> ed21168 (init)
}
