# Go-video-streaming-server-hls

<h2>Description </h2>
HLS is a streaming protocol that allows large media files to be served as many smaller text files that are broken up into roughly  10-second increments. By breaking them up, the user’s client-side application only needs to buffer ~10 seconds in advance. This saves the user a lot of potential bandwidth and allows songs or videos to start playback almost immediately.

## Repositories :-

- Refferd from <br> - https://blog.boot.dev/golang/golang-video-stream-server/ 
- Test HLS streams in all supported browsers <br> - https://hlsjs-dev.video-dev.org/demo/


## Generate outputlist.m3u8

```bash
$ make outfile
```

## Running the app

```bash
# development
$ make run
```

## Author

- [Athun Lal](https://www.linkedin.com/in/athun-lal-0103631ba/)

