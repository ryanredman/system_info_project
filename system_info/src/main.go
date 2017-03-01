package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path[1:] {
	case "cpuinfo":
		fallthrough
	case "meminfo":
		fallthrough
	case "uptime":
		out, err := exec.Command("cat", "/proc/"+r.URL.Path[1:]).Output()
		if err != nil {
			log.Fatal(err)
		}

		out_reader := bytes.NewReader(out)

		scanner := bufio.NewScanner(out_reader)
		for scanner.Scan() {
			fmt.Fprintln(w, re_redun_whtsp.ReplaceAllString(scanner.Text(), " "))
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	default:
		//TODO - small html page with links to endpoints
		fmt.Fprintln(w, "Endpoints: cpuinfo, meminfo, uptime")
	}
}

var re_redun_whtsp *regexp.Regexp

func main() {

	re_redun_whtsp = regexp.MustCompile(`[\s\p{Zs}]{2,}`)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":1404", nil)
}
