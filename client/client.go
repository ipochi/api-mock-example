package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ipochi/api-mock-example/model"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	HttpClient *http.Client
}

func (c *Client) GetCompanies() ([]model.Company, error) {
	rel := &url.URL{Path: "/api/v1/companies"}

	fmt.Println("Rel ---- ", rel)
	u := c.BaseURL.ResolveReference(rel)

	fmt.Println("What is u ---", u)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("resp body", string(body))
	var users []model.Company
	err = json.NewDecoder(resp.Body).Decode(&users)
	return users, err
}
