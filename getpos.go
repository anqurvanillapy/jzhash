package main

import (
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	// STOCK_URL is the JSONP API for today's stock query of Tencent.
	STOCK_URL = "http://hq.sinajs.cn/list=hk00700"

	nRow    = 19
	nCol    = 8
	Colmask = nCol - 1
)

// GetPosition returns numeric data of n-th row and column.
func GetPosition() (int, int) {
	var price string

	client := &http.Client{}
	req, _ := http.NewRequest("GET", STOCK_URL, nil)
	req.Header.Set("User-Agent", "Chrome/61.0")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonpStrings := strings.Split(string(body), ",")
	for _, s := range jsonpStrings {
		_, err := strconv.ParseFloat(s, 64)
		if err == nil {
			price = s
			break
		}
	}
	if price == "" {
		log.Fatal("Get opening price error")
	}
	dateprice := fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02"), price)
	hash := sha512.New()
	hash.Write([]byte(dateprice))
	hashBytes := hash.Sum(nil)
	nthRow := (int)(binary.BigEndian.Uint64(hashBytes[:33]) % nRow)
	nthCol := (int)(binary.BigEndian.Uint64(hashBytes[32:]) & Colmask)

	return nthRow, nthCol
}
