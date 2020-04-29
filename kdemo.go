package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/bitnami-labs/flagenv"
	"k8s.io/klog"
)

var (
	audience = flag.String("audience", "world", "Greeting audence.")
	addr     = flag.String("listen", ":8080", "HTTP server listen address.")
)

func handler(w http.ResponseWriter, req *http.Request) {
	klog.V(2)
	fmt.Fprintf(w, "I now say: Hello, %s!\n", *audience)
}

func main() {
	flagenv.SetFlagsFromEnv("KDEMO", flag.CommandLine)
	klog.InitFlags(nil)
	flag.Parse()

	http.HandleFunc("/", handler)
	klog.Infof("Listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
