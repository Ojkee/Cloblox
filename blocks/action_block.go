package blocks

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

// This block is responsible only for actions on variables (not declaring variables)

type ACTION_TYPE int

const (
	MATH_OPERATIONS ACTION_TYPE = iota
	SWAP
	PRINT
	RAND
	UNSIGNED
)

type ActionBlock struct {
	BlockDefault
	next *Block

	actionInput string
	actionType  ACTION_TYPE
	keys        []string
	actionKVP   map[string]float64
}

func NewActionBlock() *ActionBlock {
	return &ActionBlock{
		BlockDefault: BlockDefault{
			id:   -1,
			name: "Action Block",
		},
		next: nil,

		actionInput: "",
		actionType:  UNSIGNED,
		keys:        make([]string, 0),
		actionKVP:   make(map[string]float64),
	}
}

func (block *ActionBlock) GetNext() (*Block, error) {
	return block.next, nil
}

func (block *ActionBlock) SetNext(next Block) {
	block.next = &next
}

func (block *ActionBlock) ParseFromUserInput(input string) error {
	block.Flush()
	block.actionInput = input
	var err error
	block.actionType, err = block.getActionType(&input)
	return err
}

func (block *ActionBlock) GetKeys() []string {
	return block.keys
}

func (block *ActionBlock) SetActionKVP(actionKVP *map[string]float64) {
	block.actionKVP = *actionKVP
}

func (block *ActionBlock) PerformGetUpdateKVP() (updateVars map[string]float64, logMessage string, err error) {
	switch block.actionType {
	case SWAP:
		vals := block.actionSwap()
		return vals, "", nil
	case PRINT:
		message := block.actionPrint()
		return nil, message, nil
	case RAND:
		vals, err := block.actionRand()
		return vals, "", err
	case MATH_OPERATIONS:
		vals := block.actionMathOperations()
		return vals, "", nil
	case UNSIGNED:
		return nil, "", errors.New("Not initiated action block, might be syntax error")
	}
	return nil, "", errors.New("No action Type")
}

func (block *ActionBlock) getActionType(input *string) (ACTION_TYPE, error) {
	lowerInput := strings.ToLower(*input)
	if strings.Contains(lowerInput, "swap") {
		if err := block.parseKeysIfValidSwap(input); err != nil {
			return UNSIGNED, err
		}
		return SWAP, nil
	} else if strings.Contains(lowerInput, "print") {
		return PRINT, nil
	} else if strings.Contains(lowerInput, "rand") {
		if err := block.parseKeysIfValidRand(input); err != nil {
			return UNSIGNED, err
		}
		return RAND, nil
	}
	ops := []string{"=", "+=", "-=", "/=", "*="}
	for _, op := range ops {
		if strings.Contains(*input, op) {
			return MATH_OPERATIONS, nil
		}
	}
	return UNSIGNED, nil
}

// MATH OPERATIONS
func (block *ActionBlock) actionMathOperations() map[string]float64 { // TODO
	retVal := make(map[string]float64)
	return retVal
}

// SWAP
func (block *ActionBlock) parseKeysIfValidSwap(input *string) error {
	r := regexp.MustCompile(
		`[a-zA-Z_][a-zA-Z0-9_]*\[[a-zA-Z0-9_+\-*/\s()]*\]|[a-zA-Z_][a-zA-Z0-9_]*`,
	)
	trimmed := strings.TrimLeft(*input, " ")
	inputNoPrefix, _ := strings.CutPrefix(trimmed, "swap")
	foundKeys := r.FindAllString(inputNoPrefix, -1)
	if len(foundKeys) != 2 {
		return errors.New("Invalid number of variables")
	}
	block.keys = foundKeys
	return nil
}

func (block *ActionBlock) actionSwap() map[string]float64 {
	keys := []string{}
	vals := []float64{}
	retVal := make(map[string]float64)
	for key, val := range block.actionKVP {
		keys = append(keys, key)
		vals = append(vals, val)
	}
	retVal[keys[0]] = vals[1]
	retVal[keys[1]] = vals[0]
	return retVal
}

// PRINT
func (block *ActionBlock) actionPrint() string {
	var retVal bytes.Buffer
	for key, value := range block.actionKVP {
		retVal.WriteString(fmt.Sprintf("%s = %.3f\n", key, value))
	}
	return retVal.String()
}

// RAND
func (block *ActionBlock) parseKeysIfValidRand(input *string) error {
	tokens := strings.Split(*input, "rand")
	if len(tokens) != 2 {
		return errors.New("Invalid input")
	}
	r := regexp.MustCompile(
		`[a-zA-Z_][a-zA-Z0-9_]*\[[a-zA-Z0-9_+\-*/\s()]*\]|[a-zA-Z_][a-zA-Z0-9_]*`,
	)
	numsFound := r.FindAllString(tokens[0], -1)
	if len(numsFound) != 1 {
		return errors.New("Only one variable can be randomized per block")
	}
	numsFoundRange := r.FindAllString(tokens[1], -1)
	block.keys = append(numsFound, numsFoundRange...)
	return nil
}

func (block *ActionBlock) actionRand() (map[string]float64, error) {
	// x = rand 2, 10 only nambers for now
	trimmed := strings.TrimLeft(block.actionInput, " ")
	tokens := strings.Split(trimmed, "rand")
	for key, value := range block.actionKVP {
		numStr := strconv.FormatFloat(value, 'f', -1, 64)
		tokens[1] = strings.ReplaceAll(tokens[1], key, numStr)
	}
	mmin, mmax, err := parseRandomMinMax(&tokens[1])
	if err != nil {
		return nil, err
	}
	rVal := rand.Float64()*(mmax-mmin) + mmin
	for key := range block.actionKVP {
		if strings.Contains(tokens[0], key) {
			block.actionKVP[key] = rVal
			break
		}
	}
	return block.actionKVP, nil
}

func parseRandomMinMax(input *string) (float64, float64, error) {
	// Finds every real floating point number
	r := regexp.MustCompile(`-?\b\d+(\.\d+)?\b`)
	numsFound := r.FindAllString(*input, -1)
	if len(numsFound) != 2 {
		return 0, 0, errors.New("Invalid number of parameters")
	}
	mmin, err := strconv.ParseFloat(numsFound[0], 10)
	if err != nil {
		return 0, 0, errors.New("Can't parse parameter")
	}
	mmax, err := strconv.ParseFloat(numsFound[1], 10)
	if err != nil {
		return 0, 0, errors.New("Can't parse parameter")
	}
	if mmin >= mmax {
		return 0, 0, errors.New("max should be greater that min")
	}
	return mmin, mmax, nil
}

func (block *ActionBlock) Flush() {
	block.actionInput = ""
	block.actionType = UNSIGNED
	block.keys = make([]string, 0)
	block.actionKVP = make(map[string]float64)
}
