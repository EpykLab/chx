package shared

import (
	"encoding/json"
	"fmt"
	"log"
)

func Out(data any) {

	v, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	output := string(v)
	fmt.Println(output)
}
