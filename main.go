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

	// record := make([]string, 9)
	// heading := []string{"id", "created_at", "type", "status", "subject", "requester_id", "submitter_id", "recipient", "description"}

	// file, err := os.Create("zendesk.csv")
	// if err != nil {
	// 	glog.Fatalln("Error creating output file: ", err)
	// }
	// defer file.Close()

	// w := csv.NewWriter(file)
	// if err = w.Write(heading); err != nil {
	// 	glog.Fatalln("Error writing header to csv: ", err)
	// }

	// tickets, _ := auth.ListTickets()
	// for _, t := range tickets.Tickets {
	// 	record[0] = strconv.Itoa(int(t.Id))
	// 	record[1] = t.CreatedAt
	// 	record[2] = t.Type
	// 	record[3] = t.Status
	// 	record[4] = t.Subject
	// 	record[5] = strconv.Itoa(int(t.RequesterId))
	// 	record[6] = strconv.Itoa(int(t.SubmitterId))
	// 	record[7] = t.Recipient
	// 	record[8] = t.Description
	// 	if err := w.Write(record); err != nil {
	// 		glog.Fatalln("Error writing record to csv: ", err)
	// 	}
	// 	// fmt.Print(t.Subject + " ")
	// }
	// w.Flush()
}
