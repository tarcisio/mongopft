package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"time"

	mongopft "github.com/tarcisio/mongopft/pkg"

	"golang.org/x/sync/errgroup"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fHostname := flag.String("h", "169.57.230.40", "Hostname")
	fPort := flag.String("p", "30556", "Porta")
	fDatabase := flag.String("d", "db01", "Banco de dados utilizado no teste")
	fCollection := flag.String("c", "collection01", "collection utilizado no teste")
	nThreads := flag.Int("n", 1, "Numero de threads")
	flag.Parse()

	fmt.Printf("Database: %s\n", *fDatabase)
	fmt.Printf("Collection: %s\n", *fCollection)
	fmt.Printf("Numero de threads: %d\n", *nThreads)

	ctx := context.Background()

	var eg errgroup.Group
	for i := 0; i < *nThreads; i++ {
		eg.Go(func() error {
			return mongopft.TestThread(ctx, *fHostname, *fPort)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println("Encountered error:", err)
	}
	fmt.Println("Successfully finished.")
}
