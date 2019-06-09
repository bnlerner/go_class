package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map
and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal()
to create a JSON object from the map, and then your program should print the JSON object.
*/

///type struct Address {
//	name string
//	address string
//}

func main() {
	var idMap map[string]string
	idMap = make(map[string]string)

	fmt.Printf("enter name:")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	inputName := in.Text()

	fmt.Printf("enter address:")
	inAdr := bufio.NewScanner(os.Stdin)
	inAdr.Scan()
	inputAddress := inAdr.Text()

	idMap["name"] = inputName
	idMap["address"] = inputAddress

	barr, err := json.MarshalIndent(idMap, "", " ")

	if err == nil {
		strOut := string(barr)
		fmt.Println(strOut)
	} else {
		fmt.Println("error!")
	}
}
