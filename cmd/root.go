/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"os"

	dbInfra "github.com/DerivedPuma7/go-hexagonal/adapters/db"
	"github.com/DerivedPuma7/go-hexagonal/application"

	"github.com/spf13/cobra"
)

var db, _ = sql.Open("sqlite3", "db.sqlite")
var productDb = dbInfra.NewProductDb(db)
var productService = application.ProductService{Persistence: productDb}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-hexagonal",
	Short: "Simple application using golang and Ports and Adapters Architecture",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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


