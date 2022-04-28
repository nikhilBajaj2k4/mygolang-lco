package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// As it works on my machine. However, for the next block you need a pointer
	// to the variables you're assigning the input to. Try replacing fmt.Scanln(text2)
	// with fmt.Scanln(&text2). Don't use Sscanln, because it parses a string already
	// in memory instead of from stdin. If you want to do something like what you were
	// trying to do, replace it with fmt.Scanf("%s", &ln)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a string")

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println(input)
}
