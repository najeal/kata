package main

import (
	"log"

	"github.com/spf13/cobra"
)

var kata = &cobra.Command{
	Use:           "kata",
	Short:         "kata - kata exercises",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func main() {
	if err := kata.Execute(); err != nil {
		log.Fatal(err)
	}
}
