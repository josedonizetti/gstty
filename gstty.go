package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("gstty", "prints window size")

	cols = app.Command("cols", "prints columns")
	rows = app.Command("rows", "prints rows")
)

func main() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	res := strings.Split(string(out), " ")

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case cols.FullCommand():
		cols := strings.TrimSpace(res[0])
		fmt.Println(cols)
	case rows.FullCommand():
		rows := strings.TrimSpace(res[1])
		fmt.Println(rows)
	}
}
