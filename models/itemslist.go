package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Item struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ItemResponse struct {
	Item *Item `json:"item"`
}

type ItemsResponse struct {
	Items *[]Item `json:"items"`
}

type ItemslistModel struct {
	Endpoint string
}

func (m *ItemslistModel) GetAll() (*[]Item, error) {
	resp, err := http.Get(m.Endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var itemsResp ItemsResponse
	err = json.Unmarshal(data, &itemsResp)
	if err != nil {
		return nil, err
	}

	return itemsResp.Items, nil
}

func (m *ItemslistModel) Get(id int64) (*Item, error) {
	url := fmt.Sprintf("%s/%d", m.Endpoint, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var itemResp ItemResponse
	err = json.Unmarshal(data, &itemResp)
	if err != nil {
		return nil, err
	}

	return itemResp.Item, nil
}
