# extractaudiofromvideo
- download FFmpeg

## import go packages
	go get github.com/gofiber/fiber/v2
	go get github.com/gofiber/fiber/v2/middleware/logger

## RUN
    go run .
### use Get curl 
    curl --location 'localhost:3000/extract-audio?input=sample-5s.mp4'