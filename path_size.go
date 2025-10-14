package code

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	//"bytes"
	"github.com/urfave/cli/v3"
)

type FileInf struct {
    Name string
    Size  int64
	FormatSize float64
}

// 2 шаг
func Cli() {
	/*cmd := &cli.Command{
        Name:  "hexlet-path-size",
        Usage: "print size of a file or directory;",
    }*/
	cmd := &cli.Command{
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:  "human",
                Usage: "human-readable sizes (auto-select unit)",
            },
        },
        Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.NArg() > 0 && cmd.String("human") == "false"{
                input := cmd.Args().First()
				fileHandler := &FileInf{}
				GetSize(input,fileHandler)
				res := fmt.Sprintf("%vB %s \n", fileHandler.Size, fileHandler.Name)
				_, err := io.WriteString(os.Stdout, res)
				if err != nil {
					panic(err)
				}
            } else {
				input := cmd.Args().First()
				fileHandler := &FileInf{}
				GetSize(input,fileHandler)
        		FormatSize(fileHandler)
				res := fmt.Sprintf("%.1fM %s \n", fileHandler.FormatSize, fileHandler.Name)
				_, err := io.WriteString(os.Stdout, res)
				if err != nil {
				panic(err)
				}
            }
            return nil
        },
    }
    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}

func FormatSize(f *FileInf) {
	f.FormatSize = float64(f.Size) * 1e-06
}

// 3 шаг
func GetSize(filename string,f *FileInf) {
	f.Name = filename
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
			f.Size = size
		}
	} else {
		size = fileInfo.Size()
		f.Size = size
	}
}
