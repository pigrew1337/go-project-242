package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
)

// ответ мерчанту
type Response struct {
	Size  int64  `json:"size"`
	Count int    `json:"count"`
	Type  string `json:"type"`
}

const (
	Port = "8080"
)

func startServer() {
	mux := http.NewServeMux()
	// подвязываем html
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r,
			filepath.Join("cmd", "hexlet-path-size", "web", "index.html"),
		)
	})
	// подвязываем статику
	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(
				filepath.Join("cmd", "hexlet-path-size", "web", "static"),
			)),
		),
	)
	// обрабатываем файлы
	mux.HandleFunc("/upload", uploadHandler)
	// коносоль вывод о порте
	log.Printf("Server started: http://localhost:%s", Port)
	log.Fatal(http.ListenAndServe(":"+Port, mux))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// парсер
	err := r.ParseMultipartForm(200 << 20)
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	// список файлов из формы
	files := r.MultipartForm.File["files"]

	var total int64
	count := len(files)

	//проход по файлам
	for _, fh := range files {
		f, _ := fh.Open()

		buf := make([]byte, 1024)
		// читаем файлы
		for {
			n, err := f.Read(buf)
			total += int64(n)
			if err != nil {
				break
			}
		}

		f.Close()
	}
	// ответ json
	json.NewEncoder(w).Encode(Response{
		Size:  total,
		Count: count,
	})
}
