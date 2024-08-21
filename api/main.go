package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "math/rand"
    "time"
)

func videoHandler(w http.ResponseWriter, r *http.Request) {
    videoFiles, err := os.ReadDir("videos/")
    if err != nil {
        http.Error(w, "Error reading video directory.", http.StatusInternalServerError)
        return
    }

    numFiles := len(videoFiles)

    randomInt := rand.Intn(numFiles)

    selectedFile := videoFiles[randomInt].Name()
    videoFilePath := fmt.Sprintf("videos/%s", selectedFile)

    videoFile, err := os.Open(videoFilePath)
    if err != nil {
        http.Error(w, "Video not found.", http.StatusNotFound)
        return
    }
    defer videoFile.Close()

    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")

    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Content-Type", "video/mp4")

    _, err = io.Copy(w, videoFile)
    if err != nil {
        http.Error(w, "Error writing video data.", http.StatusInternalServerError)
        return
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())

    http.HandleFunc("/video", videoHandler)
    http.ListenAndServe("0.0.0.0:8080", nil)
}
