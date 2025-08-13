package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/rvfet/rich-go"
	"github.com/spf13/cobra"
)

func PrintLogo(phrase, font, color string) {
	logo := figure.NewColorFigure(phrase, font, color, true)
	logo.Print()
}

var rootCmd = &cobra.Command{
	Use:   "terminal-logos",
	Short: "terminal-logos is a CLI tool for printing ASCII-art logos to your terminal.",
	Long:  "terminal-logos is a CLI tool for printing beautiful and colorful ASCII-art logos to your terminal. See the supported ASCII art fonts here: https://github.com/common-nighthawk/go-figure?tab=readme-ov-file#supported-fonts",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing grab-those-docs '%s'\n", err)
		os.Exit(1)
	}
}

var phrase string
var color string
var font string

var printCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Create a logo with the specified phrase, color and font.",
	Long:    "Create a logo with the specified phrase (-p, --phrase flag), color (-c, --color flag) and font (-f, --font flag). Example usage: `terminal-logos create --phrase 'Hello Logos' --color 'cyan' --font 'roman'`",
	Run: func(cmd *cobra.Command, args []string) {
		if phrase == "" {
			fmt.Fprintf(os.Stderr, "Error: a phrase to be printed is required. Use -p or --phrase flag.\n")
			os.Exit(1)
		}

		colors := []string{
			"gray",
			"red",
			"green",
			"purple",
			"cyan",
			"white",
			"blue",
			"yellow",
		}

		if !slices.Contains(colors, color) {
			rich.Warning("Color " + color + " not available! Available colors: " + strings.Join(colors, ", "))
			rich.Warning("Defaulting to white...")
			color = "white"
		}

		PrintLogo(phrase, font, color)
	},
}

func init() {
	printCmd.Flags().StringVarP(&phrase, "phrase", "p", "", "The phrase to be printed as logo")
	printCmd.Flags().StringVarP(&color, "color", "c", "white", "The color to print the logo with")
	printCmd.Flags().StringVarP(&font, "font", "f", "alphabet", "The font to print the logo with")

	printCmd.MarkFlagRequired("phrase")

	rootCmd.AddCommand(printCmd)
}
