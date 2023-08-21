package main
 
import (
    "encoding/json"
    "io/ioutil"
    "log"
)
 
// The data struct for the decoded data
// Notice that all fields must be exportable!
type Data struct {
    Origin string
    User   string
    Active bool
	List []string
}

func readJsonFile() Data {
	 // Let's first read the `config.json` file
	 content, err := ioutil.ReadFile("./config.json")
	 if err != nil {
		 log.Fatal("Error when opening file: ", err)
	 }
  
	 // Now let's unmarshall the data into `payload`
	 var payload Data
	 err = json.Unmarshal(content, &payload)
	 if err != nil {
		 log.Fatal("Error during Unmarshal(): ", err)
		
	 }
	
	 return payload
}
 
func main() {
	var config Data
	config = readJsonFile()

 
    // Let's print the unmarshalled data!
    log.Printf("origin: %s\n", config.Origin)
    // log.Printf("user: %s\n", payload.User)
    // log.Printf("status: %t\n", payload.Active)
	// log.Printf("list: %#v\n", payload.List[0])
}