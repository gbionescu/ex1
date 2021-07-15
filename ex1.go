package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
)

const LISTEN_PORT = "8080"
const VERSION = "v0.0.1"

// Return the version number
func version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, VERSION)
}

func getDuration() ([]byte, error) {
	// Run systemd-analyze
	out, err := exec.Command("systemd-analyze", "time").Output()
	if err != nil {
		return []byte(""), errors.New("Could not get startup time due to error.")
	}

	// First line contains the startup time:
	// https://github.com/systemd/systemd/blob/main/src/analyze/analyze.c#L524
	// With the typical output looking like:
	// Startup finished in 2.960s (kernel) + 12.078s (userspace) = 15.038s
	// graphical.target reached after 12.069s in userspace
	//
	// Source code also indicates that more elements can be found in the string:
	// D1 (firmware) D2 (loader) D3 (kernel) D4 (initrd) = sum(D1 + D2 + D3 + D4)
	//
	// Assuming here that we only want the sum (e.g. what's after '=').
	out_split := bytes.Split(out, []byte("\n"))

	if len(out_split) == 0 {
		return []byte(""), errors.New("could not split systemd-analyze output.")
	}

	// Get the string after ' = '
	first_line_split := bytes.Split(out_split[0], []byte(" = "))

	if len(first_line_split) == 0 {
		return []byte(""), errors.New("could not find the total duration.")
	}

	// Last element is the duration
	return first_line_split[len(first_line_split)-1], nil
}

// Return the duration
func duration(w http.ResponseWriter, r *http.Request) {
	startup_duration, err := getDuration()

	if err != nil {
		fmt.Fprintf(w, "Error getting the startup time:\n%s", err)
	}

	fmt.Fprintf(w, "%s", startup_duration)
}

func main() {
	// Set up handlers for each endpoint.
	http.HandleFunc("/version", version)
	http.HandleFunc("/duration", duration)

	// Open the port to make sure that we can create the server here.
	listener, err := net.Listen("tcp", "localhost:"+LISTEN_PORT)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server ready, endpoints: /version and /duration")

	log.Fatal(http.Serve(listener, nil))
}
