package main

import (
	"flag"
	"fmt"
	"myuniq/uniq"
	"os"
)

type UniqConfig struct {
	count          bool
	duplicate      bool
	unique         bool
	ignoreRegister bool
	numFields      int
	numChars       int
	inputFile      string
	outputFile     string
}

func parseFlags() UniqConfig {
	var uq UniqConfig
	flag.BoolVar(&uq.count, "c", false, "подсчитать количество встречаний")
	flag.BoolVar(&uq.duplicate, "d", false, "вывести только повторяющиеся строки")
	flag.BoolVar(&uq.unique, "u", false, "вывести только уникальные строки")
	flag.BoolVar(&uq.ignoreRegister, "i", false, "игнорировать регистр")
	flag.IntVar(&uq.numFields, "f", 0, "игнорировать первые N полей")
	flag.IntVar(&uq.numChars, "s", 0, "игнорировать первые N символов")

	flag.Parse()
	argv := flag.Args()
	if len(argv) == 1 {
		uq.inputFile = argv[0]
	}
	if len(argv) == 2 {
		uq.inputFile = argv[0]
		uq.outputFile = argv[1]
	}
	return uq
}

func validate(uq UniqConfig) {
	if (uq.count && uq.duplicate) || (uq.count && uq.unique) || (uq.unique && uq.duplicate) || (len(os.Args) > 6) {
		fmt.Fprintf(os.Stderr, "Ошибка! Утилита имеет вид: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		os.Exit(1)
	}
}

func openInputFile(filename string) *os.File {
	var input *os.File
	var err error
	if filename != "" {
		input, err = os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Возникла ошибка при открытии файла!")
			os.Exit(1)
		}
	} else {
		input = os.Stdin
	}
	return input
}

func createOutputFile(filename string) *os.File {
	var output *os.File
	var err error
	if filename != "" {
		output, err = os.Create(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Возникла ошибка при создании файла!")
			os.Exit(1)
		}
	} else {
		output = os.Stdout
	}
	return output
}

func main() {
	uq := parseFlags()
	validate(uq)
	input := openInputFile(uq.inputFile)
	defer input.Close()
	output := createOutputFile(uq.outputFile)
	defer output.Close()
	lines, err := uniq.ReadLines(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения: %v", err)
		os.Exit(1)
	}
	opts := uniq.Options{
		Count:          uq.count,
		Duplicate:      uq.duplicate,
		Unique:         uq.unique,
		IgnoreRegister: uq.ignoreRegister,
		NumFields:      uq.numFields,
		NumChars:       uq.numChars,
	}
	result, err := uniq.ProcessLines(lines, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка обработки: %v", err)
		os.Exit(1)
	}
	if err := uniq.WriteLines(output, result); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка записи: %v", err)
		os.Exit(1)
	}
}
