package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	interpolate "interpolate/internal"
	"interpolate/internal/util"
	"io/ioutil"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "ip",
	Short: "Easily interpolate variables based on go templates!",
	Long:  ``,
	RunE:  run,
}

var argVars = make(map[string]string)
var argFiles = make([]string, 5)

func init() {
	cobra.OnInitialize()

	rootCmd.Flags().StringToStringVarP(&argVars, "var", "v", nil, "Key-value pair. Multiple are allowed. E.g. -v URL=http://google.com or --var timeout=50")
	rootCmd.Flags().StringArrayVarP(&argFiles, "file", "f", nil, "Path to a file containing variables. Multiple are allowed. Currently supported formats: env.")
}

// Boilerplate
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) (err error) {
	vars, err := aggregateVars()
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	return interpolate.Execute(vars, string(bytes))
}

func aggregateVars() (finalVars map[string]string, err error) {
	finalVars = make(map[string]string)

	for _, e := range os.Environ() {
		envVar := strings.SplitN(e, "=", 2)
		finalVars[envVar[0]] = envVar[1]
	}

	for _, file := range argFiles {
		fileVars, err := godotenv.Read(file)
		if err != nil {
			return nil, err
		}

		util.Merge(finalVars, fileVars)
	}

	util.Merge(finalVars, argVars)

	return finalVars, nil
}
