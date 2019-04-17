package main

import (
	"fmt"
	"log"
	"os"

	"github.com/najeal/kata/pkg/chain"

	"github.com/najeal/kata/pkg/common"
	"github.com/spf13/cobra"
)

const defaultChainsInput = "./inputchains.txt"

var chainsCmd = &cobra.Command{
	Use:   "chains",
	Short: "start anagram searches",
	RunE:  startChain,
}

var inputChainsFile string

func init() {
	chainsCmd.Flags().StringVarP(&inputChainsFile, "input", "i", defaultChainsInput, "Input File")
	kata.AddCommand(chainsCmd)
}

func startChain(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		log.Println("you should give 2 args: start word and end word")
	}
	fileReader := common.NewFileReader()
	store, err := chain.NewStore(args[0], args[1])
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to initiate word store")
	}
	file, err := os.OpenFile(inputChainsFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Chains: Failed to open file %v", inputChainsFile)
	}
	defer file.Close()
	fileReader.FeedFromFile(file, store.Add)
	graph := chain.NewMatrixGraph()
	store.FeedGraph(graph)
	res, err := graph.BFS(args[0], args[1])
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Chains: BFS search failed")
	}
	log.Println(res)
	return nil
}
