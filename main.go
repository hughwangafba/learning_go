package main

import (
	"fmt"
	"github.com/elgs/gojq"
)

func main() {
	inputs :=
		`[ 
             {
                "name" : "water",
                 "type": "liquid",
			     "boiling_point": {
			       "units" : "C",
			       "value":100
			      }
             },
            {
                "name" : "carbon monoxide",
                 "type": "gas",
                 "dangerous":true,
				"features": [
					 {
					  "units" : "C",
					  "value":100
					 },
					 {
						"type":"density",
						"units":"kg/m3",
						"value":389
					  }
                 ]
            }
         ]
         `
	parser, err := gojq.NewStringQuery(inputs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parser.Query("[0].name"))
	fmt.Println(parser.Query("[1].features.[1].type"))

}
