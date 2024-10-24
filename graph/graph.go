package graph

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/Pramod-Devireddy/go-exprtk"

	"Cloblox/blocks"
)

type Graph struct {
	blocksSlice []blocks.Block
	head        *blocks.Block
	current     *blocks.Block

	blockCounter int
	lengthLimit  int

	allCurrentVars map[string]any
	isFinished     bool
}

func NewGraph(blocksSlice *[]blocks.Block) *Graph {
	if blocksSlice == nil {
		blocksS := make([]blocks.Block, 0)
		blocksSlice = &blocksS
	} else {
		for i, block := range *blocksSlice {
			block.SetId(i)
		}
	}
	return &Graph{
		blocksSlice:    *blocksSlice,
		head:           nil,
		current:        nil,
		blockCounter:   len(*blocksSlice),
		lengthLimit:    100,
		allCurrentVars: make(map[string]any),
		isFinished:     false,
	}
}

// Checks if is fully connected
// doesn't check if it's completly transitive
// E.G. doens't detect if infinite loop exists
func (g *Graph) IsFullyConnected() bool {
	if len(g.blocksSlice) < 2 {
		return false
	}
	if idx := g.findStartIdx(); idx != -1 {
		visitedIds := make([]int, 0)
		return dfsStop(&(g.blocksSlice)[idx], &visitedIds)
	}
	return false
}

func (g *Graph) GetAllBlocks() []blocks.Block {
	return g.blocksSlice
}

func (g *Graph) AppendBlock(block blocks.Block) {
	block.SetId(g.blockCounter)
	g.blockCounter += 1
	g.blocksSlice = append(g.blocksSlice, block)
}

func (g *Graph) ConnectByIds(idFrom, idTo int, isNextTrue ...bool) error {
	if idFrom == idTo {
		return errors.New("graph/ConnectByIDs fail:\n\tCan't connect to itself")
	}
	sliceIdFrom, sliceIdTo := g.findIdsInSlice(idFrom, idTo)
	if sliceIdFrom == -1 || sliceIdTo == -1 {
		return errors.New(fmt.Sprintf(
			"graph/ConnectByIDs fail:\n\tId doesn't exist\n\tFrom: %d To: %d",
			sliceIdFrom, sliceIdTo,
		))
	}
	if err := g.connectBlocks(sliceIdFrom, sliceIdTo, isNextTrue...); err != nil {
		return err
	}
	return nil
}

func (g *Graph) IsConnectedByIds(idFrom, idTo int) bool {
	for _, block := range g.blocksSlice {
		if singleBlock, ok := block.(blocks.BlockSingleOut); ok {
			next, _ := singleBlock.GetNext()
			if next == nil {
				continue
			}
			if (*next).GetId() == idTo {
				return true
			}
		} else if manyBlock, ok := block.(blocks.BlockManyOut); ok {
			nextTrue := *manyBlock.GetNextTrue()
			if nextTrue == nil {
				continue
			}
			if nextTrue.GetId() == idTo {
				return true
			}
			nextFalse := *manyBlock.GetNextFalse()
			if nextFalse == nil {
				continue
			}
			if nextFalse.GetId() == idTo {
				return true
			}
		}
	}
	return false
}

func (g *Graph) GetAllVars() map[string]any {
	return g.allCurrentVars
}

func (g *Graph) findIdsInSlice(idFrom, idTo int) (int, int) {
	sliceIdFrom := -1
	sliceIdTo := -1
	for i, block := range g.blocksSlice {
		if idFrom == block.GetId() {
			sliceIdFrom = i
		} else if idTo == block.GetId() {
			sliceIdTo = i
		}
	}
	return sliceIdFrom, sliceIdTo
}

func (g *Graph) connectBlocks(src, dst int, isNextTrue ...bool) error {
	if manyBlock, ok := g.blocksSlice[src].(blocks.BlockManyOut); ok {
		if len(isNextTrue) == 0 {
			return errors.New(
				"graph/ConnectByIDs fail:\n\tTrying connect to ManyOut without path specified",
			)
		}
		if isNextTrue[0] {
			manyBlock.SetNextTrue(g.blocksSlice[dst])
		} else {
			manyBlock.SetNextFalse(g.blocksSlice[dst])
		}
	} else if singleBlock, ok := g.blocksSlice[src].(blocks.BlockSingleOut); ok {
		singleBlock.SetNext(g.blocksSlice[dst])
	}
	return nil
}

func (g *Graph) findStartIdx() int {
	for i, block := range g.blocksSlice {
		if isStart(&block) {
			return i
		}
	}
	return -1
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
	case *blocks.IfBlock:
		return true
	}
	return false
}

func isVariable(block *blocks.Block) bool {
	switch (*block).(type) {
	case *blocks.IfBlock:
		return true
	}
	return false
}

func dfsStop(node *blocks.Block, visitedIds *[]int) bool {
	if node == nil {
		return false
	}
	if slices.Contains(*visitedIds, (*node).GetId()) {
		return false
	}
	if isStop(node) {
		return true
	}
	visited := append(*visitedIds, (*node).GetId())
	visitedIds = &visited
	if manyOutBlock, ok := (*node).(blocks.BlockManyOut); ok {
		trueBlock := manyOutBlock.GetNextTrue()
		falseBlock := manyOutBlock.GetNextFalse()
		return dfsStop(trueBlock, visitedIds) ||
			dfsStop(falseBlock, visitedIds)
	}
	next, _ := (*node).GetNext()
	return dfsStop(next, visitedIds)
}

func (g *Graph) Log() { // Debug
	fmt.Println("******************************************")
	for _, block := range g.blocksSlice {
		fmt.Printf("%d ", block.GetId())
		if mBlock, ok := block.(blocks.BlockManyOut); ok {
			nextString := "  ->  "
			nextFalse := mBlock.GetNextFalse()
			if nextFalse != nil {
				nextString += fmt.Sprintf("%d, ", (*nextFalse).GetId())
			} else {
				nextString += "nil, "
			}
			nextTrue := mBlock.GetNextTrue()
			if nextTrue != nil {
				nextString += fmt.Sprintf("%d, ", (*nextTrue).GetId())
			} else {
				nextString += "nil "
			}
			fmt.Print(nextString)
		} else {
			next, _ := block.GetNext()
			if next == nil {
				fmt.Print("  ->  nil")
			} else {
				fmt.Printf("  ->  %d", (*next).GetId())
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func (g *Graph) SetAllVars(allVars map[string]any) {
	g.allCurrentVars = allVars
}

func (g *Graph) GetKvpByKeys(keys *[]string) (map[string]float32, error) {
	retVal := make(map[string]float32)
	for _, key := range *keys {
		if arrayKey, idxer, found := getIfArrayKey(&key); found {
			if g.isKeyInVars(arrayKey) {
				idxParsed, err := g.parseArrayIdxer(idxer)
				if err != nil {
					return nil, err
				}
				valSlice := g.allCurrentVars[arrayKey]
				val, err := g.getValFromSliceIfValid(valSlice, idxParsed)
				if err != nil {
					return nil, err
				}
				retVal[key] = val
			} else {
				return nil, errors.New("Invalid array variable")
			}
		} else {
			if g.isKeyInVars(key) {
				if val, err := g.getValIfValid(key); err != nil {
					return nil, err
				} else {
					retVal[key] = val
				}
			} else {
				return nil, errors.New("Invalid array variable 3")
			}
		}
	}
	return retVal, nil
}

func getIfArrayKey(key *string) (string, string, bool) {
	r := regexp.MustCompile(`\[(.+?)\]`)
	keyFound := r.Find([]byte(*key))
	if len(keyFound) == 0 {
		return "", "", false
	}
	arrayKey := strings.TrimRight(*key, string(keyFound))
	idxer := string(keyFound[1 : len(keyFound)-1])
	return arrayKey, idxer, true
}

func (g *Graph) parseArrayIdxer(idxer string) (int, error) {
	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()
	exprtkObj.SetExpression(idxer)
	r := regexp.MustCompile(`\b[a-zA-Z_][a-zA-Z0-9_]*\b`)
	keysFound := r.FindAllString(idxer, -1)

	temp := make(map[string]float32)
	for _, key := range keysFound {
		val, err := g.getValIfValid(key)
		if err != nil {
			return 0, err
		}
		exprtkObj.AddDoubleVariable(key)
		temp[key] = val
	}
	err := exprtkObj.CompileExpression()
	if err != nil {
		return 0, err
	}
	for key, val := range temp {
		exprtkObj.SetDoubleVariableValue(key, float64(val))
	}
	retVal := exprtkObj.GetEvaluatedValue()
	return int(retVal), nil
}

func (g *Graph) isKeyInVars(key string) bool {
	if _, ok := g.allCurrentVars[key]; ok {
		return true
	}
	return false
}

func (g *Graph) getValFromSliceIfValid(valSlice any, idxer int) (float32, error) {
	switch s := valSlice.(type) {
	case []float32:
		if idxer >= 0 && idxer <= len(s) {
			return s[idxer], nil
		}
		return 0, errors.New("Invalid index of the array")
	}
	return 0, errors.New("Variable isn't an array of numbers")
}

func (g *Graph) getValIfValid(key string) (float32, error) {
	switch v := g.allCurrentVars[key].(type) {
	case float64:
		return float32(v), nil
	case float32:
		return float32(v), nil
	case int:
		return float32(v), nil
	}
	fmt.Println(key)
	return 0, errors.New("Value isn't number")
}
