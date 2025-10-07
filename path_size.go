package code

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

// 2 шаг
func Cli() {
	cmd := &cli.Command{
        Name:  "hexlet-path-size",
        Usage: "print size of a file or directory;",
    }
    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}

// 3 шаг
func GetSize() {
	var filename string //ввод
	fmt.Scan(&filename)
	fileInfo, err := os.Lstat(filename) // файл или папка
	if err != nil {
		log.Fatal("распознование файл/папка: ", err)
	}
	var size int64
	if fileInfo.IsDir() { // если папка
		readFile, err := os.ReadDir(filename) // читаем папку
		if err != nil {
			log.Fatal("чтение папки: ", err)
		}
		for _, entry := range readFile { // проходимся по папке
			info, err := entry.Info() // получаем инфо о файлах
			if err != nil {
				log.Fatal("инфо о файле: ", err)
			}
			size += info.Size()
		}
	} else {
		size = fileInfo.Size()
	}
	res := fmt.Sprintf("%vB %s \n", size, filename)
	io.WriteString(os.Stdout, res) // можно просто через println
}
