package main

import (
    "io"
    "net/http"
    "os"
)

func videoHandler(w http.ResponseWriter, r *http.Request) {
    videoFile, err := os.Open("videos/video.mp4")
    if err != nil {
        http.Error(w, "Video not found.", http.StatusNotFound)
        return
    }
    defer videoFile.Close()

    // Set appropriate cache-control headers to prevent browser caching
    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")

    // Set CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    w.Header().Set("Content-Type", "video/mp4")

    // Read the entire video file into memory
    videoData, err := io.ReadAll(videoFile)
    if err != nil {
        http.Error(w, "Error reading video file.", http.StatusInternalServerError)
        return
    }

    // Write the video data directly to the response writer
    _, err = w.Write(videoData)
    if err != nil {
        http.Error(w, "Error writing video data.", http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/video", videoHandler)
    http.ListenAndServe(":8080", nil)
}
