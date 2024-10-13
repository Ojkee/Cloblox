package graph

import (
	"slices"

	"Cloblox/blocks"
)

type Graph struct {
	blocksSlice *[]blocks.Block
	head        *blocks.Block
	current     *blocks.Block

	lengthLimit int
}

func NewGraph(blocksSlice *[]blocks.Block) *Graph {
	return &Graph{
		blocksSlice: blocksSlice,
		head:        nil,
		current:     nil,
		lengthLimit: 100,
	}
}

func (g *Graph) findStartIdx() (int, bool) {
	for i, block := range *g.blocksSlice {
		if isStart(&block) {
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

func depthFirstSearchStop(node *blocks.Block, visitedIds *[]int) bool {
	if node == nil {
		return false
	}
	if slices.Contains(*visitedIds, (*node).GetId()) {
		return false
	}

	visited := append(*visitedIds, (*node).GetId())
	visitedIds = &visited
	if isStop(node) {
		return true
	}
	if manyOutBlock, ok := (*node).(blocks.ManyOutBlock); ok {
		trueBlock := manyOutBlock.GetNextTrue()
		falseBlock := manyOutBlock.GetNextFalse()
		return depthFirstSearchStop(trueBlock, visitedIds) ||
			depthFirstSearchStop(falseBlock, visitedIds)
	}
	next := (*node).GetNext()
	return depthFirstSearchStop(next, visitedIds)
}

// Checks if is fully connected
// doesn't check if it's completly transitive
// E.G. doens't detect if infinite loop exists
func (g *Graph) IsFullyConnected() bool {
	if len(*g.blocksSlice) < 2 {
		return false
	}
	if idx, found := g.findStartIdx(); found {
		visitedIds := make([]int, 0)
		return depthFirstSearchStop(&(*g.blocksSlice)[idx], &visitedIds)
	}
	return false
}

func (g *Graph) GetAllBlocks() *[]blocks.Block {
	return g.blocksSlice
}
