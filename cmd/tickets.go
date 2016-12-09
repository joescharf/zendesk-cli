// Copyright © 2016 Automox, Inc. <support@automox.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/joescharf/zego/zego"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var outFile string

// exportticketsCmd represents the tickets command
var exportticketsCmd = &cobra.Command{
	Use:   "tickets",
	Short: "Export all zendesk tickets to csv",
	Long:  `Output will be placed in zendesk.csv`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("-> Using Zendesk: username: %s, Subdomain: %s\n", viper.GetString("username"), viper.GetString("subdomain"))
		auth := zego.Auth{viper.GetString("username") + "/token", viper.GetString("token"), viper.GetString("subdomain")}

		tickets, err := auth.ListTickets()
		if err != nil {
			fmt.Println("** Error listing Zendesk tickets: ", err)
			return
		}

		// out, err := gocsv.MarshalString(tickets.Tickets)
		// if err != nil {
		// 	fmt.Printf("Error converting to CSV", err)
		// 	return
		// }
		fmt.Println("-> Creating Output File ", outFile)
		file, err := os.Create(outFile)
		if err != nil {
			fmt.Println("** Error creating output file: ", err)
			return
		}
		defer file.Close()

		fmt.Println("-> Writing CSV Data")
		err = gocsv.MarshalFile(tickets.Tickets, file)
		if err != nil {
			fmt.Println("** Error writing to outputfile: ", err)
			return
		}

		fmt.Println("Zendesk Ticket export complete.")

	},
}

func init() {
	exportCmd.AddCommand(exportticketsCmd)
	exportCmd.Flags().StringVarP(&outFile, "out", "o", "zendesk.csv", "Set csv output filename")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ticketsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ticketsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
