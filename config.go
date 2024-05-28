package version

import (
	"flag"
	"fmt"
	"github.com/bryanaustin/yaarp"
	"os"
)

type Config struct {
	Provided  *Version
	Greater   *Version
	Lesser    *Version
	Base      *Version
	Increment *Version
	Set       *Version
	Minimum   *Version
	Format    *Version
	Pad       *Version
}

// ConfigureFromArgs will automatically configure based on process arguments. On error it will print to stderr and exit.
func ConfigureFromArgs() *Config {
	greaterstr := flag.String("greater", "", "test to see if this version is greater (separators ignored)")
	lesserstr := flag.String("lesser", "", "test to see if this version is lesser (separators ignored)")
	basestr := flag.String("base", "", "increment by the largest value that is smaller than this value (separators ignored)")
	incrementstr := flag.String("increment", "", "increase the version but the amount provided (separators ignored)")
	setstr := flag.String("set", "", "set all non-zero numbers to the value provided (separators ignored)")
	minstr := flag.String("minimum", "", "pad the numbers by the integer provided (separators ignored)")
	formatstr := flag.String("format", "", "format result in the format provided (numbers ignored)")
	padstr := flag.String("pad", "", "pad the numbers by the integer provided (separators ignored)")

	yaarp.Parse()
	if len(yaarp.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Expected exactly one argument")
		os.Exit(1)
	}
	providedstr := yaarp.Arg(0)
	c := new(Config)
	c.Provided = mustParse(&providedstr)
	c.Greater = mustParse(greaterstr)
	c.Lesser = mustParse(lesserstr)
	c.Base = mustParse(basestr)
	c.Increment = mustParse(incrementstr)
	c.Set = mustParse(setstr)
	c.Minimum = mustParse(minstr)
	c.Format = mustParse(formatstr)
	c.Pad = mustParse(padstr)

	lessorgreat := c.Greater != nil || c.Lesser != nil
	if lessorgreat && (c.Base != nil || c.Increment != nil || c.Set != nil ||
		c.Minimum != nil || c.Format != nil || c.Pad != nil) {
		// Have a lesser or greater flag plus one other option
		fmt.Fprintln(os.Stderr, "The --lesser and --greater options don't work with any other arguments")
		os.Exit(2)
	}

	return c
}

func mustParse(x *string) *Version {
	if x == nil || len(*x) < 1 {
		return nil
	}
	v := Parse(*x)
	if err := v.Vaild(); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing %q: %s\n", *x, err)
		os.Exit(1)
	}
	return v
}
