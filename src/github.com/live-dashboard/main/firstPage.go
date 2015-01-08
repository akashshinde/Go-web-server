package main

import (
	//"github.com/gin-gonic/gin"
	//"io"
	"io/ioutil"
	//"github.com/go-martini/martini"
	"encoding/json"
	"log"
	"net/http"
)

type Results struct {
	elements []Elements `json:"elements"`
}

type Elements struct {
	denominator int32   `json:"denominator"`
	numerator   int32   `json:"numerator"`
	name        string  `json:"name"`
	ratio       float32 `json:"ratio"`
}

func getJSON() {
	resp, _ := http.Get("http://controller.c3.mtv/job/Live-Usher-Cobertura/345/cobertura/api/json?depth=2")
	//JsonBody := resp.Body
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("data s-is : %s", string(body))

	//decoder := json.NewDecoder(resp.Body)

	var r Results
	_ = json.Unmarshal(body, &r)
	/*for {
		if err := decoder.Decode(&r); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}*/
	//fmt.Println(r.elements.name[1])
	log.Printf("%+v", r)
}

func main() {
	/*
		router := gin.Default()
		router.GET("/", func(c *gin.Context) {
			c.String(200, "hello world")
		})
		router.Run(":8080")
	*/
	getJSON()
}
