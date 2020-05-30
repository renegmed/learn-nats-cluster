package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	nats "github.com/nats-io/nats.go"
	"github.com/renegmed/driveragent/agent"
	"github.com/renegmed/driveragent/kit"
)

func main() {
	var (
		showHelp    bool
		showVersion bool
		natsServers string
		agentType   string
	)

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage: driver-agent [options...]\n\n")
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, "\n")
	}

	flag.BoolVar(&showHelp, "help", false, "Show help")
	flag.BoolVar(&showVersion, "version", false, "Show version")
	//flag.StringVar(&natsServers, "nats", "nats://172.20.0.3:4222,nats://172.20.0.4:4223,nats://172.20.02:4224", "List of NATS Servers to connect")
	flag.StringVar(&natsServers, "nats", "", "List of NATS Servers to connect") // nats cluster urls is required
	flag.StringVar(&agentType, "type", "regular", "Kind of vehicle")
	flag.Parse()

	switch {
	case showHelp:
		flag.Usage()
		os.Exit(0)
	case showVersion:
		fmt.Fprint(os.Stderr, "NATS Rider Driver Agent v%s\n", agent.Version)
		os.Exit(0)
	}
	log.Printf("Stating NATS Rider Driver Agent version %s", agent.Version)

	log.Println("nats servers:", natsServers)

	comp := kit.NewComponent("driver-agent")

	fmt.Print(comp)

	// Set infinite retries to never stop reconnecting
	err := comp.SetupConnectionToNATS(natsServers, nats.MaxReconnects(-1))
	if err != nil {
		log.Fatal(err)
	}

	ag := agent.Agent{
		Component: comp,
		AgentType: agentType,
	}

	err = ag.SetupSubscriptions()
	if err != nil {
		log.Fatal(err)
	}

	runtime.Goexit()

}
