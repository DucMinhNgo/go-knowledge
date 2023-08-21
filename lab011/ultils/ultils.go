package ultils

import (
	"encoding/json"
    "io/ioutil"
    "log"
)

type ConfigStruct struct {
	Origin string
    User   string
    Active bool
}

func get_config() {
	// Let's first read the `config.json` file
    content, err := ioutil.ReadFile("../config.json")
    if err != nil {
        log.Fatal("Error when opening file: ", err)
    }
 
    // Now let's unmarshall the data into `payload`
    var payload ConfigStruct
    err = json.Unmarshal(content, &payload)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }
 
    // Let's print the unmarshalled data!
    log.Printf("origin: %s\n", payload.Origin)
    log.Printf("user: %s\n", payload.User)
    log.Printf("status: %t\n", payload.Active)
}