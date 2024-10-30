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

const (
	INVALID_PATH_ERR_MESSAGE  = "At least one path doesn't end in 'Stop' block"
	INVALID_START_ERR_MESSAGE = "There must be exacly one 'Start' block"
)

type Graph struct {
	blockCounter int
	blocksSlice  []blocks.Block
	head         *blocks.Block
	current      *blocks.Block

	stepCounter int
	stepLimit   int

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
		blockCounter: len(*blocksSlice),
		blocksSlice:  *blocksSlice,
		head:         nil,
		current:      nil,

		stepCounter: 0,
		stepLimit:   1024,

		allCurrentVars: make(map[string]any),
		isFinished:     false,
	}
}

// Assigns head node if there is only one Start block
// and all paths lead to Stop block/blocks
func (graph *Graph) InitIfValid() error {
	if !isOneStart(&graph.blocksSlice) {
		return errors.New(INVALID_START_ERR_MESSAGE)
	}

	startNode := &graph.blocksSlice[graph.findStartIdx()]
	graph.head = startNode
	graph.current = startNode
	visitedIds := make([]int, 0)

	allInStop, anyInNil := dfsAllPathsStops(graph.head, &visitedIds)
	if anyInNil || !allInStop {
		return errors.New(INVALID_PATH_ERR_MESSAGE)
	}
	return nil
}

func (graph *Graph) GetHead() *blocks.Block {
	return graph.head
}

func (graph *Graph) GetCurrent() *blocks.Block {
	return graph.current
}

func (graph *Graph) GetAllBlocks() []blocks.Block {
	return graph.blocksSlice
}

func (graph *Graph) AppendBlock(block blocks.Block) {
	block.SetId(graph.blockCounter)
	graph.blockCounter += 1
	graph.blocksSlice = append(graph.blocksSlice, block)
}

func (graph *Graph) ConnectByIds(idFrom, idTo int, isNextTrue ...bool) error {
	if idFrom == idTo {
		return errors.New("graph/ConnectByIDs fail:\n\tCan't connect to itself")
	}
	sliceIdFrom, sliceIdTo := graph.findIdsInSlice(idFrom, idTo)
	if sliceIdFrom == -1 || sliceIdTo == -1 {
		return errors.New(fmt.Sprintf(
			"graph/ConnectByIDs fail:\n\tId doesn't exist\n\tFrom: %d To: %d",
			sliceIdFrom, sliceIdTo,
		))
	}
	if err := graph.connectBlocks(sliceIdFrom, sliceIdTo, isNextTrue...); err != nil {
		return err
	}
	return nil
}

func (graph *Graph) IsConnectedByIds(idFrom, idTo int) bool {
	for _, block := range graph.blocksSlice {
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

func (graph *Graph) GetAllVars() map[string]any {
	return graph.allCurrentVars
}

func (graph *Graph) Log() { // Debug
	fmt.Println("******************************************")
	for _, block := range graph.blocksSlice {
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

func (graph *Graph) SetAllVars(allVars map[string]any) {
	graph.allCurrentVars = allVars
}

func (graph *Graph) GetKvpByKeys(keys *[]string) (map[string]float32, error) {
	// idxer - expression within square brackets
	retVal := make(map[string]float32)
	for _, key := range *keys {
		if arrayKey, idxer, found := getIfArrayKey(&key); found {
			if !graph.isKeyInVars(arrayKey) {
				return nil, errors.New("Invalid array variable")
			}
			idxParsed, err := graph.parseArrayIdxer(idxer)
			if err != nil {
				return nil, err
			}
			valSlice := graph.allCurrentVars[arrayKey]
			val, err := graph.getValFromSliceIfValid(valSlice, idxParsed)
			if err != nil {
				return nil, err
			}
			retVal[key] = val
		} else if graph.isKeyInVars(key) {
			val, err := graph.getValIfValid(key)
			if err != nil {
				return nil, err
			}
			retVal[key] = val
		} else {
			return nil, errors.New("Invalid array variable")
		}
	}
	return retVal, nil
}

func (graph *Graph) MakeStep() error {
	if graph.stepCounter >= graph.stepLimit {
		return errors.New("Limit of steps exceeded")
	}

	return nil
}

func dfsAllPathsStops(node *blocks.Block, visitedIds *[]int) (inStop, inNil bool) {
	if node == nil {
		return false, true
	}
	if slices.Contains(*visitedIds, (*node).GetId()) {
		return false, false
	}
	if isStop(node) {
		return true, false
	}
	*visitedIds = append(*visitedIds, (*node).GetId())

	if manyOutBlock, ok := (*node).(blocks.BlockManyOut); ok {
		trueBlock := manyOutBlock.GetNextTrue()
		falseBlock := manyOutBlock.GetNextFalse()
		inStopTrue, inNilTrue := dfsAllPathsStops(trueBlock, visitedIds)
		inStopFalse, inNilFalse := dfsAllPathsStops(falseBlock, visitedIds)
		if inNilTrue || inNilFalse {
			return false, true
		}
		return inStopTrue || inStopFalse, false
	}
	next, err := (*node).GetNext()
	if err != nil {
		panic(err)
	}
	return dfsAllPathsStops(next, visitedIds)
}

func isOneStart(bSlice *[]blocks.Block) bool {
	startCounter := 0
	for _, block := range *bSlice {
		if isStart(&block) {
			startCounter += 1
		}
	}
	if startCounter != 1 {
		return false
	}
	return true
}

func (graph *Graph) findIdsInSlice(idFrom, idTo int) (int, int) {
	sliceIdFrom := -1
	sliceIdTo := -1
	for i, block := range graph.blocksSlice {
		if idFrom == block.GetId() {
			sliceIdFrom = i
		} else if idTo == block.GetId() {
			sliceIdTo = i
		}
	}
	return sliceIdFrom, sliceIdTo
}

func (graph *Graph) connectBlocks(src, dst int, isNextTrue ...bool) error {
	if manyBlock, ok := graph.blocksSlice[src].(blocks.BlockManyOut); ok {
		if len(isNextTrue) == 0 {
			return errors.New(
				"graph/ConnectByIDs fail:\n\tTrying connect to ManyOut without path specified",
			)
		}
		if isNextTrue[0] {
			manyBlock.SetNextTrue(graph.blocksSlice[dst])
		} else {
			manyBlock.SetNextFalse(graph.blocksSlice[dst])
		}
	} else if singleBlock, ok := graph.blocksSlice[src].(blocks.BlockSingleOut); ok {
		singleBlock.SetNext(graph.blocksSlice[dst])
	}
	return nil
}

func (graph *Graph) findStartIdx() int {
	for i, block := range graph.blocksSlice {
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

func (graph *Graph) parseArrayIdxer(idxer string) (int, error) {
	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()
	exprtkObj.SetExpression(idxer)
	r := regexp.MustCompile(`\b[a-zA-Z_][a-zA-Z0-9_]*\b`)
	keysFound := r.FindAllString(idxer, -1)

	foundKVP := make(map[string]float32)
	for _, key := range keysFound {
		val, err := graph.getValIfValid(key)
		if err != nil {
			return 0, err
		}
		exprtkObj.AddDoubleVariable(key)
		foundKVP[key] = val
	}
	err := exprtkObj.CompileExpression()
	if err != nil {
		return 0, err
	}
	for key, val := range foundKVP {
		exprtkObj.SetDoubleVariableValue(key, float64(val))
	}
	retVal := exprtkObj.GetEvaluatedValue()
	return int(retVal), nil
}

func (graph *Graph) isKeyInVars(key string) bool {
	if _, ok := graph.allCurrentVars[key]; ok {
		return true
	}
	return false
}

func (graph *Graph) getValFromSliceIfValid(valSlice any, idxer int) (float32, error) {
	idxErr := errors.New("Invalid index of the array")
	switch s := valSlice.(type) {
	case []float64:
		if idxer >= 0 && idxer <= len(s) {
			return float32(s[idxer]), nil
		}
		return 0, idxErr
	case []float32:
		if idxer >= 0 && idxer <= len(s) {
			return s[idxer], nil
		}
		return 0, idxErr
	case []int:
		if idxer >= 0 && idxer <= len(s) {
			return float32(s[idxer]), nil
		}
		return 0, idxErr
	}
	return 0, errors.New("Variable isn't an array of numbers")
}

func (graph *Graph) getValIfValid(key string) (float32, error) {
	switch v := graph.allCurrentVars[key].(type) {
	case float64:
		return float32(v), nil
	case float32:
		return float32(v), nil
	case int:
		return float32(v), nil
	}
	return 0, errors.New("Value isn't number")
}

func (graph *Graph) UpdateVarsFromKVP(newVars *map[string]float64) error {
	for key, val := range *newVars {
		if arrayKey, idxer, found := getIfArrayKey(&key); found {
			if !graph.isKeyInVars(arrayKey) {
				return errors.New("Invalid array variable")
			}
			idxParsed, err := graph.parseArrayIdxer(idxer)
			if err != nil {
				return err
			}
			err = graph.updateValInSlice(&arrayKey, &idxParsed, &val)
			if err != nil {
				return err
			}
		} else if graph.isSliceByKey(&key) {
			return errors.New("Can't assign value to array variable")
		} else {
			graph.allCurrentVars[key] = val
		}
	}
	return nil
}

func (graph *Graph) updateValInSlice(arrayKey *string, idx *int, val *float64) error {
	switch valSlice := graph.allCurrentVars[*arrayKey].(type) {
	case []float64:
		if *idx < 0 || *idx > len(valSlice) {
			return errors.New("Idx out of range")
		}
		valSlice[*idx] = *val
		graph.allCurrentVars[*arrayKey] = valSlice
	}
	return nil
}

func (graph *Graph) isSliceByKey(key *string) bool {
	switch graph.allCurrentVars[*key].(type) {
	case []float64:
		return true
	}
	return false
}
