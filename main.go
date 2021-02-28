package main

import (
	"fmt";
	"os/exec";
	//"os";
)

func catchErrors(err interface{}) {
	if err != nil {
		return
	}
}

func main() {
	fmt.Println("Welcome to Sheepella")
	fmt.Println("I'll help to enrich your English :)")
	fmt.Println("")

	scrap()

	cmd := exec.Command("cmd", "/c", "start", "https://translate.google.com.br/?hl=pt-BR&sl=en&tl=pt&op=translate")
	_, _ = cmd.Output()

	fmt.Println("Click on the link you want to access")
}