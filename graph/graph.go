package graph

import "Cloblox/blocks"

type Graph struct {
	blocksSlice []*blocks.Block
	head        *blocks.Block
	current     *blocks.Block

	lengthLimit int
}

func NewGraph(blocksSlice []*blocks.Block) *Graph {
	return &Graph{
		blocksSlice: blocksSlice,
		head:        nil,
		current:     nil,
		lengthLimit: 100,
	}
}

func (g *Graph) findStartIdx() (int, bool) {
	for i, block := range g.blocksSlice {
		if isStart(block) {
			return i, true
		}
	}
	return -1, false
}

func isStart(block *blocks.Block) bool {
	switch (*block).(type) {
	case *blocks.StartBlock:
		return true
	}
	return false
}

func isStop(block *blocks.Block) bool {
	switch (*block).(type) {
	case *blocks.StopBlock:
		return true
	}
	return false
}

func isIf(block *blocks.Block) bool {
	switch (*block).(type) {
	case
		*blocks.IfBlock[float32],
		*blocks.IfBlock[int],
		*blocks.IfBlock[string]:
		return true
	}
	return false
}

func DepthFirstSearchStop(head *blocks.Block, visitedIds *[]int) bool {
	return false
}

// Checks if is fully connected,
// doesn't check if it's completly transitive
// E.G. doens't detect if infinite loop exists
func (g *Graph) IsFullyConnected() bool {
	if len(g.blocksSlice) < 2 { // Starts with Start; ends with End
		return false
	}
	if idx, found := g.findStartIdx(); found {
		visitedIds := make([]int, 0)
		return DepthFirstSearchStop(g.blocksSlice[idx], &visitedIds)
	}
	return false
}
