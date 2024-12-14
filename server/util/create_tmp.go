package util

import (
	"log"
	"os"
)

func CreateTmp() {
	dirName := "tmp"
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
