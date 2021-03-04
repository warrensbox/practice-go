package main

func main() {

}

type Words struct {
	smalls   []string
	tens     []string
	bigs     []string
	hundred  string
	negative string
}

func initialize() {
	w := Words{}
	w.smalls = []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twele",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}
	w.tens = []string{
		"",
		"",
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}
	w.bigs = []string{
		"",
		"thousands",
		"million",
		"billion",
	}
	w.hundred = "hundred"
	w.negative = "negative"
}

// func (w *Words) convert(num int) string {
// 	if num == 0 {
// 		return w.smalls[0]
// 	} else if num < 0 {
// 		return w.negative + w.convert(-1*num)
// 	}

// 	parts := []string{}

// 	chunkCount := 0

// 	for num > 0 {
// 		if num%1000 != 0 {
// 			chunks = convertChunks()
// 		}
// 	}
// }

// func convertChunks(number int){

// 	if num >= 100 {

// 	}
// }
