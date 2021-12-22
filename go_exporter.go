package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	// exporter conf
	addr = "127.0.0.1"
	port = "8080"
)

func main() {
	fmt.Println("Server run at " + addr + ":" + port)
	http.HandleFunc("/metrics", HelloServer)
	http.ListenAndServe(addr+":"+port, nil)
}

func Out(export string) {

}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	sys_info := func(exp_name, exp_val string, exp_descr string) {
		fmt.Fprintln(w, "# HELP Debug "+exp_name+exp_descr)
		fmt.Fprintln(w, exp_name+"{nodename=\""+hostname+"\"} "+exp_val)
	}
	// print to web
	sys_info("custom_exporter_addr", addr, " Custom exporter running address.")
	sys_info("custom_exporter_port", port, " Custom exporter running port.")
}
