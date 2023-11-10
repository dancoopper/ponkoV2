package main

import (
	"log"
	"os"
)

func main() {
	//fmt.Println("getting pic")
	//GetPic()
	//GetName()
	//fmt.Print(Geturl())

	err := os.Mkdir("pic", 777)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	DownLoadPic()
	MainWin()

}
