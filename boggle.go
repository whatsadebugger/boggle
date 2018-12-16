package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/whatsadebugger/algorithms-data-structures/DataStructures/trie"
)

func main() {

	letters := [][]string{
		{"r", "h", "r", "e"},
		{"y", "p", "c", "s"},
		{"w", "n", "s", "n"},
		{"t", "e", "g", "o"},
	}
	f, err := os.Open("wordsEn.txt")
	if err != nil {
		panic(err)
	}

	// load up our words into the trie for fast lookups
	s := bufio.NewScanner(f)
	words := trie.NewNode()

	for s.Scan() {
		words.Insert(s.Text())
	}

	boggle(letters, words)
}

type point struct {
	x int
	y int
}

var moves = []point{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}

func boggle(letters [][]string, words *trie.Node) {

	// visited := make(map[point]struct{})

	for i := 0; i < len(letters); i++ {
		for j := 0; j < len(letters[i]); j++ {
			fmt.Println(letters[i][j])
			// visited[newPoint(i, j)] = struct{}{}

		}
	}
}

// Point make points
func newPoint(x, y int) point {
	return point{
		x: x,
		y: y,
	}
}
