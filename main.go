package main

import (
	"log"
	"net/http"
)

func main() {
	const port = 8080
	const hlsDir = "hls"

	http.Handle("/", addHeaders(http.FileServer(http.Dir(hlsDir))))
	log.Printf("Serving HLS files from directory '%s' on port %d\n", hlsDir, port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

// ffmpeg -i football.mp4 -c:v libx264 -preset veryfast -g 48 -keyint_min 48 -sc_threshold 0 -b:v 2500k -maxrate 2500k -bufsize 5000k -c:a aac -b:a 128k -hls_time 10 -hls_playlist_type vod -hls_segment_filename "output%03d.ts" playlist.m3u8
