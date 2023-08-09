package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func sendHello() {
	resp, err := http.Get("http://localhost:8080/test/whoami")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func headLyrics() int {
	req, _ := http.NewRequest("HEAD", "http://localhost:8080/download/lyrics", nil)
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	length := resp.Header.Get("Content-Length")
	result, _ := strconv.Atoi(length)
	return result
}

func downloadPrallel() {
	length := headLyrics()
	for length > 0 {
		fmt.Printf("length: %d\n", length)
		/** download segment */
		from := length - 100
		if from < 0 {
			from = 0
		}
		downloadSegment(from, length)
		length -= 100
	}
	// time.Sleep(time.Second * 10)
}

func downloadSegment(from, to int) {
	/** download segment */
	req, _ := http.NewRequest("GET", "http://localhost:8080/download/lyrics/segment", nil)
	req.Header.Set("Range", "bytes="+strconv.Itoa(from)+"-"+strconv.Itoa(to))
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	putFile(body, from)
}

func putFile(body []byte, offset int) {
	/** write to local */
	filepath := "./statics/target.txt"
	file, _ := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	file.WriteAt(body, int64(offset))
	// write := bufio.NewWriter(file)
	// write.WriteString(string(body))
	// write.Flush()
	// fmt.Println("Download segment success!")
}
