package io

import (
	"bufio"
	"os"
)

func Read_file(filename string, source_file *string, text_buffer *[][]rune) {
	file, err := os.Open(filename)
	if err != nil {
		*source_file = filename
		*text_buffer = append(*text_buffer, []rune{})
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	line_number := 0

	for scanner.Scan() {
		line := scanner.Text()
		*text_buffer = append(*text_buffer, []rune{})

		for i := 0; i < len(line); i++ {
			(*text_buffer)[line_number] = append((*text_buffer)[line_number], rune(line[i]))
		}
		line_number++
	}
	if line_number == 0 {
		*text_buffer = append(*text_buffer, []rune{})
	}
}
