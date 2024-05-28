package seeds

import (
	"flag"
	"fmt"
)

func Seeds() bool {
	var seedVal string
	flag.StringVar(&seedVal, "seed", "", "Get value flag --seed={seeder}")
	flag.Parse()
	switch seedVal {
	case "user":
		fmt.Print("Creating...\n")
		UserSeed()
		fmt.Print("User Created")
		return true
	}

	return false
}
