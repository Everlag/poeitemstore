package main

import (
	"fmt"
	"os"

	"time"

	"github.com/Everlag/gothing/cmd"
	"github.com/Everlag/gothing/db"
)

func main() {

	db, err := db.Boot("")
	if err != nil {
		fmt.Printf("failed to open db, err=%s\n", err)
		os.Exit(-1)
	}
	defer db.Close()

	start := time.Now()
	cmd.HandleCommands(db)

	end := time.Now()
	fmt.Printf("command took %s\n", end.Sub(start))
}
