package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/kyokan/namegrind"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var inFile io.ReadCloser
	if len(os.Args) == 2 {
		f, err := os.Open(strings.TrimSpace(os.Args[1]))
		if err != nil {
			printErr(err)
		}
		inFile = f
	} else {
		inFile = ioutil.NopCloser(os.Stdin)
	}



	reservationsExist, err := namegrind.ReservationsExist()
	if err != nil {
		printErr(err)
	}
	if !reservationsExist {
		if err := namegrind.FetchReservations(); err != nil {
			printErr(err)
		}
	}

	reservations, err := namegrind.ParseReservations()
	if err != nil {
		printErr(err)
	}

	scan := bufio.NewScanner(inFile)
	fmt.Println("name,height,week,reserved")
	for scan.Scan() {
		name := scan.Text()
		nameHash, err := namegrind.HashName(name)
		if err != nil {
			printErr(err)
		}
		isReserved := reservations[hex.EncodeToString(nameHash)]
		height, week := namegrind.Rollout(nameHash)
		fmt.Printf("%s,%d,%d,%t\n", name, height, week, isReserved)
	}
}

func printErr(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
