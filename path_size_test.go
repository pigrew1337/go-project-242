package code_test

import(
	"code"
	"testing"
	"os"
	"bytes"
	//"fmt"
	"io"
	//"github.com/stretchr/testify/assert"
	//"code/cmd/hexlet-path-size"
)

func TestGetSize(t *testing.T) {
	rtdin,wstdin,_ := os.Pipe()
	oldStdin := os.Stdin 
	os.Stdin = rtdin 

	rstdount,wstdount,_ := os.Pipe()
	oldStdout := os.Stdout
	os.Stdout = wstdount

	io.WriteString("testadata/test.txt\n")
	wstdin.Close()

	code.
}