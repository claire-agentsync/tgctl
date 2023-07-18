package main

import (
	"flag"
	"fmt"
	"os"
	"tgctl/scaffold"
)


func main() {
	// Scaffold flag functions
	scaffoldCmd := flag.NewFlagSet("scaffold", flag.ExitOnError)
	scaffoldAccount := scaffoldCmd.String("account", "", "Creates an account scaffold, e.g. 421929404181")
	scaffoldRegion := scaffoldCmd.String("region", "", "Scaffold region e.g. us-west-2")
	scaffoldService := scaffoldCmd.String("service", "", "Service that scaffold supports, e.g. ingress")
	scaffoldStack := scaffoldCmd.String("stack", "", "Stack that scaffold is a type of, e.g. service-stack")
	scaffoldKind := scaffoldCmd.String("kind", "", "Kind, e.g. networking, security, util")
	scaffoldResource := scaffoldCmd.String("resource", "", "Resource or module type, e.g. s3")
	scaffoldTarget := scaffoldCmd.String("target", ".", "Target directory")
	scaffoldDryRun := scaffoldCmd.Bool("dry-run", false, "Take no action but output what would happen")
	scaffoldHelp := scaffoldCmd.Bool("help", false, "Output help and exit")
	// Help functions
	//helpCmd := flag.NewFlagSet("help", flag.ExitOnError)
	// account/region/service/stack/kind/resource
	if len(os.Args) < 2 {
		Help()
		os.Exit(1)
		scaffoldCmd.PrintDefaults()
	}
	switch os.Args[1] {
	case "scaffold":
		scaffoldCmd.Parse(os.Args[2:])
		if *scaffoldHelp == true {
			scaffoldCmd.PrintDefaults()
			os.Exit(1)
		}
		s := scaffold.Scaff{Account: *scaffoldAccount,
			Region:   *scaffoldRegion,
			Service:  *scaffoldService,
			Stack:    *scaffoldStack,
			Kind:     *scaffoldKind,
			Resource: *scaffoldResource,
			Target:   *scaffoldTarget,
			DryRun:   *scaffoldDryRun,
		}
		dir, err := s.Build()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("success")
			fmt.Println(dir)
		}
        case "account":

	case "help":
		Help()
		scaffoldCmd.PrintDefaults()
		os.Exit(0)
	default:
		fmt.Println("Unknown usage of:", os.Args[1:])
		Help()
		os.Exit(1)
	}
	os.Exit(0)
}
