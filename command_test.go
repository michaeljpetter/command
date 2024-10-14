package command_test

import (
	"bytes"
	"github.com/michaeljpetter/command"
	"github.com/michaeljpetter/command/check"
	"github.com/michaeljpetter/command/flag"
	"github.com/michaeljpetter/ptr"
	"io"
	"slices"
	"testing"
)

func usageString(c *command.Command) string {
	buf := new(bytes.Buffer)
	out := c.Output()
	c.SetOutput(buf)
	c.Usage()
	c.SetOutput(out)
	return buf.String()
}

func TestCommandEmpty(t *testing.T) {
	cmd := command.New("empty", "this is only a test", flag.ContinueOnError)

	if cmd.Name() != "empty" {
		t.Errorf("wrong name %v, expected %v", cmd.Name(), "empty")
	}

	if usageString(cmd) !=
		`Usage: empty

  this is only a test
` {
		t.Errorf("wrong usage:\n%v", usageString(cmd))
	}

	if cmd.HasFlags() {
		t.Error("HasFlags returned true")
	}
	if cmd.HasSubcommands() {
		t.Error("HasSubcommands returned true")
	}
	if cmd.HasPositional() {
		t.Error("HasPositional returned true")
	}

	args := []string{"some", "thing"}
	err := cmd.Parse(args)

	if err != nil {
		t.Fatalf("parse failed with %v", err)
	}
	if !cmd.Parsed() {
		t.Error("parsed returned false")
	}
	if !slices.Equal(cmd.Args(), args) {
		t.Errorf("wrong args %v, expected %v", cmd.Args(), args)
	}
}

func TestCommandFlags(t *testing.T) {
	var f int
	var cab string

	buildCommand := func() *command.Command {
		f, cab = 0, ""
		cmd := command.New("trucker", "make a ford truck", flag.ContinueOnError)
		cmd.IntVar(&f, "f", 150, "model num", check.AtLeast(150))
		cmd.StringVar(&cab, "cab", "", "cab feature", check.OneOf("extended", "super"))
		return cmd
	}

	t.Run("Info", func(t *testing.T) {
		cmd := buildCommand()

		if usageString(cmd) !=
			`Usage: trucker [options]

  make a ford truck

Options:
  -cab value
    	cab feature
  -f value
    	model num (default 150)
` {
			t.Errorf("wrong usage:\n%v", usageString(cmd))
		}

		if !cmd.HasFlags() {
			t.Error("HasFlags returned false")
		}
		if cmd.HasSubcommands() {
			t.Error("HasSubcommands returned true")
		}
		if cmd.HasPositional() {
			t.Error("HasPositional returned true")
		}
	})

	t.Run("ValidArgs", func(t *testing.T) {
		cmd := buildCommand()
		err := cmd.Parse([]string{"-f", "250", "-cab", "super", "wut?"})

		if err != nil {
			t.Fatalf("parse failed with %v", err)
		}
		if f != 250 {
			t.Errorf("wrong -f value %v, expected %v", f, 250)
		}
		if cab != "super" {
			t.Errorf("wrong -cab value %v, expected %v", cab, "super")
		}
		if !slices.Equal(cmd.Args(), []string{"wut?"}) {
			t.Errorf("wrong args %v, expected %v", cmd.Args(), []string{"wut?"})
		}
	})

	t.Run("NoArgs", func(t *testing.T) {
		cmd := buildCommand()
		err := cmd.Parse(nil)

		if err != nil {
			t.Fatalf("parse failed with %v", err)
		}
		if f != 150 {
			t.Errorf("wrong -f value %v, expected %v", f, 150)
		}
		if cab != "" {
			t.Errorf("wrong -cab value %v, expected empty", cab)
		}
		if cmd.NArg() != 0 {
			t.Errorf("wrong args %v, expected empty", cmd.Args())
		}
	})

	t.Run("ArgFailsParse", func(t *testing.T) {
		cmd := buildCommand()
		cmd.SetOutput(io.Discard)
		err := cmd.Parse([]string{"-f", "ancy"})

		if err == nil {
			t.Fatal("parse succeeded")
		}
		if err.Error() != `invalid value "ancy" for flag -f: parse error` {
			t.Errorf("wrong error %v", err)
		}
	})

	t.Run("ArgFailsCheck", func(t *testing.T) {
		cmd := buildCommand()
		cmd.SetOutput(io.Discard)
		err := cmd.Parse([]string{"-cab", "brown"})

		if err == nil {
			t.Fatal("parse succeeded")
		}
		if err.Error() != `invalid value "brown" for flag -cab: must be one of [extended super]` {
			t.Errorf("wrong error %v", err)
		}
	})
}

func TestCommandPositional(t *testing.T) {
	var f int
	var cab string

	buildCommand := func() *command.Command {
		f, cab = 0, ""
		cmd := command.New("trucker", "make a ford truck", flag.ContinueOnError)
		cmd.PositionalIntVar(&f, "f", nil, "model num", check.AtLeast(150))
		cmd.PositionalStringVar(&cab, "cab", ptr.To(""), "cab feature", check.OneOf("extended", "super"))
		return cmd
	}

	t.Run("Info", func(t *testing.T) {
		cmd := buildCommand()

		if usageString(cmd) !=
			`Usage: trucker <f> [cab]

  make a ford truck

Arguments:
  f     model num
  cab   cab feature (default "")
` {
			t.Errorf("wrong usage:\n%v", usageString(cmd))
		}

		if cmd.HasFlags() {
			t.Error("HasFlags returned true")
		}
		if cmd.HasSubcommands() {
			t.Error("HasSubcommands returned true")
		}
		if !cmd.HasPositional() {
			t.Error("HasPositional returned false")
		}
	})

	t.Run("ValidArgs", func(t *testing.T) {
		cmd := buildCommand()
		err := cmd.Parse([]string{"250", "super", "wut?"})

		if err != nil {
			t.Fatalf("parse failed with %v", err)
		}
		if f != 250 {
			t.Errorf("wrong f value %v, expected %v", f, 250)
		}
		if cab != "super" {
			t.Errorf("wrong cab value %v, expected %v", cab, "super")
		}
		if !slices.Equal(cmd.Args(), []string{"wut?"}) {
			t.Errorf("wrong args %v, expected %v", cmd.Args(), []string{"wut?"})
		}
	})

	t.Run("RequiredOnlyArgs", func(t *testing.T) {
		cmd := buildCommand()
		err := cmd.Parse([]string{"350"})

		if err != nil {
			t.Fatalf("parse failed with %v", err)
		}
		if f != 350 {
			t.Errorf("wrong -f value %v, expected %v", f, 350)
		}
		if cab != "" {
			t.Errorf("wrong -cab value %v, expected empty", cab)
		}
		if cmd.NArg() != 0 {
			t.Errorf("wrong args %v, expected empty", cmd.Args())
		}
	})

	t.Run("NoArgs", func(t *testing.T) {
		cmd := buildCommand()
		cmd.SetOutput(io.Discard)
		err := cmd.Parse(nil)

		if err == nil {
			t.Fatal("parse succeeded")
		}
		if err.Error() != `missing argument for <f>` {
			t.Errorf("wrong error %v", err)
		}
	})

	t.Run("ArgFailsParse", func(t *testing.T) {
		cmd := buildCommand()
		cmd.SetOutput(io.Discard)
		err := cmd.Parse([]string{"ancy"})

		if err == nil {
			t.Fatal("parse succeeded")
		}
		if err.Error() != `invalid value "ancy" for argument f: parse error` {
			t.Errorf("wrong error %v", err)
		}
	})

	t.Run("ArgFailsCheck", func(t *testing.T) {
		cmd := buildCommand()
		cmd.SetOutput(io.Discard)
		err := cmd.Parse([]string{"50"})

		if err == nil {
			t.Fatal("parse succeeded")
		}
		if err.Error() != `invalid value "50" for argument f: must be at least 150` {
			t.Errorf("wrong error %v", err)
		}
	})
}

func TestCommandSubcommands(t *testing.T) {
	var designCalled, buyCalled bool

	buildCommand := func() *command.Command {
		designCalled, buyCalled = false, false
		cmd := command.New("trucker", "truck utility", flag.ContinueOnError)
		cmd.Subcommand("design", "design a new truck", func(cmd command.Bound) {
			designCalled = true
		})
		cmd.Subcommand("buy", "buy a stock truck", func(cmd command.Bound) {
			buyCalled = true
		})
		return cmd
	}

	t.Run("Info", func(t *testing.T) {
		cmd := buildCommand()

		if usageString(cmd) !=
			`Usage: trucker <command>

  truck utility

Commands:
  buy     buy a stock truck
  design  design a new truck
` {
			t.Errorf("wrong usage:\n%v", usageString(cmd))
		}

		if cmd.HasFlags() {
			t.Error("HasFlags returned true")
		}
		if !cmd.HasSubcommands() {
			t.Error("HasSubcommands returned false")
		}
		if cmd.HasPositional() {
			t.Error("HasPositional returned true")
		}
		if designCalled || buyCalled {
			t.Error("called handler(s) before parse")
		}
	})

	t.Run("ValidCommand", func(t *testing.T) {
		cmd := buildCommand()
		err := cmd.Parse([]string{"buy", "f150"})

		if err != nil {
			t.Fatalf("parse failed with %v", err)
		}
		if cmd.NArg() != 0 {
			t.Errorf("wrong args %v, expected empty", cmd.Args())
		}
		if !buyCalled {
			t.Error("did not call buy handler")
		}
		if designCalled {
			t.Error("called design handler")
		}
	})

	t.Run("NoCommand", func(t *testing.T) {
		cmd := buildCommand()
		cmd.SetOutput(io.Discard)
		err := cmd.Parse(nil)

		if err == nil {
			t.Fatal("parse succeeded")
		}
		if err.Error() != `missing command` {
			t.Errorf("wrong error %v", err)
		}
		if designCalled || buyCalled {
			t.Error("called handler(s)")
		}
	})

	t.Run("InvalidCommand", func(t *testing.T) {
		cmd := buildCommand()
		cmd.SetOutput(io.Discard)
		err := cmd.Parse([]string{"impound"})

		if err == nil {
			t.Fatal("parse succeeded")
		}
		if err.Error() != `unknown command: impound` {
			t.Errorf("wrong error %v", err)
		}
		if designCalled || buyCalled {
			t.Error("called handler(s)")
		}
	})
}

func TestCommandAll(t *testing.T) {
	calledHandler := false

	cmd := command.New("trucker", "truck utility", flag.ContinueOnError)
	make := cmd.String("make", "FORD", "truck manufacturer")

	cmd.Subcommand("design", "design a new truck", func(cmd command.Bound) {
		calledHandler = true
		budget := cmd.Int("budget", 50000, "design budget")
		model := cmd.PositionalString("model", nil, "truck model")

		if usageString(cmd.Command) !=
			`Usage: trucker design [options] <model>

  design a new truck

Options:
  -budget value
    	design budget (default 50000)

Arguments:
  model  truck model
` {
			t.Errorf("wrong usage:\n%v", usageString(cmd.Command))
		}

		err := cmd.Parse()

		if err != nil {
			t.Fatalf("parse failed with %v", err)
		}

		if *make != "CHEVY" {
			t.Errorf("wrong make %v, expected %v", *make, "CHEVY")
		}
		if *model != "Silverado" {
			t.Errorf("wrong model %v, expected %v", *model, "Silverado")
		}
		if *budget != 45000 {
			t.Errorf("wrong budget %v, expected %v", *budget, 45000)
		}
	})

	cmd.Subcommand("buy", "buy a stock truck", func(cmd command.Bound) {
		t.Fatal("called buy handler")
	})

	if usageString(cmd) !=
		`Usage: trucker [options] <command>

  truck utility

Options:
  -make value
    	truck manufacturer (default "FORD")

Commands:
  buy     buy a stock truck
  design  design a new truck
` {
		t.Errorf("wrong usage:\n%v", usageString(cmd))
	}

	err := cmd.Parse([]string{"-make", "CHEVY", "design", "-budget", "45000", "Silverado"})

	if err != nil {
		t.Fatalf("parse failed with %v", err)
	}
	if !calledHandler {
		t.Error("did not call handler")
	}
}
