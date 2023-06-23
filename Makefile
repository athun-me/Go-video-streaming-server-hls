outfile:
	ffmpeg -i football.mp4 -c:v libx264 -preset veryfast -g 48 -keyint_min 48 -sc_threshold 0 -b:v 2500k -maxrate 2500k -bufsize 5000k -c:a aac -b:a 128k -hls_time 10 -hls_playlist_type vod -hls_segment_filename "output%03d.ts" playlist.m3u8

run:
	go run main.go