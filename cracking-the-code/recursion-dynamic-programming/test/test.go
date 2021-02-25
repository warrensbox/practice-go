package main

import "fmt"

func main() {

	//fmt.Println(sumForLoop(4))
	//fmt.Println(sumRecursion(4))

	array1 := []string{"/aws/lambda/esp-ENVS-lambda-asyncwebservicelambda", "/aws/lambda/esp-ENVS-lambda-email-generation", "/aws/lambda/esp-ENVS-lambda-email-notification", "/aws/lambda/esp-ENVS-lambda-importer-checkevidencestatus", "/aws/lambda/esp-ENVS-lambda-importer-checkimportrequeststatus", "/aws/lambda/esp-ENVS-lambda-importer-handlecancellationerror", "/aws/lambda/esp-ENVS-lambda-importer-monitorincomplete", "/aws/lambda/esp-ENVS-lambda-importer-postevidenceupload", "/aws/lambda/esp-ENVS-lambda-pre-deployment-check", "/aws/lambda/esp-ENVS-lambda-scoring-checkscoresysrequeststatus", "/aws/lambda/esp-ENVS-lambda-oc-monitorincomplete", "/aws/lambda/esp-ENVS-lambda-portfolio-postsubmit", "/aws/lambda/esp-ENVS-lambda-portfolio-setpurgedate", "/aws/lambda/esp-ENVS-lambda-portfolio-submit", "/aws/lambda/esp-ENVS-lambda-repository-cleanup", "/aws/lambda/esp-ENVS-lambda-repository-precleanup", "/aws/lambda/esp-ENVS-lambda-scoring-monitorincomplete", "/esp-ENVS/ppp-external-webservice", "/esp-ENVS/webservices", "/esp-ENVS/web-ui", "/esp-ENVS/quartz", "/esp-ENVS/external-webservice", "/esp-ENVS/flyway", "/aws/lambda/esp-ENVS-lambda-email-generation", "/aws/lambda/esp-ENVS-lambda-email-notification", "/aws/lambda/esp-ENVS-lambda-importer-checkevidencestatus", "/aws/lambda/esp-ENVS-lambda-importer-checkimportrequeststatus", "/aws/lambda/esp-ENVS-lambda-reposync-post-processing", "/aws/lambda/esp-ENVS-lambda-reposync-conversion-status-check", "/aws/lambda/esp-ENVS-lambda-reposync-original-status-check"}
	array2 := []string{"/aws/lambda/esp-ENVS-lambda-asyncwebservicelambda", "/aws/lambda/esp-ENVS-lambda-email-generation", "/aws/lambda/esp-ENVS-lambda-email-notification", "/aws/lambda/esp-ENVS-lambda-importer-checkevidencestatus", "/aws/lambda/esp-ENVS-lambda-importer-checkimportrequeststatus", "/aws/lambda/esp-ENVS-lambda-importer-handlecancellationerror", "/aws/lambda/esp-ENVS-lambda-importer-monitorincomplete", "/aws/lambda/esp-ENVS-lambda-importer-postevidenceupload", "/aws/lambda/esp-ENVS-lambda-pre-deployment-check", "/aws/lambda/esp-ENVS-lambda-scoring-checkscoresysrequeststatus", "/aws/lambda/esp-ENVS-lambda-oc-monitorincomplete", "/aws/lambda/esp-ENVS-lambda-portfolio-postsubmit", "/aws/lambda/esp-ENVS-lambda-portfolio-setpurgedate", "/aws/lambda/esp-ENVS-lambda-portfolio-submit", "/aws/lambda/esp-ENVS-lambda-repository-cleanup", "/aws/lambda/esp-ENVS-lambda-repository-precleanup", "/aws/lambda/esp-ENVS-lambda-scoring-monitorincomplete", "/esp-ENVS/ppp-external-webservice", "/esp-ENVS/webservices", "/esp-ENVS/web-ui", "/esp-ENVS/quartz", "/esp-ENVS/external-webservice", "/esp-ENVS/flyway", "/aws/lambda/esp-ENVS-lambda-reposync-post-processing", "/aws/lambda/esp-ENVS-lambda-reposync-conversion-status-check", "/aws/lambda/esp-ENVS-lambda-reposync-original-status-check"}
	r := unique(array1)
	fmt.Println(r)
	fmt.Println(array2)
}

func sumForLoop(n int) int {
	i := 1
	prev := 0
	for i <= n {
		prev = prev + i
		i++
	}
	return prev
}

func sumRecursion(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumRecursion(n-1)
}

func unique(arr []string) []string {
	occured := map[string]bool{}
	result := []string{}
	for e := range arr {

		// check if already the mapped
		// variable is set to true or not
		if occured[arr[e]] != true {
			occured[arr[e]] = true

			// Append to result slice.
			result = append(result, arr[e])
		} else {
			fmt.Println("theres duplicate")
			fmt.Println(arr[e])
		}
	}

	return result
}

// func main() {
//     array1 := []int{1, 5, 3, 4, 1, 6, 6, 6, 8, 7, 13, 5}
//     fmt.Println(array1)
//     unique_items := unique(array1)
//     fmt.Println(unique_items)
// }
