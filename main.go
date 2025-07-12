package main

import(
	"os"
	"os/user"
	"fmt"
	"monkey/repl"
)

func main() {

	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf( "Hello %s, Welcome!\n", u.Username)
	fmt.Printf( "This is Monkey REPL: \n")

	repl.Start(os.Stdin, os.Stdout)

}
