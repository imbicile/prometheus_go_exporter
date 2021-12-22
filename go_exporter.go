package main

import (
	"fmt"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"net/http"
	"os"
)

func main() {
	var (
		// exporter address and port
		addr = kingpin.Flag("listen-address", "Exporter address.").Default("0.0.0.0").String()
		port = kingpin.Flag("listen-port", "Exporter port.").Default("8080").String()
	)

	HelloServer := func(w http.ResponseWriter, r *http.Request) {

		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}

		sys_info := func(exp_name, exp_val string, exp_descr string) {
			fmt.Fprintln(w, "# HELP Debug "+exp_name+exp_descr)
			fmt.Fprintln(w, exp_name+"{nodename=\""+hostname+"\"} "+exp_val)
		}
		// print to web sysinfo
		sys_info("custom_exporter_addr", *addr, " Custom exporter running address.")
		sys_info("custom_exporter_port", *port, " Custom exporter running port.")
	}

	kingpin.CommandLine.UsageWriter(os.Stdout)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	fmt.Println("Server run at " + *addr + ":" + *port)
	http.HandleFunc("/metrics", HelloServer)
	http.ListenAndServe(*addr+":"+*port, nil)

}
