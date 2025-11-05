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

func main() {
	uq := parseFlags()
	if (uq.count && uq.duplicate) || (uq.count && uq.unique) || (uq.unique && uq.duplicate) || (len(os.Args) > 6) {
		fmt.Fprintf(os.Stderr, "Ошибка! Утилита имеет вид: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		os.Exit(1)
	}

	var input *os.File
	if uq.inputFile != "" {
		var err error
		input, err = os.Open(uq.inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Возникла ошибка при открытии файла!")
			os.Exit(1)
		}
		defer input.Close()
	} else {
		input = os.Stdin
	}

	var output *os.File
	if uq.outputFile != "" {
		var err error
		output, err = os.Create(uq.outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Возникла ошибка при создании файла!")
			os.Exit(1)
		}
		defer output.Close()
	} else {
		output = os.Stdout
	}

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
