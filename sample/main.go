package main

import (
	"fmt"
	"os"

	"github.com/rehacktive/motion-detect"
)

func main() {

	d := motion.New(motion.DefaultThresold, motion.DefaultSensitivity, 10000, "output.jpg")
	detected, err := d.DetectMotion("/ramfs/1.jpg", "/ramfs/2.jpg")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println(detected)
}
