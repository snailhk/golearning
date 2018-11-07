package main

import (
	"learning/errorhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		defer func() {
			r := recover()
			log.Printf("Panic: %v", r)
			http.Error(writer, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}()
		err := handler(writer, request)
		if err != nil {
			log.Println("error", err)
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/list/",
		errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
