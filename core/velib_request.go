package core

import (
	"../model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

/* This func will fetch the current status of the velibs station in a 500 m around Splio's HQ
 and return it as a go Response struct by Unmarshalling the json inside one */
func GetVelibRecords() model.Response {
	url := "https://opendata.paris.fr/api/records/1.0/search/?dataset=velib-disponibilite-en-temps-reel&q=&lang=en&sort=-dist&geofilter.distance=48.870986%2C2.3353436%2C500&timezone=Europe%2FParis"
	response, err := http.Get(url)
	if err != nil {
		log.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err)
	}

	var result model.Response
	jsonErr := json.Unmarshal(responseData, &result)
	if jsonErr!= nil {
		log.Fatal(jsonErr)
	}

	/* Every record formatted*/
	for i := 0; i<len(result.Records); i++ {
		(&result.Records[i]).FormatTimestamp()
		(&result.Records[i]).FormatDistance()
	}

	return result
}
