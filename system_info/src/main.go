package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var endpoints = []string{"cpuinfo", "meminfo", "uptime"}

func handler(w http.ResponseWriter, r *http.Request) {

	target := r.URL.Path[1:]

	for _, e := range endpoints {
		if target == e {
			out, err := exec.Command("cat", "/proc/"+r.URL.Path[1:]).Output()
			if err != nil {
				log.Fatal(err)
			}

			out_hash := make(map[string]string)
			out_reader := bytes.NewReader(out)
			scanner := bufio.NewScanner(out_reader)
			for scanner.Scan() {
				var kv_arr []string
				line := scanner.Text()

				switch target {
				case "cpuinfo":
					line = strings.Replace(line, "  ", "", -1)
					line = strings.Replace(line, "\t", "", -1)
					line = strings.Replace(line, ": ", ":", -1)
					kv_arr = strings.Split(line, ":")

					if kv_arr[0] != "" {
						if len(kv_arr) > 1 {
							out_hash[kv_arr[0]] = kv_arr[1]
						} else {
							out_hash[kv_arr[0]] = ""
						}
					}
				case "meminfo":
					line = strings.Replace(line, "  ", "", -1)
					line = strings.Replace(line, ": ", ":", -1)
					kv_arr = strings.Split(line, ":")

					if kv_arr[0] != "" {
						if len(kv_arr) > 1 {
							out_hash[kv_arr[0]] = kv_arr[1]
						} else {
							out_hash[kv_arr[0]] = ""
						}
					}
				case "uptime":
					out_hash["uptime"] = strings.Split(line, " ")[0]
					out_hash["idle"] = strings.Split(line, " ")[1]
				}

				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
			}

			json_bytes, err := json.Marshal(out_hash)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintln(w, string(json_bytes))
			return
		}
	}

	for _, e := range endpoints {
		fmt.Fprintln(w, "<a href='/"+e+"'>"+e+"</a>\n"+"<br>\n")
	}
	return
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":1404", nil)
}
