package main

import (
	"fmt"
	"log"
	"os"

	"github.com/najeal/kata/pkg/anagrams"
	"github.com/najeal/kata/pkg/common"
	"github.com/spf13/cobra"
)

const defaultInput = "./example_files/kata_anagrams.txt"

var inputFileAnagrams string

var anagramsCmd = &cobra.Command{
	Use:   "anagrams",
	Short: "start anagram searches",
	RunE:  startAnagrams,
}

func init() {
	anagramsCmd.Flags().StringVarP(&inputFileAnagrams, "file", "f", defaultInput, "Input File")
	kata.AddCommand(anagramsCmd)
}

func startAnagrams(cmd *cobra.Command, args []string) error {
	store := anagrams.NewSizeDispatcher(new(anagrams.LenExtractor), new(common.WordCleaner))
	feeder := new(common.FileReader)
	printer := new(anagrams.LoggerPrinter)
	file, err := os.OpenFile(inputFileAnagrams, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Anagrams: Failed to open file %v", inputFileAnagrams)
	}
	defer file.Close()
	feeder.FeedFromFile(file, store.Add)
	store.PrintAll(printer)
	store.PrintLongestWordsAnagrams(printer)
	store.PrintLongestSet(printer)
	store.PrintTotalWords(printer)
	store.PrintTotalSets(printer)
	return nil
}
