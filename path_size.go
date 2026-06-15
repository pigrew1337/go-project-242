package code

import (
<<<<<<< HEAD
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
=======
	"fmt"
	"io/fs"
	"path/filepath"
)

type FileInf struct {
	Name  string
	Size  int64
	Count int64
}

func Calculate(path string, format string) (string, error) {
	file := &FileInf{}

	err := GetSize(path, file)
	if err != nil {
		return "", err
	}

	switch format {

	case "b":
		return fmt.Sprintf("%dB %s (%d files)", file.Size, file.Name, file.Count), nil

	case "k":
		return fmt.Sprintf("%.2fKB %s (%d files)", float64(file.Size)/1024, file.Name, file.Count), nil

	case "m":
		return fmt.Sprintf("%.2fMB %s (%d files)", float64(file.Size)/1024/1024, file.Name, file.Count), nil

	case "g":
		return fmt.Sprintf("%.2fGB %s (%d files)", float64(file.Size)/1024/1024/1024, file.Name, file.Count), nil

	default:
		return fmt.Sprintf("%s %s (%d files)", HumanSize(file.Size), file.Name, file.Count), nil
	}
}

func HumanSize(size int64) string {
	const unit = 1024.0

	if size < int64(unit) {
		return fmt.Sprintf("%dB", size)
	}

	units := []string{"KB", "MB", "GB", "TB", "PB"}

	value := float64(size)
	i := 0

	for value >= unit && i < len(units)-1 {
		value /= unit
		i++
	}

	value = float64(int(value*100)) / 100

	return fmt.Sprintf("%.2f%s", value, units[i])
}

func GetSize(filename string, f *FileInf) error {
	path, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	f.Name = path

	var size int64
	var count int64

	err = filepath.WalkDir(path, func(_ string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		count++
		size += info.Size()

		return nil
	})

	if err != nil {
		return err
	}

	f.Size = size
	f.Count = count

	return nil
}
>>>>>>> ed21168 (init)
