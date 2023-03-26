package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	stickersHash := make(map[int]struct{})
	stickers := make([]int, 0)

	line, _ = reader.ReadString('\n')
	for _, sticker := range strings.Fields(line) {
		number, _ := strconv.Atoi(sticker)

		if _, ok := stickersHash[number]; !ok {
			stickersHash[number] = struct{}{}
			stickers = append(stickers, number)
		}
	}

	sort.Ints(stickers)

	line, _ = reader.ReadString('\n')
	line, _ = reader.ReadString('\n')

	writer := bufio.NewWriter(os.Stdout)

	for _, sticker := range strings.Fields(line) {
		num, _ := strconv.Atoi(sticker)

		if num > stickers[len(stickers)-1] {
			_, _ = fmt.Fprintln(writer, len(stickers))
		} else if num <= stickers[0] {
			_, _ = fmt.Fprintln(writer, 0)
		} else {
			result := sort.SearchInts(stickers, num)

			if result > 0 && stickers[result] >= num {
				result--
			}

			_, _ = fmt.Fprintln(writer, result+1)
		}
	}

	_ = writer.Flush()
}
