package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/wol", getWol)
	err := http.ListenAndServe(":8888", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal("server closed\n")
	} else if err != nil {
		log.Fatal(err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "HELLO!\n")
}
func getWol(w http.ResponseWriter, r *http.Request) {
	ip := os.Getenv("IP")
	mac := os.Getenv("MAC")
	cmd := exec.Command("wakeonlan", "-i", ip, mac)
	stdout, err := cmd.Output()
	fmt.Printf("%s\n", stdout)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, "Hello, WOL!\n")
}
