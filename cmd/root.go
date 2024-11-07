package cmd

import (
	"fmt"
	"os"

	"github.com/raghavendra-talur/yeshto/internal/ast"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yeshto",
	Short: "yeshto is a CLI tool to provide insights on your Go repositories",
	Run: func(cmd *cobra.Command, args []string) {
		// get current working directory
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		modInfo, err := ast.BuildModuleInfo(cwd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, pkg := range modInfo.Packages {

			fmt.Printf("Package: %s\n", pkg.Name)
			fmt.Printf("    Path: %s\n", pkg.PkgPath)

			if pkg.Module != nil {
				fmt.Printf("    Module Path: %s\n", pkg.Module.Path)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
