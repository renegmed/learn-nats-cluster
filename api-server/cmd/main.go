package main

import (
	"flag"
	"log"
	"os"
	"runtime"

	"github.com/renegmed/riders-client/kit"
	"github.com/renegmed/riders-client/server"
)

func main() {
	var (
		showHelp     bool
		showVersion  bool
		serverListen string
		natsServers  string
	)
	flag.Usage = func() {
		log.Print(os.Stderr, "Usage: api-server [options...]\n\n")
		flag.PrintDefaults()
		log.Print(os.Stderr, "\n")
	}

	// Set default flags
	flag.BoolVar(&showHelp, "help", false, "Show help")
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.StringVar(&serverListen, "listen", "0.0.0.0:9090", "Network host:port to listen on")
	flag.StringVar(&natsServers, "nats", "", "List of NATS Servers to connect") // required to pass nats URLs

	flag.Parse()

	switch {
	case showHelp:
		flag.Usage()
		os.Exit(0)
	case showVersion:
		log.Print(os.Stderr, "NATS Rider API Server v%s\n", server.Version)
		os.Exit(0)
	}

	log.Printf("Starting NATS Rider API Server version %s", server.Version)

	// Register new component within the system
	comp := kit.NewComponent("api-server")

	log.Println("nats servers:", natsServers)

	err := comp.SetupConnectionToNATS(natsServers)
	if err != nil {
		log.Fatalf("Error on setting up NATS connection:", err)
	}

	log.Print(comp)

	server := server.NewServer(comp)

	err = server.ListenAndServe(serverListen)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening for HTTP request on %v", serverListen)
	runtime.Goexit()
}
