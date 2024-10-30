package blocks

import (
	"fmt"
	"regexp"
	"strings"
)

type BlockType interface {
	float32 | []float32
}

type Block interface {
	GetId() int
	SetId(id int)
	GetName() string
	GetNext() (*Block, error)
}

type BlockSingleOut interface {
	Block
	SetNext(next Block)
}

type BlockManyOut interface {
	Block
	GetNextTrue() *Block
	GetNextFalse() *Block
	SetNextTrue(next Block)
	SetNextFalse(next Block)
}

type BlockWithVars interface {
	GetVars() map[string]any
}

type BlockDefault struct {
	id   int
	name string
}

func (b *BlockDefault) GetId() int {
	return b.id
}

func (b *BlockDefault) SetId(id int) {
	b.id = id
}

func (b *BlockDefault) GetName() string {
	return b.name
}

// Searching arrays that may has non-leading numbers
// such as my_tab2[(i-1)*2], but they aren't non-array variables
// such as x, i2, my_x etc.
func getKeysFromString(input *string) []string {
	r := regexp.MustCompile(
		`[a-zA-Z_][a-zA-Z0-9_]*\[[a-zA-Z0-9_+\-*/\s()]*\]|[a-zA-Z_][a-zA-Z0-9_]*`,
	)
	keysFound := r.FindAllString(*input, -1)
	return keysFound
}

// a[4]         // In    arr[i]       // In    arr[x + i] // In
// tab[dd]      // In    [dkf]        // Out   arr[x + 2] // In
// [4]          // Out   myArray[45]  // In
// dkjf[4k]     // Out   my_array[df] // In
// t[my_x]      // In    dlsk[df sf]  // Out
func findReplaceArrayKeys(input, replaceKey string) (string, map[string]string) {
	r := regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*\[(?:[a-zA-Z0-9_+\-*/\s()]+)\]`)
	retVal := make(map[string]string)
	arrayKeysFound := r.FindAllString(input, -1)
	for i, key := range arrayKeysFound {
		nextReplaceKey := fmt.Sprintf("%s%d", replaceKey, i)
		input = strings.ReplaceAll(input, key, nextReplaceKey)
		retVal[key] = nextReplaceKey
	}
	return input, retVal
}
