package main

import (
	"log"

	"github.com/yudhapratama10/search-consumer/model"
	"github.com/yudhapratama10/search-consumer/repository"
	"github.com/yudhapratama10/search-consumer/usecase"
)

func consume(data model.Message) {

	repo := repository.NewFootballRepository(elasticClient)
	usecase := usecase.NewFootballClubUsecase(repo)

	if data.Operation == "insert" {
		// Inserting to ES
		_, err = usecase.Insert(data.Data)
		if err != nil {
			log.Fatal("Error on inserting docs: ", err)
		}

	} else if data.Operation == "update" {

		_, err = usecase.Update(data.Data)
		if err != nil {
			log.Fatal("Error on updating docs: ", err)
		}

	} else if data.Operation == "delete" {

		_, err = usecase.Delete(data.Data)
		if err != nil {
			log.Fatal("Error on deleting docs: ", err)
		}

	}
}
