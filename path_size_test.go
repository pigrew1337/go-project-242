package code_test

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
}
