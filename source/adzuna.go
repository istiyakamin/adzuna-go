package source

import (
	"adzuna/config"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type RequestData struct {
	Page  string `json:"page"` 
	What   string  `json:"what"`
	Where  string  `json:"where"`
	Results_per_page string `json:"results_per_page"`
}


type ResponseData struct {
	Class   string  `json:"__CLASS__"`
	Count   int     `json:"count"`
	Mean    float64 `json:"mean"`
	Results []struct {
		Created     time.Time `json:"created"`
		Title       string    `json:"title"`
		SalaryMax   int       `json:"salary_max"`
		RedirectURL string    `json:"redirect_url"`
		Description string    `json:"description"`
		Location    struct {
			DisplayName string   `json:"display_name"`
			Class       string   `json:"__CLASS__"`
			Area        []string `json:"area"`
		} `json:"location"`
		Longitude float64 `json:"longitude,omitempty"`
		SalaryMin int     `json:"salary_min"`
		Class     string  `json:"__CLASS__"`
		Latitude  float64 `json:"latitude,omitempty"`
		Category  struct {
			Class string `json:"__CLASS__"`
			Tag   string `json:"tag"`
			Label string `json:"label"`
		} `json:"category"`
		SalaryIsPredicted string `json:"salary_is_predicted"`
		Adref             string `json:"adref"`
		ID                string `json:"id"`
		ContractTime      string `json:"contract_time"`
		Company           struct {
			DisplayName string `json:"display_name"`
			Class       string `json:"__CLASS__"`
		} `json:"company"`
	} `json:"results"`
}

func Adzuna(req RequestData) ResponseData {

	if req.Page == "" || req.Page == "0" {
		req.Page = "1"
	}

	if req.Results_per_page == "" || req.Results_per_page == "0" {
		req.Results_per_page = "10"
	}

	baseURL := "https://api.adzuna.com"
    resource := "/v1/api/jobs/us/search/"+ req.Page
    params := url.Values{}
    params.Add("app_id", config.ADZUNA_APP_ID)
    params.Add("app_key", config.ADZUNA_APP_KEY)
	params.Add("results_per_page", req.Results_per_page)

	if len(req.What) > 0 {
		params.Add("what", req.What)
	}

	if len(req.Where) > 0 {
		params.Add("what", req.Where)
	}

    u, _ := url.ParseRequestURI(baseURL)
    u.Path = resource
    u.RawQuery = params.Encode()
    urlStr := fmt.Sprintf("%v", u) // "http://example.com/path?param1=value1&param2=value2"
	fmt.Println(urlStr)

	res, err := http.Get(urlStr)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(responseData)

	myJsonString := responseData;

	var adzuna_data ResponseData;
	json.Unmarshal([]byte(myJsonString), &adzuna_data)
	// // fmt.Printf("%T", &adzuna_data.Results)
	// for key, value := range adzuna_data.Results {
	// 	fmt.Println(key, ":", value.Title)
	// }

	return adzuna_data
}