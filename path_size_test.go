package code_test

import(
	"code"
	"testing"
	"os"
	"bytes"
	"fmt"
	"io"
	"github.com/stretchr/testify/require"
)

func TestGetSize(t *testing.T) {
	rtdin,wstdin,_ := os.Pipe()
	oldStdin := os.Stdin 
	os.Stdin = rtdin 

	rstdount,wstdount,_ := os.Pipe()
	oldStdout := os.Stdout
	os.Stdout = wstdount

	_, err := io.WriteString(wstdin,"testdata/test.txt\n")
	if err != nil {
		panic(err)
	}
	err = wstdin.Close()
	if err != nil {
		panic(err)
	}

	code.GetSize()
	err = wstdount.Close()
	if err != nil {
		panic(err)
	}
	os.Stdout = oldStdout
	os.Stdin  = oldStdin

	var buf bytes.Buffer
	_, err = io.Copy(&buf, rstdount)
	if err != nil {
		panic(err)
	}

    output := buf.String()
	result := fmt.Sprintf("%vB %s \n", 55, "testdata/test.txt")
	require.Equal(t,output,result)
}
