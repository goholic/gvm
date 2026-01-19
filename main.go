package main

import (
	"fmt"
	"os/user"

	"github.com/goholic/gvm/lexer"
	"github.com/goholic/gvm/token"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the GVM programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	input := `let five = 5;
	let ten = 10;
	
	fn add(x, y) {
		x + y;
	}
	
	let result = add(five, ten);
	`

	fmt.Println("Tokenizing input:")
	fmt.Println(input)
	fmt.Println("-----------------")

	// l.ch = l
	// l.position = 0
	// l.readPosition = 1
	l := lexer.New(input)

	for {
		// emit tokens
		// on each iteration
		// bingo!
		tok := l.NextToken()
		fmt.Printf("%+v\n", tok)
		if tok.Type == token.EOF {
			break
		}
	}

}
