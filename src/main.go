package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "chrnr-cli",
		Short: "k8s file churner cli",
		Long:  "A CLI tool to simulate file churn in a k8s cluster",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("chrnr-cli")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
