package main

func numMatchingSubseq(S string, words []string) (num int) {
	waiting := map[rune][]string{' ': words}
	for _, c := range " " + S {
		//fmt.Println("str",string(c))
		advance := waiting[c]
		delete(waiting, c)
		for _, word := range advance {
			//fmt.Println("word",word)
			if len(word) == 0 {
				num++
				//fmt.Println("Adding")
			} else {
				c := rune(word[0])
				//fmt.Println("char",string(word[0]))
				waiting[c] = append(waiting[c], word[1:])
				//fmt.Println(waiting)
			}
		}
		// fmt.Println("======")
	}
	return
}
