package main

import (
	"flag"
	"log"
	"os"

	"inet.af/tcpproxy"
)

var (
	listen string
	target string
)

func init() {
	flag.StringVar(&listen, "listen", "", "listen address (required)")
	flag.StringVar(&target, "target", "", "target address (required)")
}

func main() {
	flag.Parse()

	if listen == "" || target == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var p tcpproxy.Proxy
	p.AddRoute(listen, tcpproxy.To(target))

	log.Printf("listen=%s, target=%s", listen, target)
	if err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
