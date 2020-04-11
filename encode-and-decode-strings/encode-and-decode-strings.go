package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	var codec Codec
	input := []string{"Hola World"}
	output := codec.Decode(codec.Encode(input))

	fmt.Println(output)

}

const DELIM = '|'

type Codec struct {
	empty bool
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {

	out := []byte{}

	if len(strs) == 0 {
		codec.empty = true
		return string(out)
	}

	fmt.Println(len(strs))
	for i := 0; i < len(strs); i++ {
		if i > 0 {
			out = append(out, DELIM)
		}

		data := base64.StdEncoding.EncodeToString([]byte(strs[i]))
		out = append(out, data...)
	}

	fmt.Println(string(out))
	return string(out)
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {

	out := []string{}

	if codec.empty == true {
		return out
	}

	if len(strs) == 0 {
		return []string{""}
	}

	buffer := []byte{}
	for i := 0; i < len(strs); i++ {

		c := strs[i]

		if c == DELIM {
			out = append(out, dumpBuffer(buffer))
			buffer = []byte{}
		} else {
			buffer = append(buffer, c)
		}
	}

	out = append(out, dumpBuffer(buffer))

	return out
}

func dumpBuffer(buffer []byte) string {

	if len(buffer) == 0 {
		return ""
	}

	data, _ := base64.StdEncoding.DecodeString(string(buffer))

	return string(data)
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
