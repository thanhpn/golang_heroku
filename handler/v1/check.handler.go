package v1

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type PageProps struct {
	IsEligible bool `json:"isEligible"`
}
type DetailResp struct {
	PageProps PageProps `json:"pageProps"`
}

func Health(ctx *gin.Context) {
	response := map[string]string{
		"message": "ok!",
	}
	ctx.JSON(http.StatusOK, response)
}

func Address(ctx *gin.Context) {
	address := ctx.Param("address")

	isAirdrop := checkAddress(address)
	if isAirdrop {
		response := map[string]string{
			"message":    "ok!",
			"address":    address,
			"isEligible": "true",
		}
		ctx.JSON(http.StatusOK, response)
		return
	}

	requestURL := fmt.Sprintf("https://arbitrum.foundation/_next/data/lLwjZPqfFxwx0LQ3IfWRy/eligibility.json?address=%v", address)
	fmt.Printf("requestURL\n" + requestURL)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/112.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("x-nextjs-data", "1")
	req.Header.Set("TE", "trailers")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}
	fmt.Printf("client: response body: %s\n", resBody)

	defer res.Body.Close()
	var detail DetailResp
	err = json.Unmarshal(resBody, &detail)
	if err != nil {
		return
	}

	response := map[string]string{
		"message":    "ok!",
		"address":    address,
		"isEligible": fmt.Sprintf("%v", detail.PageProps.IsEligible),
	}
	ctx.JSON(http.StatusOK, response)
}

var addressLines []string

func CheckWhitelistAddress(ctx *gin.Context) {
	loadWhitelistAddress()
}

func checkAddress(address string) bool {
	loadWhitelistAddress()
	for _, addressLine := range addressLines {
		res := strings.EqualFold(address, addressLine)
		if res {
			return true
		}
	}
	return false
}

func loadWhitelistAddress() {
	if len(addressLines) > 0 {
		fmt.Println("address loaded")
		return
	}
	file, err := os.Open("address.txt")

	if err != nil {
		fmt.Printf("failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		addressLines = append(addressLines, scanner.Text())
	}
	file.Close()

}
