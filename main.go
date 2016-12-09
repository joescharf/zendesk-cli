package main

import (
	"fmt"
	"github.com/joescharf/zendesk-cli/cmd"
	// "io/ioutil"
	// "os"
	// "strconv"
)

// type ZendeskTicket struct {
// 	Id int `csv:"id"`
// 	CreatedAt
// }

func main() {

	fmt.Println("** Zendesk CLI**\n")
	cmd.Execute()
	fmt.Println()

}
