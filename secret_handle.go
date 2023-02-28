package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func getSecretByName(name string) string {
	f, err := os.Open("_tokens.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")

		if splitLine[0] == name {
			return splitLine[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ""
}
