package main

import (
    "fmt"
    "log"
    "time"

    "github.com/fsnotify/fsnotify"
)

func main() {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    done := make(chan bool)

    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                fmt.Println("Event detected:", event)
                if event.Op&fsnotify.Write == fsnotify.Write {
                    fmt.Println("Modified file:", event.Name)
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                fmt.Println("Error:", err)
            }
        }
    }()

    err = watcher.Add("/path/to/your/scripts")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Watching folder...")
    <-done
}
