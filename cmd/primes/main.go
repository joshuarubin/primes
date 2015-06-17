package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/joshuarubin/primes"
)

var (
	app  = cli.NewApp()
	args = make([]uint64, 2)
	algo primes.SieveAlgo
)

func init() {
	app.Name = "primes"
	app.Version = "1.0.0"
	app.Usage = "primes [-p] <int> <int>"
	app.Authors = []cli.Author{
		{Name: "Joshua Rubin", Email: "jrubin@zvelo.com"},
	}
	app.Before = before
	app.Action = run
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "print, p",
			Usage: "print the primes to stdout, comma separated",
		},
		cli.StringFlag{
			Name:  "algorithm, a",
			Usage: "which algorithm to use [eratosthenes]",
		},
		cli.StringFlag{
			Name:  "profile",
			Usage: "write profiling output to this file",
		},
	}
}

func before(c *cli.Context) error {
	if len(c.Args()) < 2 {
		cli.ShowAppHelp(c)
		log.Fatal("missing one or both integers")
	}

	for i := 0; i < 2; i++ {
		val, err := strconv.ParseUint(c.Args()[i], 10, 64)
		if err != nil {
			cli.ShowAppHelp(c)
			log.Fatalf("error parsing arg %d: %s\n", i, err)
		}
		args[i] = val
	}

	algopt := c.GlobalString("algorithm")
	switch algopt {
	case "eratosthenes", "":
		algo = primes.EratosthenesAlgo
	default:
		cli.ShowAppHelp(c)
		log.Fatalf("unknown algorithm: %s\n", algopt)
	}

	return nil
}

func run(c *cli.Context) {
	fn := c.GlobalString("profile")
	if len(fn) > 0 {
		f, err := os.Create(fn)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	ps := primes.Between(args[0], args[1], algo)
	l := len(ps)
	if c.GlobalBool("print") && l > 0 {
		for i, p := range ps {
			fmt.Printf("%d", p)
			if i < l-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Println()
	}
}

func main() {
	app.Run(os.Args)
}
