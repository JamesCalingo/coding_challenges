package loadtester

import "fmt"

func CheckStatus(url string) string {
	status := ""
	fmt.Printf("Your request for %s returned: ", url)
	// http.NewRequest("GET", url, io.Reader())
	return status
}
