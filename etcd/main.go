package main

import (
	"context"
	"flag"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	run := func(ctx context.Context) {
		// complete your controller loop here
		log.Println("housekeeping...")

		select {}
	}

	var name = flag.String("name", "", "give a name")
	flag.Parse()

	// Create a etcd client
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"localhost:2379"}})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// create a sessions to elect a Leader
	session, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	election := concurrency.NewElection(session, "/leader-election/")

	// use a Go context, so we can tell the leaderelection code when we
	// want to step down
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// listen for interrupts or the Linux SIGTERM signal and resign,
	// which the leader election code will observe and
	// step down
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		log.Println("Received termination, signaling shutdown")
		election.Resign(ctx)
	}()

	// Elect a leader (or wait that the leader resign)
	if err := election.Campaign(ctx, *name); err != nil {
		log.Fatal(err)
	}

	// run the controller loop when elected
	run(ctx)
}
