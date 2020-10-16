package main

func main() {

}

func URLify(str []string) {

	spaces := 0
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			spaces++
		}
	}

	newlength := len(str) + spaces*2
	str[newlength] = "/0"

	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == " " {
			str[newlength-1] = "0"
			str[newlength-2] = "0"
			str[newlength-3] = "0"

			newlength = newlength - 3
		} else {
			str[newlength-1] = str[i]
			newlength = newlength - 1
		}
	}

}
