package main

import (
	"fmt"
	"math/rand"
	"os"
)

func Shred(path string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("problem reading the file: %v", err)
		return
	}
	
	stat, err := file.Stat()
	if err != nil {
		fmt.Printf("problem getting file stats: %v", err)
		return
	}

	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	for s := 0; s < 3; s++ {
		err = file.Truncate(0) //sets size to zero
		if err != nil {
			fmt.Printf("problem truncating file: %v", err)
		}

		for i := 0; i < int(stat.Size()); i++ {

			index := rand.Intn(len(charSet))
			c := charSet[index]
			file.WriteString(string(c))
		}
	}
	file.Close()
	err = os.Remove(path)
	if err != nil {
		fmt.Printf("problem deleting file: %v", err)
		return
	}
}
