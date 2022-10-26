package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/wol", getWol)
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal("server closed\n")
	} else if err != nil {
		log.Fatal(err)
	}
}
func getWol(w http.ResponseWriter, r *http.Request) {
	ip := os.Getenv("IP")
	mac := os.Getenv("MAC")
	cmd := exec.Command("wakeonlan", "-i", ip, mac)
	stdout, err := cmd.Output()
	log.Printf("%s\n", stdout)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(200)
}
