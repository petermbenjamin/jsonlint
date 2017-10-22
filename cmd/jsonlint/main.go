package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/moxar/jsonlint"
)

func main() {
	jsonBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln("could not read JSON from STDIN:", err)
	}

	lintedJSON, err := jsonlint.Lint(jsonBytes)
	if err != nil {
		log.Fatalln("could not lint JSON:", err)
	}

	fmt.Println(string(lintedJSON))
}
