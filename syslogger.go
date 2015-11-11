package main

import (
	"bufio"
	"flag"
	"fmt"
	"log/syslog"
	"os"
)

func main() {
	tag := flag.String("tag", "go", "syslog tag or application name")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	logwriter, e := syslog.New(syslog.LOG_NOTICE|syslog.LOG_LOCAL4, *tag)
	defer logwriter.Close()
	if e != nil {
		os.Exit(1)
	}

	for scanner.Scan() {
		logwriter.Write([]byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
