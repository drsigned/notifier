package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/drsigned/gos"
	"github.com/drsigned/notifier/pkg/notifier"
	"github.com/logrusorgru/aurora/v3"
)

type options struct {
	oneline bool
}

var (
	co options
	so notifier.Options
)

func banner() {
	fmt.Fprintln(os.Stderr, aurora.BrightBlue(`
             _   _  __ _           
 _ __   ___ | |_(_)/ _(_) ___ _ __ 
| '_ \ / _ \| __| | |_| |/ _ \ '__|
| | | | (_) | |_| |  _| |  __/ |   
|_| |_|\___/ \__|_|_| |_|\___|_| v1.0.0
`).Bold())
}

func init() {
	flag.BoolVar(&co.oneline, "l", false, "")

	flag.Usage = func() {
		banner()

		h := "USAGE:\n"
		h += "  notifier [OPTIONS]\n"

		h += "\nOPTIONS:\n"
		h += "  -l        send message line by line (default: false)\n"

		fmt.Fprintf(os.Stderr, h)
	}

	flag.Parse()
}

func main() {
	if !gos.HasStdin() {
		os.Exit(1)
	}

	options, err := notifier.ParseOptions(&so)
	if err != nil {
		log.Fatalln(err)
	}

	notifier, err := notifier.New(options)
	if err != nil {
		log.Fatalln(err)
	}

	var lines string
	var message string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		message = line

		if co.oneline {
			notifier.SendNotification(message)
		} else {
			lines += line
			lines += "\n"
		}
	}

	if !co.oneline {
		message = lines

		notifier.SendNotification(message)
	}
}
