package main

import (
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rmpurp/knowhow/cli"
	"github.com/rmpurp/knowhow/dao"
	"log"
	"os"
)

func main() {
	db, err := dao.ConnectAndInitialize("./debug.db")

	if err != nil {
		log.Fatal(err)
	}

	err = dao.CreateFixtures(db)

	if err != nil {
		log.Fatal(err)
	}

	editCommand := flag.NewFlagSet("edit", flag.ExitOnError)
	openCommand := flag.NewFlagSet("open", flag.ExitOnError)
	searchCommand := flag.NewFlagSet("search", flag.ExitOnError)
	moveCommand := flag.NewFlagSet("move", flag.ExitOnError)

	switch os.Args[1] {
	case "o":
		fallthrough
	case "edit":
		editCommand.Parse(os.Args[2:])
	case "open":
		openCommand.Parse(os.Args[2:])
	case "s":
		fallthrough
	case "search":
		searchCommand.Parse(os.Args[2:])
	case "m":
		fallthrough
	case "move":
		moveCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if editCommand.Parsed() {
		fmt.Println("edit")
		edited, _ := cli.EditText("TEST TEST \nTEST TEST2")
		fmt.Println(edited)
	}

	if openCommand.Parsed() {
		fmt.Println("open")
	}

	if searchCommand.Parsed() {
		fmt.Println("search")
	}

	if moveCommand.Parsed() {
		fmt.Println("move")
	}

}
