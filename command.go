// Package command implements a command line parser that
// extends the stdlib flag package, adding support for
// commands, subcommands, positional parameters, and checked values.
package command

import (
	"errors"
	"fmt"
	"github.com/michaeljpetter/command/flag"
	"github.com/michaeljpetter/fp"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// Command represents a single command in a command tree, which may
// contain flags, subcommands, and positional parameters.
type Command struct {
	// The embedded FlagSet contains the name of the command, plus its flag definitions.
	*flag.FlagSet
	usage       string
	subcommands map[string]subcommand
	positional  []*positional

	// The behavior of Usage is analogous to FlagSet, but it extended by default to
	// display usage information for all flags, subcommands, and positional parameters.
	// Usage may be replaced with a user function to customize output.
	Usage func()
}

// Bound represents a [Command] that has been paired with a specific set
// of command arguments. It is returned from explicit calls
// to [Command.Bind], and implicitly in subcommand handlers.
type Bound struct {
	// The embedded command which has been bound.
	*Command
	args []string
}

// HandlerFunc defines a function that processes a [Bound] command.
// It is used during parsing to delegate to subcommands.
type HandlerFunc func(Bound)

// Value extends the [flag.Value] type, adding support for required values.
type Value interface {
	flag.Value
	Required() bool
}

type subcommand struct {
	usage   string
	handler HandlerFunc
}

type positional struct {
	name     string
	usage    string
	value    Value
	defValue string
}

// New creates a new [Command] with the given name, usage, and error handling.
func New(name, usage string, errorHandling flag.ErrorHandling) *Command {
	c := &Command{
		FlagSet:     flag.NewFlagSet(name, errorHandling),
		usage:       usage,
		subcommands: make(map[string]subcommand),
		positional:  make([]*positional, 0),
	}

	c.FlagSet.Usage = c.delegateUsage
	c.Usage = c.defaultUsage

	return c
}

// Subcommand defines a subcommand with the given name, usage, and handler.
// The handler is called only when the subcommand name has been parsed by this command,
// and is then bound to the remaining arguments.
//
// Panics if positional parameters have been defined on the same command,
// as they are mutually exclusive.
func (c *Command) Subcommand(name, usage string, handler HandlerFunc) {
	if c.HasPositional() {
		panic("subcommands and positional parameters are mutually exclusive")
	}

	c.subcommands[name] = subcommand{usage, handler}
}

// PositionalVar defines a positional parameter with the given [Value], name, and usage.
//
// Panics if subcommands have been defined on the same command,
// as they are mutually exclusive.
//
// Panics if the [Value] is required, and optional positional parameters
// have already been defined on the same command.
func (c *Command) PositionalVar(value Value, name, usage string) {
	if c.HasSubcommands() {
		panic("subcommands and positional parameters are mutually exclusive")
	}

	if c.HasPositional() && value.Required() && !c.positional[len(c.positional)-1].value.Required() {
		panic("required positional parameters must precede optional")
	}

	c.positional = append(c.positional, &positional{name, usage, value, value.String()})
}

// HasFlags indicates whether flags have been defined on this command.
func (c *Command) HasFlags() bool {
	has := false
	c.FlagSet.VisitAll(func(*flag.Flag) { has = true })
	return has
}

// HasSubcommands indicates whether subcommands have been defined on this command.
func (c *Command) HasSubcommands() bool {
	return 0 < len(c.subcommands)
}

// HasPositional indicates whether positional parameters have been defined on this command.
func (c *Command) HasPositional() bool {
	return 0 < len(c.positional)
}

func (c *Command) delegateUsage() {
	if c.Usage != nil {
		c.Usage()
	} else {
		c.defaultUsage()
	}
}

// PrintSubcommands prints, to standard error unless configured otherwise,
// the list of all defined subcommands and their usage strings.
func (c *Command) PrintSubcommands() {
	names := slices.Sorted(maps.Keys(c.subcommands))

	longest := fp.MaxOf(fp.StringLen, 4)(slices.Values(names))

	for _, name := range names {
		fmt.Fprintf(c.Output(), "  %-*s  %s\n", longest, name, c.subcommands[name].usage)
	}
}

// PrintPositional prints, to standard error unless configured otherwise,
// the list of all defined positional parameters and their usage strings.
func (c *Command) PrintPositional() {
	names := fp.Map(func(p *positional) string { return p.name })(slices.Values(c.positional))

	longest := fp.MaxOf(fp.StringLen, 4)(names)

	for _, positional := range c.positional {
		fmt.Fprintf(c.Output(), "  %-*s  %s", longest, positional.name, positional.usage)

		if !positional.value.Required() {
			fmt.Fprintf(c.Output(), " (default %s)", positional.defValue)
		}

		fmt.Fprint(c.Output(), "\n")
	}
}

func (c *Command) defaultUsage() {
	fmt.Fprintf(c.Output(), "Usage: %s", c.Name())

	if c.HasFlags() {
		fmt.Fprint(c.Output(), " [options]")
	}

	if c.HasSubcommands() {
		fmt.Fprint(c.Output(), " <command>")
	} else if c.HasPositional() {
		for _, positional := range c.positional {
			if positional.value.Required() {
				fmt.Fprintf(c.Output(), " <%s>", positional.name)
			} else {
				fmt.Fprintf(c.Output(), " [%s]", positional.name)
			}
		}
	}

	fmt.Fprintf(c.Output(), "\n\n")

	for _, line := range strings.Split(c.usage, "\n") {
		fmt.Fprintf(c.Output(), "  %s\n", line)
	}

	if c.HasFlags() {
		fmt.Fprint(c.Output(), "\nOptions:\n")
		c.FlagSet.PrintDefaults()
	}

	if c.HasSubcommands() {
		fmt.Fprint(c.Output(), "\nCommands:\n")
		c.PrintSubcommands()
	} else if c.HasPositional() {
		fmt.Fprint(c.Output(), "\nArguments:\n")
		c.PrintPositional()
	}
}

// Parse parses the given arguments according to the definition of the command.
// The behavior on error is defined by the [flag.ErrorHandling] value used to create the command.
func (c *Command) Parse(args []string) error {
	err := c.FlagSet.Parse(args)
	if err != nil {
		return err
	}

	if c.HasSubcommands() {
		err = c.parseCommand(c.FlagSet.Args())
	} else if c.HasPositional() {
		err = c.parsePositional(c.FlagSet.Args())
	}

	if err != nil {
		fmt.Fprintln(c.Output(), err)
		c.delegateUsage()

		switch c.ErrorHandling() {
		case flag.ContinueOnError:
			return err
		case flag.ExitOnError:
			os.Exit(2)
		case flag.PanicOnError:
			panic(err)
		}
	}

	return nil
}

func (c *Command) parseCommand(args []string) error {
	if 0 == len(args) {
		return errors.New("missing command")
	}

	name := strings.TrimSpace(args[0])
	subcommand, ok := c.subcommands[name]

	if !ok {
		return fmt.Errorf("unknown command: %s", name)
	}

	subcommand.handler(
		New(
			c.Name()+" "+name,
			subcommand.usage,
			c.ErrorHandling(),
		).
			Bind(args[1:]),
	)
	return nil
}

func (c *Command) parsePositional(args []string) error {
	for i, positional := range c.positional {
		if len(args) <= i {
			if positional.value.Required() {
				return fmt.Errorf("missing argument for <%s>", positional.name)
			}
			break
		}
		if err := positional.value.Set(args[i]); err != nil {
			return fmt.Errorf("invalid value \"%s\" for argument %s: %v", args[i], positional.name, err)
		}
	}
	return nil
}

// NArg returns the number of remaining arguments after parsing.
func (c *Command) NArg() int {
	if c.HasSubcommands() {
		return 0
	}
	if c.HasPositional() {
		return max(0, c.FlagSet.NArg()-len(c.positional))
	}
	return c.FlagSet.NArg()
}

// Arg provides indexed access to the remaining arguments after parsing.
// Arg returns an empty string if the requested element does not exist.
func (c *Command) Arg(i int) string {
	if c.HasSubcommands() {
		return ""
	}
	if c.HasPositional() {
		return c.FlagSet.Arg(i + len(c.positional))
	}
	return c.FlagSet.Arg(i)
}

// Args returns the remaining arguments after parsing.
func (c *Command) Args() []string {
	if c.HasSubcommands() {
		return nil
	}
	if c.HasPositional() {
		return c.FlagSet.Args()[min(len(c.FlagSet.Args()), len(c.positional)):]
	}
	return c.FlagSet.Args()
}

// Bind pairs this command with a specific set of arguments to be parsed,
// and returns a [Bound] command representing that pairing.
func (c *Command) Bind(args []string) Bound {
	return Bound{c, args}
}

// Parse calls [Command.Parse] with its bound arguments.
func (b Bound) Parse() error {
	return b.Command.Parse(b.args)
}

// Program creates a new top-level [Command] with a name extracted from
// [os.Args][0], the given usage, and [flag.ExitOnError] error handling.
// The result is then bound to [os.Args][1:].
//
// This is the typical starting point for most command-line processing.
func Program(usage string) Bound {
	return New(
		strings.TrimSuffix(
			filepath.Base(os.Args[0]),
			filepath.Ext(os.Args[0]),
		),
		usage,
		flag.ExitOnError,
	).
		Bind(os.Args[1:])
}
