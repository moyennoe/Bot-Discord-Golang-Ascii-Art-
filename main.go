package main

import (
	"fmt"

	"./bot"
	"./config"
)

func main() {
	fmt.Println(" ____         _____        __     __  _   _    ____   __      __        _____               _____    _____    _____  ")
	fmt.Println("|  _ \\   _   / ____|       \\ \\   / / | \\ | |  / __ \\  \\ \\    / /       |  __ \\      /\\     |  __ \\  |_   _|  / ____| ")
	fmt.Println("| |_) | / | | |             \\ \\_/ /  |  \\| | | |  | |  \\ \\  / /        | |__) |    /  \\    | |__) |   | |   | (___   ")
	fmt.Println("|  _ <  | | | |              \\   /   | . ` | | |  | |   \\ \\/ /         |  ___/    / /\\ \\   |  _  /    | |    \\___ \\  ")
	fmt.Println("| |_) | | | | |____           | |    | |\\  | | |__| |    \\  /          | |       / ____ \\  | | \\ \\   _| |_   ____) | ")
	fmt.Println("|____/  |_|  \\_____|          |_|    |_| \\_|  \\____/      \\/           |_|      /_/    \\_\\ |_|  \\_\\ |_____| |_____/  ")

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
