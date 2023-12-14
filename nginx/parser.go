package nginx

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ParseNginxConfig(reader io.Reader) (stack []Items) {
	scanner := bufio.NewScanner(reader)
	var currentBlock *Block
	var stackPointer []Items

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		if strings.HasSuffix(line, "{") {
			blockName := strings.TrimSuffix(line, "{")
			newBlock := &Block{Name: blockName}
			if currentBlock != nil {
				currentBlock.Items = append(currentBlock.Items, newBlock)
			} else {
				stack = append(stack, newBlock)
			}
			stackPointer = append(stackPointer, newBlock)
			currentBlock = newBlock
		} else if strings.HasPrefix(line, "}") {
			if len(stackPointer) > 1 {
				stackPointer = stackPointer[:len(stackPointer)-1]
				currentBlock = stackPointer[len(stackPointer)-1].(*Block)
			} else {
				currentBlock = nil
			}
		} else {
			parts := strings.SplitN(line, " ", 2)
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				directive := &Directive{Name: name, Value: value}
				if currentBlock != nil {
					currentBlock.Items = append(currentBlock.Items, directive)
				} else {
					stack = append(stack, directive)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return stack
}
