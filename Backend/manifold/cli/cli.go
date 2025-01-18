package cli

import (
	"encoding/json"
	"fmt"
	"gitlab.com/manifold555112/manifold/lib"
)

func Run(args Args) error {
	argstr, err := json.MarshalIndent(args, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(argstr))
	if args.CreateConfig != nil && *args.CreateConfig {
		fmt.Println("creating default config")
		if err := lib.CreateDefaultConfig(); err != nil {
			return err
		}
	}
	if args.Setup != nil && *args.Setup {
		if err := setup(); err != nil {
			return err
		}
	}
	return nil
}
