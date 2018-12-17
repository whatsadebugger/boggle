package main

import (
	"bufio"
	"os"

	"github.com/whatsadebugger/algorithms-data-structures/DataStructures/graph"
	"github.com/whatsadebugger/algorithms-data-structures/DataStructures/trie"
)

func main() {
	letters := [][]string{
		{"r", "h", "r", "e"},
		{"y", "p", "c", "s"},
		{"w", "n", "s", "n"},
		{"t", "e", "g", "o"},
	}
	g := ArrayToGraph(letters)
	f, err := os.Open("wordsEn.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

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

func ArrayToGraph(slice [][]string) graph.Graph {
	g := graph.Graph{}
	for i, row := range slice {
		for j, value := range row {
			if g.FindVertexByLabel(value) == nil {
				g.AddVertex(value)
			}

			// find all the neighbors
			neighbors := findNeighbors(i, j, len(slice), len(slice[0]))
			vertex := g.FindVertexByLabel(value)

		}
	}
	return g
}

func findNeighbors(x, y, xLimit, yLimit int) []*graph.Vertex {
	return nil
}

func boggle(letters [][]string, words *trie.Node) {

	// visited := make(map[point]struct{})

	for i := 0; i < len(letters); i++ {
		for j := 0; j < len(letters[i]); j++ {
			// for each letter recursively find words
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
