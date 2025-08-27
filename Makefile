# запуск линтера по ссылке 
lint:
	~/go/bin/golangci-lint run
# билдим в бинарник 
build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
# запуск базы
go run:
	go run  cmd/hexlet-path-size/main.go
# запуск бинарника
binarnik:
	./bin/hexlet-path-size