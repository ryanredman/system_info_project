package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	//"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//TODO - small html page with links to endpoints
	fmt.Fprintln(w, "Get: cpuinfo, meminfo, uptime")
}

func mem_handler(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("cat", "/proc/meminfo").Output()
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
}

func cpu_handler(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("cat", "/proc/cpuinfo").Output()
	if err != nil {
		log.Fatal(err)
	}

	out_reader := bytes.NewReader(out)

	scanner := bufio.NewScanner(out_reader)
	for scanner.Scan() {
		//fmt.Fprintln(w, re_redun_whtsp.ReplaceAllString(scanner.Text(), " "))
		fmt.Fprintln(w, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func upt_handler(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("cat", "/proc/uptime").Output()
	if err != nil {
		log.Fatal(err)
	}

	out_reader := bytes.NewReader(out)

	scanner := bufio.NewScanner(out_reader)
	for scanner.Scan() {
		fmt.Fprintln(w, re_redun_whtsp.ReplaceAllString(scanner.Text(), " "))
		//fmt.Fprintln(w, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var re_redun_whtsp *regexp.Regexp

func main() {

	re_redun_whtsp = regexp.MustCompile(`[\s\p{Zs}]{2,}`)

	http.HandleFunc("/", handler)
	http.HandleFunc("/cpuinfo", cpu_handler)
	http.HandleFunc("/meminfo", mem_handler)
	http.HandleFunc("/uptime", upt_handler)
	http.ListenAndServe(":1404", nil)
}
