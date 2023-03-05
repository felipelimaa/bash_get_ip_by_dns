package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// This func will return the application ready to execute.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Find ip and names on internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for IP addresses by host on the internet.",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "servers",
			Usage:  "Search for servers by host on the internet.",
			Flags:  flags,
			Action: findServers,
		},
	}

	return app
}

func findIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServers(c *cli.Context) {
	host := c.String("host")

	hosts, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, host := range hosts {
		fmt.Println(host.Host)
	}

}
