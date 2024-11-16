package blocks

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/Pramod-Devireddy/go-exprtk"

	"Cloblox/functools"
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

	replaceKey          string             // key to replace every array array key for parsing
	keys                []string           // every array with index and non-array var
	arrayKeys           map[string]string  // key: array with index val: symbole
	actionKVP           map[string]float64 // key: array with index and non-array var val: value
	actionInputRaw      string             // non edited expr
	actionInputReplaced string             // replaced with replaceKey and set rounding
	actionType          ACTION_TYPE
}

func NewActionBlock() *ActionBlock {
	return &ActionBlock{
		BlockDefault: BlockDefault{
			id:   -1,
			name: "Action Block",
		},
		next: nil,

		replaceKey:          "TEMPKEY",
		keys:                make([]string, 0),
		arrayKeys:           make(map[string]string),
		actionKVP:           make(map[string]float64),
		actionInputRaw:      "",
		actionInputReplaced: "",
		actionType:          UNSIGNED,
	}
}

func (block *ActionBlock) GetNext() (*Block, error) {
	return block.next, nil
}

func (block *ActionBlock) SetNext(next *Block) {
	block.next = next
}

func (block *ActionBlock) ParseFromUserInput(input string) error {
	block.Flush()
	block.actionInputRaw = input
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
		vals, err := block.actionMathOperations()
		return vals, "", err
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
		if err := block.parseKeysIfValidPrint(input); err != nil {
			return UNSIGNED, err
		}
		return PRINT, nil
	} else if strings.Contains(lowerInput, "rand") {
		if err := block.parseKeysIfValidRand(input); err != nil {
			return UNSIGNED, err
		}
		return RAND, nil
	}
	ops := []string{"=", "+=", "-=", "/=", "*=", "++", "--"}
	for _, op := range ops {
		if strings.Contains(*input, op) {
			if err := block.parseKeysIfValidMathOperations(input); err != nil {
				return UNSIGNED, err
			}
			return MATH_OPERATIONS, nil
		}
	}
	return UNSIGNED, errors.New("Invalid input in action block")
}

// MATH OPERATIONS
func (block *ActionBlock) parseKeysIfValidMathOperations(input *string) error {
	keysFound := getKeysFromString(input)
	if len(keysFound) == 0 {
		return errors.New("No vars")
	}
	ops := []string{"+=", "-=", "/=", "*=", "="}
	for _, op := range ops {
		if strings.Contains(*input, op) {
			tokens := strings.Split(*input, op)
			if len(tokens) < 2 {
				return errors.New("Invalid syntax")
			} else if strings.TrimSpace(tokens[1]) == "" {
				return errors.New("Invalid syntax")
			}
		}
	}
	block.keys = keysFound
	block.actionInputReplaced, block.arrayKeys = findReplaceArrayKeys(*input, block.replaceKey)
	return nil
}

// POSSIBLE INPUT:
// x++;  x += 2; x /= 3;
// x = x / 0; x = 2 * 4; x += i*g;
// s[i] = i * 2; s[d*c+2] /= d*s[i]*k + 9
// etc.
func (block *ActionBlock) actionMathOperations() (map[string]float64, error) {
	retVal := make(map[string]float64)
	ops := []string{"+=", "-=", "/=", "*=", "++", "--", "="}
	var operator string
	for _, op := range ops {
		if strings.Contains(block.actionInputReplaced, op) {
			operator = op
			break
		}
	}
	tokens := strings.Split(block.actionInputRaw, operator)
	lhsExpr := getKeysFromString(&tokens[0])
	if len(lhsExpr) != 1 {
		return nil, errors.New("Invalid number of lhs")
	}
	lhsKey := lhsExpr[0]
	if operator == "++" {
		retVal[lhsKey] = block.actionKVP[lhsKey] + 1
		return retVal, nil
	} else if operator == "--" {
		retVal[lhsKey] = block.actionKVP[lhsKey] - 1
		return retVal, nil
	}
	tokensReplaced := strings.Split(block.actionInputReplaced, operator)
	evaluated, err := block.evaluateRHS(&tokensReplaced[1])
	if err != nil {
		return nil, err
	}
	return block.validateEvalGetUpdateKVP(&evaluated, &operator, &lhsKey)
}

func (block *ActionBlock) evaluateRHS(replacedRHS *string) (float64, error) {
	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()

	exprtkObj.SetExpression(*replacedRHS + " + 0")
	for _, key := range block.keys {
		if strings.Contains(*replacedRHS, key) {
			exprtkObj.AddDoubleVariable(key)
		} else if replaceKey, ok := block.arrayKeys[key]; ok {
			if strings.Contains(*replacedRHS, replaceKey) {
				exprtkObj.AddDoubleVariable(replaceKey)
			}
		}
	}
	err := exprtkObj.CompileExpression()
	if err != nil {
		return 0, err
	}
	for key, val := range block.actionKVP {
		if strings.Contains(*replacedRHS, key) {
			exprtkObj.SetDoubleVariableValue(key, val)
		} else if replaceKey, ok := block.arrayKeys[key]; ok {
			if strings.Contains(*replacedRHS, replaceKey) {
				exprtkObj.SetDoubleVariableValue(replaceKey, val)
			}
		}
	}
	return exprtkObj.GetEvaluatedValue(), nil
}

func (block *ActionBlock) validateEvalGetUpdateKVP(
	evaluated *float64,
	operator, lhsKey *string,
) (map[string]float64, error) {
	if *evaluated == math.Inf(1) || *evaluated == math.Inf(-1) {
		return nil, errors.New("Illigal math operation")
	}
	if *evaluated == 0 && (*operator == "/" || *operator == "/=") {
		return nil, errors.New("Division by zero error")
	}
	retVal := make(map[string]float64)
	switch *operator {
	case "=":
		retVal[*lhsKey] = *evaluated
		break
	case "+=":
		retVal[*lhsKey] = block.actionKVP[*lhsKey] + *evaluated
		break
	case "-=":
		retVal[*lhsKey] = block.actionKVP[*lhsKey] - *evaluated
		break
	case "*=":
		retVal[*lhsKey] = block.actionKVP[*lhsKey] * *evaluated
		break
	case "/=":
		retVal[*lhsKey] = block.actionKVP[*lhsKey] / *evaluated
		break
	}
	return retVal, nil
}

// SWAP
func (block *ActionBlock) parseKeysIfValidSwap(input *string) error {
	inputNoPrefix := (*input)[strings.LastIndex(*input, "swap")+len("swap"):]
	foundKeys := getKeysFromString(&inputNoPrefix)
	if len(foundKeys) != 2 {
		consoleMess := fmt.Sprintf("Invalid number of variables in line: %s", *input)
		debugMess := fmt.Sprintf("action_block fail:\n\t%s", consoleMess)
		return functools.NewStrongError(consoleMess, debugMess)
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
func (block *ActionBlock) parseKeysIfValidPrint(input *string) error {
	prefixIdx := strings.LastIndex(*input, "print")
	inputNoPrefix := (*input)[prefixIdx+len("print"):]
	keysFound := getKeysFromString(&inputNoPrefix)
	if len(keysFound) == 0 {
		return errors.New("Nothing to print")
	}
	block.keys = keysFound
	return nil
}

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
	numsFound := getKeysFromString(&tokens[0])
	if len(numsFound) != 1 {
		return errors.New("Only one variable can be randomized per block")
	}
	numsFoundRange := getKeysFromString(&tokens[1])
	block.keys = append(numsFound, numsFoundRange...)
	return nil
}

func (block *ActionBlock) actionRand() (map[string]float64, error) {
	trimmed := strings.TrimLeft(block.actionInputRaw, " ")
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

// Finds every real floating point number
func parseRandomMinMax(input *string) (float64, float64, error) {
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
	block.replaceKey = "TEMPKEY"
	block.keys = make([]string, 0)
	block.arrayKeys = make(map[string]string)
	block.actionKVP = make(map[string]float64)
	block.actionInputRaw = ""
	block.actionInputReplaced = ""
	block.actionType = UNSIGNED
}
