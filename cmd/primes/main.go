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
)

func join(slice []uint64, sep string) string {
	ret := ""
	l := len(slice)
	for i, p := range slice {
		ret += fmt.Sprintf("%d", p)
		if i < l-1 {
			ret += ", "
		}
	}
	return ret

}

func init() {
	app.Name = "primes"
	app.Version = "1.0.0"
	app.Usage = "primes <int> <int>"
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
	}
}

func before(c *cli.Context) error {
	if len(c.Args()) < 2 {
		fmt.Println("need two integers")
		os.Exit(1)
	}

	for i := 0; i < 2; i++ {
		val, err := strconv.ParseUint(c.Args()[i], 10, 64)
		if err != nil {
			fmt.Printf("error parsing arg %d: %s\n", i, err)
			os.Exit(1)
		}
		args[i] = val
	}

	return nil
}

func run(c *cli.Context) {
	ps := primes.Between(args[0], args[1])
	if c.GlobalBool("print") && len(ps) > 0 {
		fmt.Printf("%s\n", join(ps, ", "))
	}
}

func main() {
	app.Run(os.Args)
}
