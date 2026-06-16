package code

import (
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
