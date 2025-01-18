package main

import (
	"flag"
	"fmt"
	"gitlab.com/manifold555112/manifold/cli"
	"gitlab.com/manifold555112/manifold/server"
	"os"
)

func main() {
	cliCmd := flag.NewFlagSet("cli", flag.ExitOnError)
	cliSetup := cliCmd.Bool("S", false, "set up the system")
	cliCreateConfig := cliCmd.Bool("cc", false, "scaffold empty config file in $XDG_CONFIG_HOME/manifold/config.toml")

	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
	serverConf := serverCmd.String("conf", "", "config filpath, defaults to $XDG_CONFIG_HOME/manifold/config.toml")

	usage := func() {
		fmt.Println("Expected either subcommand cli or server")
		cliCmd.Usage()
		serverCmd.Usage()
	}

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "cli":
		if len(os.Args[2:]) == 0 {
			cliCmd.Usage()
			os.Exit(1)
		}
		cliCmd.Parse(os.Args[2:])
		err := cli.Run(cli.Args{
			Setup:        cliSetup,
			CreateConfig: cliCreateConfig,
		})
		if err != nil {
			cliCmd.Usage()
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	case "server":
		err := serverCmd.Parse(os.Args[2:])
		if err != nil {
			serverCmd.Usage()
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = server.StartServer(server.Args{
			Conf: serverConf,
		})
		if err != nil {
			serverCmd.Usage()
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	default:
		usage()
		os.Exit(1)
	}

}
