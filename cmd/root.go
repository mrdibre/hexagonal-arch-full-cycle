package cmd

import (
	"database/sql"
	"github.com/mrdibre/hexagonal-arch-go/adapters/db"
	"github.com/mrdibre/hexagonal-arch-go/application"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hexagonal-arch-go",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var conn, _ = sql.Open("sqlite3", "db.sqlite")
var productDb = db.NewProductDb(conn)
var productService = application.ProductService{Persistence: productDb}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


