package usage

import (
	"flag"
	"fmt"
	"os"
)

func Ordered(names ...string) func() {
	return func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		for _, name := range names {
			f := *flag.Lookup(name)
			f.Value.Set(f.DefValue)
			fs := flag.NewFlagSet("", flag.PanicOnError)
			fs.Var(f.Value, f.Name, f.Usage)
			fs.PrintDefaults()
		}
	}
}
