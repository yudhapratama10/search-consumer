package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/yudhapratama10/search-consumer/model"
)

var topic string = "test-messages"

func (repo *footballRepository) Insert(data model.FootballClub) (model.FootballClub, error) {

	query := map[string]interface{}{
		"name":        data.Name,
		"tournaments": data.Tournaments,
		"nation":      data.Nation,
		"has_stadium": data.HasStadium,
		"description": data.Description,
		"rating":      data.Rating,
	}

	strQuery, err := json.Marshal(query)
	if err != nil {
		return model.FootballClub{}, err
	}

	fmt.Println(string(strQuery))

	// Perform the index request.
	req := esapi.IndexRequest{
		Index:      "footballclubs",
		DocumentID: strconv.Itoa(data.Id),
		Body:       strings.NewReader(string(strQuery)),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), repo.client)
	defer res.Body.Close()

	if err != nil {
		return model.FootballClub{}, err
	}

	return data, nil
}

func (repo *footballRepository) Update(data model.FootballClub) (model.FootballClub, error) {

	query := map[string]interface{}{
		"name":        data.Name,
		"tournaments": data.Tournaments,
		"nation":      data.Nation,
		"has_stadium": data.HasStadium,
		"description": data.Description,
		"rating":      data.Rating,
	}

	strQuery, err := json.Marshal(query)
	if err != nil {
		return model.FootballClub{}, err
	}

	// fmt.Println(string(strQuery))

	// Perform the index request.
	req := esapi.IndexRequest{
		Index:      "footballclubs",
		DocumentID: strconv.Itoa(data.Id),
		Body:       strings.NewReader(string(strQuery)),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), repo.client)
	defer res.Body.Close()
	// fmt.Println(err)

	if err != nil {
		return model.FootballClub{}, err
	}

	return data, nil
}

func (repo *footballRepository) Delete(data model.FootballClub) (model.FootballClub, error) {

	// Perform the delete request.
	req := esapi.DeleteRequest{
		Index:      "footballclubs",
		DocumentID: strconv.Itoa(data.Id),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), repo.client)
	defer res.Body.Close()

	if err != nil {
		return model.FootballClub{}, err
	}

	return data, nil
}
