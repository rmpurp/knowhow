package main

import (
	"fmt"
    "flag"
    "os"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
    openCommand := flag.NewFlagSet("open", flag.ExitOnError)
    searchCommand := flag.NewFlagSet("search", flag.ExitOnError)
    moveCommand := flag.NewFlagSet("move", flag.ExitOnError)

    switch os.Args[1] {
    case "o":
        fallthrough
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

