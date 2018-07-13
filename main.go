package main

import (
	"fmt"
	"log"
	"os"

	stan "github.com/nats-io/go-nats-streaming"
)

const (
	clusterID = "test-cluster"
	subject   = "scm.channel"
	//svrURL    = "nats://nats-streaming-local:4222"
	svrURL = "nats://192.168.99.101:31625"
)

func main() {
	fmt.Println("nats-streaming-client test")
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalf("could not get hostname: %v", err)
	}
	clientID := hn + "-stateful-set-client"
	fmt.Println(clientID)
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(svrURL),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalf("could not connect to steaming server: %v", err)
	}
	defer sc.Close()

	startOpt := stan.DeliverAllAvailable()
	i := 0
	_, err = sc.Subscribe(subject, func(msg *stan.Msg) {
		i++
	}, startOpt)
	if err != nil {
		log.Fatalf("could not subscribe to streaming server: %v", err)
	}
	fmt.Println("Fully initialized")
	select {}
}
