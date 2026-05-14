// Package cli wires up the cobra command tree for `philips-wiz-bulb`.
package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/deepanshutr/philips-wiz-bulb-cli/internal/config"
	"github.com/deepanshutr/philips-wiz-bulb-cli/internal/core"
	"github.com/spf13/cobra"
)

func NewRoot() *cobra.Command {
	cfg := config.Load()
	c := core.New(cfg.CoreURL)
	ctx := context.Background()

	root := &cobra.Command{
		Use:   "philips-wiz-bulb",
		Short: "Control Philips WiZ bulbs via philips-wiz-bulb-core",
	}

	root.AddCommand(
		&cobra.Command{
			Use: "list", Short: "List known bulbs",
			RunE: func(*cobra.Command, []string) error {
				out, err := c.List(ctx)
				if err != nil {
					return err
				}
				b, _ := json.MarshalIndent(out, "", "  ")
				fmt.Println(string(b))
				return nil
			},
		},
		simple("state", "Show current state of a bulb", 0, 1, func(args []string) error {
			t := pickTarget(args)
			out, err := c.Get(ctx, t)
			if err != nil {
				return err
			}
			b, _ := json.MarshalIndent(out, "", "  ")
			fmt.Println(string(b))
			return nil
		}),
		simple("on", "Turn a bulb on", 0, 1, func(args []string) error {
			return c.On(ctx, pickTarget(args))
		}),
		simple("off", "Turn a bulb off", 0, 1, func(args []string) error {
			return c.Off(ctx, pickTarget(args))
		}),
		simple("bri", "Set brightness 10..100", 1, 2, func(args []string) error {
			v, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return c.Brightness(ctx, pickTargetAt(args, 1), v)
		}),
		simple("temp", "Set CCT 2200..6500 K", 1, 2, func(args []string) error {
			v, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return c.Temp(ctx, pickTargetAt(args, 1), v)
		}),
		simple("color", "Set RGB color: <r> <g> <b> [target]", 3, 4, func(args []string) error {
			r, _ := strconv.Atoi(args[0])
			g, _ := strconv.Atoi(args[1])
			b, _ := strconv.Atoi(args[2])
			return c.Color(ctx, pickTargetAt(args, 3), r, g, b)
		}),
		simple("scene", "Set built-in scene by name or id", 1, 2, func(args []string) error {
			return c.Scene(ctx, pickTargetAt(args, 1), args[0])
		}),
		&cobra.Command{
			Use: "discover", Short: "Force a LAN rescan",
			RunE: func(*cobra.Command, []string) error {
				out, err := c.Discover(ctx)
				if err != nil {
					return err
				}
				b, _ := json.MarshalIndent(out, "", "  ")
				fmt.Println(string(b))
				return nil
			},
		},
		&cobra.Command{
			Use: "name <target> <new-name>", Short: "Rename a bulb",
			Args: cobra.ExactArgs(2),
			RunE: func(_ *cobra.Command, args []string) error {
				return c.Rename(ctx, args[0], args[1])
			},
		},
	)
	return root
}

func pickTarget(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return "_default"
}

func pickTargetAt(args []string, idx int) string {
	if len(args) > idx {
		return args[idx]
	}
	return "_default"
}

func simple(use, short string, minArgs, maxArgs int, fn func([]string) error) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Args:  cobra.RangeArgs(minArgs, maxArgs),
		RunE:  func(_ *cobra.Command, args []string) error { return fn(args) },
	}
}
