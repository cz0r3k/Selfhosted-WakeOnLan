package main

import (
	"encoding/hex"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
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
	sendMagicPacker(mac, ip)
	w.WriteHeader(200)
}
func sendMagicPacker(mac string, ip string) {
	mp := createMagicPacket(mac)

	laddr, err := net.ResolveUDPAddr("udp", ":8089")
	if err != nil {
		log.Fatal(err)
	}

	raddr, err := net.ResolveUDPAddr("udp", ip+":9")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	conn.Write(mp)
}

func createMagicPacket(mac string) []byte {
	s := strings.Repeat("FF", 6) + strings.Repeat(strings.ReplaceAll(mac, ":", ""), 16)
	data, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
