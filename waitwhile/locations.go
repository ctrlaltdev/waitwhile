package waitwhile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ctrlaltdev/waitwhile/utils"
)

type Location struct {
	BusinessName     string `json:"businessName"`
	BusinessType     string `json:"businessType"`
	Created          string `json:"created"`
	CreatedBy        string `json:"createdBy"`
	Email            string `json:"email"`
	ID               string `json:"id"`
	IsActive         bool   `json:"isActive"`
	IsForceClosed    bool   `json:"isForceClosed"`
	IsForceOpen      bool   `json:"isForceOpen"`
	IsPublicBooking  bool   `json:"isPublicBooking"`
	IsPublicCheckIn  bool   `json:"isPublicCheckIn"`
	IsPublicWaitlist bool   `json:"isPublicWaitlist"`
	Name             string `json:"name"`
	ShortName        string `json:"shortName"`
	Updated          string `json:"updated"`
	UpdatedBy        string `json:"updatedBy"`
	Locale           string `json:"locale"`
}

func GetLocation(id *string) Location {
	url := baseUrl + "/locations/" + *id

	req, err := http.NewRequest("GET", url, nil)
	utils.CheckErr(err)

	req.Header.Set("apikey", *apiKey)

	client := &http.Client{Timeout: time.Second * 10}

	res, err := client.Do(req)
	utils.CheckErr(err)

	defer res.Body.Close()

	var data Location
	err = json.NewDecoder(res.Body).Decode(&data)
	utils.CheckErr(err)

	return data
}

type GetLocationsResponse struct {
	EndAt   string     `json:"endAt"`
	Limit   int        `json:"limit"`
	Results []Location `json:"results"`
	StartAt string     `json:"startAt"`
}

func GetLocations(limit int, startAfter *string) GetLocationsResponse {
	urlStr := baseUrl + "/locations"

	u, err := url.Parse(urlStr)
	utils.CheckErr(err)

	q, err := url.ParseQuery(u.RawQuery)
	utils.CheckErr(err)

	q.Add("limit", fmt.Sprint(limit))
	if startAfter != nil {
		q.Add("startAfter", *startAfter)
	}

	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	utils.CheckErr(err)

	req.Header.Set("apikey", *apiKey)

	client := &http.Client{Timeout: time.Second * 10}

	res, err := client.Do(req)
	utils.CheckErr(err)

	defer res.Body.Close()

	var data GetLocationsResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	utils.CheckErr(err)

	return data
}

func UpdateLocation(id *string, payload *string) error {
	url := baseUrl + "/locations/" + *id

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(*payload)))
	if err != nil {
		return err
	}

	req.Header.Set("apikey", *apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("the API returned %d: %s", res.StatusCode, http.StatusText(res.StatusCode))
	}

	return nil
}

func DeleteLocation(id string) error {
	url := baseUrl + "/locations/" + id

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("apikey", *apiKey)

	client := &http.Client{Timeout: time.Second * 10}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("the API returned %d: %s", res.StatusCode, http.StatusText(res.StatusCode))
	}

	return nil
}
