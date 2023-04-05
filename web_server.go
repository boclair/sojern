package main

import (
	"log"
	"net/http"
)

const httpPort = ":8080"
const filename = "/tmp/ok"

func main() {
	fileChannel := NewFileWatcher(filename)
	StartServer(fileChannel, httpPort)

	emptyChan := make(chan struct{})
	<-emptyChan
}

func StartServer(fileChannel <-chan FileExists, httpPort string) http.Server {
	mux := http.NewServeMux()
	logWrapper := NewLoggingWrapper()

	okHandler := NewOkHandler()
	go connectWatcher(fileChannel, okHandler)

	imgHandler := logWrapper.wrap(NewImgHandler())

	mux.Handle("/ping", okHandler)
	mux.Handle("/img", imgHandler)

	log.Println("Starting http server on port", httpPort)

	srv := http.Server{
		Addr:    httpPort,
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	return srv
}

// Repeatedly reads from the provided channel to check if the file exists.
// If it does, then 
func connectWatcher(ch <-chan FileExists, okHandler OkHandler) {
	for {
		result, ok := <-ch
		if !ok {
			return
		}
		
		switch result {
		case FILE_EXISTS:
			okHandler.setStatus()
		
		case FILE_ABSENT:
			okHandler.clearStatus()

		default:
			log.Fatal("Unexpected file status", result)
		}
	}
}
