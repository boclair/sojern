package main

import (
	"os"
	"time"
)

// TODO: "github.com/fsnotify/fsnotify" seems like it would nicely
// create events on file system changes.  However, I could not get it
// to work quickly.  In the meantime, here is a mechanism to poll for
// the filename.

// WARNING: If ever there are multiple instances of the webserver,
// there may arise a race condition where some servers will return 200
// and some will still return 503, and visa-versa.  This method cannot
// be used for server synchronization.

type FileExists int

const (
	FILE_EXISTS FileExists = iota
	FILE_ABSENT
)

const poleInterval = 2 * time.Second

// Creates a filewatch to poll the given filename to test if the file exists.
// Writes the result to the returned channel.
func NewFileWatcher(filename string) <-chan FileExists {
	ch := make(chan FileExists)

	go func() {
		for {
			_, err := os.Stat(filename)
			if os.IsNotExist(err) {
				ch <- FILE_ABSENT
			} else {
				ch <- FILE_EXISTS
			}
		}
	}()

	return ch
}
