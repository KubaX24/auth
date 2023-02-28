package main

import (
	"auth/utils"
	"bufio"
	"log"
	"os"
	"strings"
)

func getWebsiteByCode(code string) (string, string) {
	f, err := os.Open("_websites.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")

		if splitLine[1] == code {
			return splitLine[0], splitLine[2]
		}
	}

	log.Println(utils.Red + code + " [Website] was not found!")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return "", ""
}
