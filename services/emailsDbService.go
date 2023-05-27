package services

import (
	"os"
	"log"
	"bufio"
)

func ReadEmailsList(file *os.File) []string {
	var emails []string 
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		emails = append(emails, scanner.Text())
	}


	return emails
}

func FindEmailInFile(email string, file *os.File) bool {	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if email == scanner.Text() {
			return true
		}
	}
	
	return false
}

func WriteEmailToFile(email string, file *os.File) {
	_, err := file.WriteString(email + "\n")
	if err != nil {
		log.Println(err.Error())
	}
}
