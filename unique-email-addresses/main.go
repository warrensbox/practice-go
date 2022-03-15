package main

import "strings"

func numUniqueEmails(emails []string) int {

	mapLocal := make(map[string]bool)
	mapDomain := make(map[string]bool)
	count := 0
	for _, val := range emails {
		arr := strings.Split(val, "@")

		domain := arr[1]
		local := strings.ReplaceAll(arr[0], ".", "")
		if strings.Contains(local, "+") {
			pos := strings.Index(local, "+")
			local = local[:pos]
		}

		if !mapDomain[domain] {
			count++
		} else if !mapLocal[local] {
			count++
		}
		mapLocal[local] = true
		mapDomain[domain] = true
	}

	return count
}

/*

split it by the @ sign

arr[0] = localname
arr[1] = domainname

testemail+alex
@leetcode.com

testemail+bobcathy
@leetcode.com

testemail+david
@lee.tcode.com


 t := strings.Trim(s, "[")

*/
