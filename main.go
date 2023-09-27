package main
import (
	"encoding/json"
	"fmt"
)
// ###################################### Main Function ######################################
func main() {
	
	var json_data = make(map[string]string)
	json_data["last"] = "arora"
	json_data["no1"] = "ankit"
	json_data["age"] = "28"
	json_data["zebra"] = "stripes"
	json_content,err := json.Marshal(json_data)
	fmt.Println(string(json_content))

	
}

