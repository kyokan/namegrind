package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/kyokan/namegrind"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Usage: namegrind <infile>")
		os.Exit(1)
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

	inFile := strings.TrimSpace(os.Args[1])
	f, err := os.Open(inFile)
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		name := scan.Text()
		nameHash, err := namegrind.HashName(name)
		if err != nil {
			printErr(err)
		}
		isReserved := reservations[hex.EncodeToString(nameHash)]
		week, height := namegrind.Rollout(nameHash)
		fmt.Printf("%s,%d,%d,%t\n", name, week, height, isReserved)
	}
}

func printErr(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
