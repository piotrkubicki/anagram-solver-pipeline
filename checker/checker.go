package checker

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

func check(phrase string, hashes []string) bool {
	decodedHash := md5.Sum([]byte(phrase))
	hashValue := hex.EncodeToString(decodedHash[:])
	for _, hash := range hashes {
		if hashValue == hash {
			log.Printf("Found password: %v", phrase)
			return true
		}
	}
	return false
}

func Run(hashes []string, passwords *[]string, inputChannel chan string) {
	counter := 0
	for {
		if counter == len(hashes) {
			break
		}
		phrase := <-inputChannel
		if check(phrase, hashes) {
			counter++
			*passwords = append(*passwords, phrase)
		}
	}
}
