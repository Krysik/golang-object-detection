package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"github.com/otiai10/gosseract"
	"io"
    "log"
    "os"
    "net/http"
)


const IMAGE_NAME = "image.png"

func main() {
	deviceID := 1

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Capture Window")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	fmt.Printf("Start reading device: %v\n", deviceID)
	record := true
	
	
	for record {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}
		window.IMShow(img)
		if window.WaitKey(1) == 32 {
			gocv.IMWrite(IMAGE_NAME, img)
			text := detectText(&img)
			err := WriteToFile(text)
		    if err != nil {
		        log.Fatal(err)
		    }
			record = false
		}
	}
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fs)
	http.ListenAndServe(":8081", nil)
}

func detectText(img *gocv.Mat) string {
	fmt.Println("TEXT DETECTING")
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(IMAGE_NAME)
	text, _ := client.Text()
	fmt.Println("DETECTED TEXT\n" + text)
	return text
}

func WriteToFile(data string) error {
    file, err := os.Create("text.txt")
    if err != nil {
        return err
    }
    defer file.Close()

    textToSave := "Detected Text\n" + data

    _, err = io.WriteString(file, textToSave)
    if err != nil {
        return err
    }
    return file.Sync()
}


