package utilities

import (
	"fmt"
"regexp"
	"log"
	"os"
"encoding/json"
"reflect"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

// fire http requests
// unmarshall func
// compare response  functions
// header prepation functiond
// update headers function

func Get_common_headers() map[string]string {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	return headers
}

func Fire_get_request(headers map[string]string, qparams string, endpoint string) (response *resty.Response){

	client := resty.New()
	resp, err := client.R().
		SetQueryString(qparams).
		SetHeaders(headers).
		Get(endpoint)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	return resp
}

func Fire_post_request(headers map[string]string, payload string, endpoint string) (response *resty.Response){

	client := resty.New()
	resp, err := client.R().
		
		SetHeaders(headers).
		SetBody(payload).
		Post(endpoint)

	if err != nil {
		log.Println("error:" , err)
	}
	log.Println(resp)
	return resp
}

func Fire_put_request(headers map[string]string, payload string, endpoint string) (response *resty.Response){

	client := resty.New()
	resp, err := client.R().
		
		SetHeaders(headers).
		SetBody(payload).
		Put(endpoint)

	if err != nil {
		log.Println("error:" , err)
	}
	log.Println(resp)
	return resp
}

func Fire_delete_request(headers map[string]string, endpoint string) (response *resty.Response){

	client := resty.New()
	resp, err := client.R().
		
		SetHeaders(headers).
		Delete(endpoint)

	if err != nil {
		log.Println("error:" , err)
	}
	log.Println(resp)
	return resp
}

func Get_env_val(val string) string {
	return viper.Get(val).(string)
}


func Get_testtarget_envname() string {
	return os.Getenv("env")
}

func Load_envconfig(){

	log.Println("Loading env config")
	
	viper.SetConfigType("env")
	viper.SetConfigFile("../../"+ "dev.env")
	
	//viper.ReadInConfig()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %w \n", err)
	}

	log.Println("Loaded the env config file > " + Get_testtarget_envname())
	log.Println(viper.Get("registerpostendpoint"))
	
	
}


func ReadFile(partialfilepath string) (string, error) {
	data, err := os.ReadFile("../../data/" +  partialfilepath)
	if err != nil {
		log.Println(" File not found")
		return "", err
	}
	temp:= string(data)
	re := regexp.MustCompile(` +\r?\n +`)
	temp1:= re.ReplaceAllString(temp, "")
	log.Println(temp1)
	return temp1, nil
}

func JSONBytesEqual(a, b []byte) (bool, error) {
    var j, j2 interface{}
    if err := json.Unmarshal(a, &j); err != nil {
        return false, err
    }
    if err := json.Unmarshal(b, &j2); err != nil {
        return false, err
    }
    return reflect.DeepEqual(j2, j), nil
}
