# GoLang Camera Streaming

### How to run
```
# serve the default webcam /dev/video0 
go run main.go
``` 
```
# serve an ip camera stream
$ go run main.go rstp://camera_ip:554/
```
```
# serve another a usb camera
$ go run main.go /dev/video1
```
