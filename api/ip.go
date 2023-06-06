package api

import (
	"encoding/json"
	"io"
	"net/http"
)

const apikey = "X-RapidAPI-Key"
const apihost = "X-RapidAPI-Host"

type RapidApi struct {
	RapidKey     string
	RapidKeyVal  string
	RapidHost    string
	RapidHostVal string
}
type IpInfo struct {
	Vaild        bool   `json:"is_vaild"`
	Country      string `json:"country"`
	Country_Code string `json:"country_code"`
	Region_Code  string `json:"region_code"`
	Region       string `json:"region"`
	City         string `json:"city"`
	Zip          string `json:"zip"`
	Time_zone    string `json:"timezone"`
	Isp          string `json:"isp"`
	Address      string `json:"address"`
}

func NewRapidApi(keyval, hostval string) *RapidApi {
	return &RapidApi{RapidKey: apikey, RapidKeyVal: keyval, RapidHost: apihost, RapidHostVal: hostval}
}
func (s *RapidApi) GetIpInfo(ip string) (IpInfo, error) {
	url := "https://ip-lookup-by-api-ninjas.p.rapidapi.com/v1/iplookup?address="
	req, err := http.NewRequest("GET", url+ip, nil)
	if err == nil {
		req.Header.Add(s.RapidKey, s.RapidKeyVal)
		req.Header.Add(s.RapidHost, s.RapidHostVal)
		resp, err := http.DefaultClient.Do(req)
		if err == nil {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err == nil {
				ans := new(IpInfo)
				err = json.Unmarshal(body, &ans)
				if err == nil {
					return *ans, nil
				}
			}
		}
	}
	return IpInfo{}, err
}
