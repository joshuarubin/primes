package main

import (
	"fmt"
	"os"
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
	}
}

func before(c *cli.Context) error {
	if len(c.Args()) < 2 {
		cli.ShowAppHelp(c)
		fmt.Fprintf(os.Stderr, "missing one or both integers\n")
		os.Exit(1)
	}

	for i := 0; i < 2; i++ {
		val, err := strconv.ParseUint(c.Args()[i], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing arg %d: %s\n", i, err)
			os.Exit(1)
		}
		args[i] = val
	}

	algopt := c.GlobalString("algorithm")
	switch algopt {
	case "eratosthenes", "":
		fmt.Fprintf(os.Stderr, "using sieve of eratosthenes algorithm\n")
		algo = primes.EratosthenesAlgo
	default:
		fmt.Fprintf(os.Stderr, "unknown algorithm: %s\n", algopt)
	}

	return nil
}

func run(c *cli.Context) {
	ps := primes.Between(args[0], args[1], algo)
	l := len(ps)
	if c.GlobalBool("print") && l > 0 {
		// fmt.Printf("%s\n", join(ps, ", "))
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
