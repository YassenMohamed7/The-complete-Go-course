package main

import (
	"bufio"
	genericAdd "example/module5-pactice/generic_Add"
	"example/module5-pactice/note"
	"example/module5-pactice/summation"
	"example/module5-pactice/todo"
	"fmt"
	"os"
	"strings"
)

type Saver interface {
	Save() error
}
type displayer interface {
	Display()
}

func main() {
	fmt.Println(summation.Add(3, 5))
	fmt.Println(summation.Add(3.14, 5))
	fmt.Println(summation.Add("Hello ", "World"))
	// fmt.Println(summation.Add(3, "Hello")) casues panic because it is not supported type for b

	fmt.Println(genericAdd.GenericAdd(3, 5))
	fmt.Println(genericAdd.GenericAdd(3.14, 5.14))
	fmt.Println(genericAdd.GenericAdd("Hello ", "World"))

	myNote, err := note.New(getUserInput("Title: "), getUserInput("Content: "))
	if err != nil {
		panic(err)
	}

	display(myNote)
	err = saveToFile(myNote)
	if err != nil {
		panic(err)
	}
	var n *note.Note = new(note.Note)
	n.Content = "Hello World"
	n.Title = "My Note"
	n.Display()

	todo := todo.New("Todo")
	display(todo)
	err = saveToFile(todo)
	if err != nil {
		panic(err)
	}

	printSomething(1)
	printSomething("Hello")
	printSomething(3.14)

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	//  Remove the newline character from the end of the string
	// we can not do it using str = strings.TrimSuffix(str, "\n\r") because it will not work on all platforms, so we need to trim both \n and \r separately
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")
	if err != nil {
		panic(err)
	}
	return str
}

func saveToFile(s Saver) error {
	err := s.Save()
	if err != nil {
		return err
	}
	return nil
}

func display(s displayer) {
	s.Display()
}

func printSomething(value interface{}) {
	switch value.(type) {
	case string:
		fmt.Println("String value: ", value.(string))
	case int:
		fmt.Println("Integer value: ", value.(int))
	case float64:
		fmt.Println("Float value: ", value.(float64))
	default:
		fmt.Println("Unknown type")
	}
}
