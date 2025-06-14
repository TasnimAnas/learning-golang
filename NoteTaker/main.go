package main

import (
	"NoteTaker/note"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	title, err := takeUserInput("Enter note title")
	if err != nil {
		fmt.Println(err)
	}
	content, err := takeUserInput("Enter note content")
	if err != nil {
		fmt.Println(err)
	}
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
	}
	userNote.Display()
	err = userNote.SaveJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Note saved!")

}

func takeUserInput(prompt string) (string, error) {
	fmt.Printf(prompt + ": ")
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	data = strings.TrimSuffix(data, "\n")
	data = strings.TrimSuffix(data, "\r")

	return data, nil
}
