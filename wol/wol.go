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
	web_port := os.Getenv("WEB_PORT")
	sending_port := os.Getenv("SENDING_PORT")
	receiving_port := os.Getenv("RECEIVING_PORT")
	broadcast_ip := os.Getenv("BROADCAST_IP")
	mac := os.Getenv("MAC")

	magic_packet := createMagicPacket(mac)
	local_addr, err := net.ResolveUDPAddr("udp", (":" + sending_port))
	if err != nil {
		log.Fatal(err)
	}
	remote_addr, err := net.ResolveUDPAddr("udp", (broadcast_ip + ":" + receiving_port))
	if err != nil {
		log.Fatal(err)
	}

	arg := wol_arg{
		local_addr:   local_addr,
		remote_addr:  remote_addr,
		magic_packet: magic_packet}

	http.Handle("/", fs)
	http.HandleFunc("/wol", arg.getWol)
	err_server := http.ListenAndServe(":"+web_port, nil)
	if errors.Is(err_server, http.ErrServerClosed) {
		log.Fatal("server closed\n")
	} else if err_server != nil {
		log.Fatal(err_server)
	}
}

type wol_arg struct {
	local_addr   *net.UDPAddr
	remote_addr  *net.UDPAddr
	magic_packet []byte
}

func (arg wol_arg) getWol(w http.ResponseWriter, r *http.Request) {
	sendMagicPacket(arg)
	w.WriteHeader(200)
}

func sendMagicPacket(arg wol_arg) {
	conn, err := net.DialUDP("udp", arg.local_addr, arg.remote_addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	conn.Write(arg.magic_packet)
}

func createMagicPacket(mac string) []byte {
	s := strings.Repeat("FF", 6) + strings.Repeat(strings.ReplaceAll(mac, ":", ""), 16)
	data, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
