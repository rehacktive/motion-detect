package main

import (
	"fmt"
	"os"

	"github.com/rehacktive/motion-detect"
)

func main() {

	d := motion.New(motion.DefaultThresold, motion.DefaultSensitivity, motion.DefaultMinArea, "output.jpg")
	detected, err := d.DetectMotion("/home/aw4y/motion_images/img/diff1/1.jpg", "/home/aw4y/motion_images/img/diff1/2.jpg")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println(detected)
}
