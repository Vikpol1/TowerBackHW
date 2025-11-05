package uniq

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Options struct {
	Count          bool
	Duplicate      bool
	Unique         bool
	IgnoreRegister bool
	NumFields      int
	NumChars       int
}

func ProcessLines(str []string, opts Options) ([]string, error) {
	if len(str) == 0 {
		return []string{}, nil
	}

	res := make([]string, 0)
	processed := make([]string, len(str))
	for i, line := range str {
		processed[i] = processLine(line, opts)
	}
	i := 0
	for i < len(processed) {
		j := i + 1
		c := 1
		for j < len(processed) && processed[j] == processed[i] {
			c++
			j++
		}
		line := str[i]
		switch {
		case opts.Count:
			res = append(res, fmt.Sprintf("%d %s", c, line))
		case opts.Duplicate:
			if c > 1 {
				res = append(res, line)
			}
		case opts.Unique:
			if c == 1 {
				res = append(res, line)
			}
		default:
			res = append(res, line)
		}
		i = j
	}
	if res == nil {
		return []string{}, nil
	}
	return res, nil
}

func processLine(line string, opts Options) string {
	res := line
	if opts.NumFields > 0 {
		fields := strings.Fields(res)
		if len(fields) > opts.NumFields {
			res = strings.Join(fields[opts.NumFields:], " ")
		} else {
			res = ""
		}
	}
	if opts.NumChars > 0 {
		if len(res) > opts.NumChars {
			res = res[opts.NumChars:]
		} else {
			res = ""
		}
	}
	if opts.IgnoreRegister {
		res = strings.ToLower(res)
	}
	return res
}

func ReadLines(reader io.Reader) ([]string, error) {
	var str []string
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		str = append(str, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if str == nil {
		return []string{}, nil
	}
	return str, nil
}

func WriteLines(writer io.Writer, str []string) error {
	w := bufio.NewWriter(writer)
	for _, line := range str {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return w.Flush()
}
