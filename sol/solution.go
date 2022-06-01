package sol

import "fmt"

func Run(commands []string, values [][]int) []string {
	result := []string{"null"}
	cLen := len(commands)
	medianFinder := Constructor()
	for idx := 1; idx < cLen; idx++ {
		command := commands[idx]
		switch command {
		case "addNum":
			medianFinder.AddNum(values[idx][0])
			result = append(result, "null")
		case "findMedian":
			result = append(result, fmt.Sprintf("%v", medianFinder.FindMedian()))
		}
	}
	return result
}
