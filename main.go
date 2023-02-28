package main

import (
	"auth/utils"
	"fmt"
)

var version = "b1"
var port = "8080"

func main() {
	start()
}

func start() {
	fmt.Println(utils.White + "\n             _   _       _____ _    ___     _________       _____ " + utils.Reset)
	fmt.Println(utils.White + "            | | | |     / ____| |  | \\ \\   / /__   __|/\\   / ____|" + utils.Reset)
	fmt.Println(utils.White + "  __ _ _   _| |_| |__  | |    | |__| |\\ \\_/ /   | |  /  \\ | |     " + utils.Reset)
	fmt.Println(utils.White + " / _` | | | | __| '_ \\ | |    |  __  | \\   /    | | / /\\ \\| |     " + utils.Reset)
	fmt.Println(utils.White + "| (_| | |_| | |_| | | || |____| |  | |  | |     | |/ ____ \\ |____ " + utils.Reset)
	fmt.Println(utils.White + " \\__,_|\\__,_|\\__|_| |_(_)_____|_|  |_|  |_|     |_/_/    \\_\\_____|" + utils.Reset)
	fmt.Println(utils.White + "  (Version " + version + ")\n" + utils.Reset)
	fmt.Println(utils.Red + "Starting server..." + utils.Reset)

	Server()
}
