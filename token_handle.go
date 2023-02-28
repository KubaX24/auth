package main

import (
	"fmt"
	"sync"
	"time"
)

var tokenToToken = sync.Map{}

func ManageTimeTokens(timeToken string, userToken string) {
	tokenToToken.Store(timeToken, userToken)
	fmt.Println("[Time] token created! " + timeToken + " and map to " + userToken)
	fmt.Println("http://localhost:8080/time/" + timeToken) //TODO: Delete
	timeStamp := time.Now().Unix()

	for range time.Tick(time.Second * 1) {
		if (timeStamp + 15) < time.Now().Unix() {
			removeToken(timeToken)
			return
		}
	}
}

func removeToken(timeToken string) {
	if _, exist := tokenToToken.Load(timeToken); exist {
		tokenToToken.Delete(timeToken)

		fmt.Println("[Time] token removed! ")
	}
}

func getTokenFromTime(timeToken string) (string, int) {
	token, exist := tokenToToken.Load(timeToken)

	if exist {
		return fmt.Sprint(token), 200
	}

	return "Time not found", 404
}
