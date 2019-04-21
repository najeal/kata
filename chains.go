package main

import (
	"fmt"
	"log"
	"os"

	"github.com/najeal/kata/pkg/chain"

	"github.com/najeal/kata/pkg/common"
	"github.com/spf13/cobra"
)

const defaultChainsInput = "./example_files/wordlist.txt"

var inputFileChains string

var chainsCmd = &cobra.Command{
	Use:   "chains",
	Short: "start anagram searches",
	RunE:  startChain,
}

func init() {
	chainsCmd.Flags().StringVarP(&inputFileChains, "file", "f", defaultChainsInput, "Input File")
	kata.AddCommand(chainsCmd)
}

func startChain(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		log.Println("you should give 2 args: start word and end word")
		return fmt.Errorf("not enough argument")
	}
	fileReader := common.NewFileReader()
	store, err := chain.NewStore(args[0], args[1])
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to initiate word store")
	}
	file, err := os.OpenFile(inputFileChains, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Chains: Failed to open file %v", inputFileChains)
	}
	defer file.Close()
	fileReader.FeedFromFile(file, store.Add)
	if store.AreReferencesIn() != true {
		log.Println("At least one of your words is not in the dictionnary")
		return fmt.Errorf("Cannot search chain")
	}
	graph := chain.NewMatrixGraph()
	store.FeedGraph(graph)
	res, err := graph.BFS(args[0], args[1])
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Chains: BFS search failed")
	}
	fmt.Println(res)
	return nil
}
