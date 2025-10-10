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
	rtdin,wstdin,_ := os.Pipe() // пайпуем Stdin вход
	oldStdin := os.Stdin // сохраняем оригинал
	os.Stdin = rtdin // перенапрвляем ввод на наш Pipe

	rstdount,wstdount,_ := os.Pipe() // пайпуем Stdout выход
	oldStdout := os.Stdout // сохраняем оригинал
	os.Stdout = wstdount // перенапрвляем вывод на наш Pipe

	_, err := io.WriteString(wstdin,"testdata/test.txt\n") // запись в Stdin
	if err != nil {
		panic(err)
	}
	err = wstdin.Close() // закрываем чтобы fmt не ждал бесконечно
	if err != nil {
		panic(err)
	}

	code.GetSize() // вызовы функции которая принимает захваченный Stdin
	err = wstdount.Close() // закрываем чтобы завершить чтение вывода
	if err != nil {
		panic(err)
	}
	os.Stdout = oldStdout // возвращаем в исходное состояние
	os.Stdin  = oldStdin

	var buf bytes.Buffer // буффер для чтения вывода
	_, err = io.Copy(&buf, rstdount) // копируем ответ GetSize в stdout
	if err != nil {
		panic(err)
	}

    output := buf.String() // байты -> строка
	result := fmt.Sprintf("%vB %s \n", 55, "testdata/test.txt") // ождиаемый результат
	require.Equal(t,output,result) // проверка через testify
}
