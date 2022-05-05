package checker

import (
	"fmt"
	"log"
	"net/http"
)

func CheckWebServer(urlStr string) {
	resp, err := http.Get(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The status code we got is:", resp.StatusCode)
}

//func CheckWebServer(urlStr string) (code int, err error) {
//	resp, err := http.Get(urlStr)
//	if err != nil {
//		//log.Fatal(err)
//		return 0, err
//	}
//	return resp.StatusCode, nil
//	//fmt.Println("The status code we got is:", resp.StatusCode)
//}
//
//func CheckBulk(listUrls []string) map[string]int {
//	result := make(map[string]int)
//	for _, u := range listUrls {
//		res, err := CheckWebServer(u)
//		if err != nil {
//			result[u] = 1000
//		} else {
//			result[u] = res
//		}
//	}
//	return result
//}
