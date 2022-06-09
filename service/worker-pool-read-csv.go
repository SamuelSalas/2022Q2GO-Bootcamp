package service

import (
	"strconv"
	"sync"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
)

type Job struct {
	ExecFn func(csvRow []string) entity.Character
	Args   []string
}

func worker(w *sync.WaitGroup, job <-chan Job, data chan<- entity.Character) {
	for c := range job {
		data <- c.ExecFn(c.Args)
	}
	w.Done()
}

func makeWP(poolSize int, job <-chan Job, data chan<- entity.Character) {
	var w sync.WaitGroup
	for i := 0; i < poolSize; i++ {
		w.Add(1)
		go worker(&w, job, data)
	}
	w.Wait()
	close(data)
}

func create(idType string, poolSize int, data [][]string, jobs chan<- Job, execFn func(csvRow []string) entity.Character) {
	count := 0
	id := 0
	for _, row := range data {
		if count < poolSize {
			id, _ = strconv.Atoi(row[0])
			if idType == "even" {
				if id%2 == 0 {
					j := Job{
						ExecFn: execFn,
						Args:   row,
					}
					jobs <- j
				}
			}

			if idType == "odd" {
				if id%2 == 1 {
					j := Job{
						ExecFn: execFn,
						Args:   row,
					}
					jobs <- j
				}
			}

		}
	}

	close(jobs)
}

func (s *csvService) ReadCsvWorkerPool(idType string, workers, items, itemsWorkerLimit int) (*entity.ResponseBody, error) {
	csv, errs := s.csvRepo.ExtractCsvData()
	if errs != nil {
		return nil, errs
	}

	responseBody := entity.ResponseBody{}
	jobs := make(chan Job, itemsWorkerLimit)
	characters := make(chan entity.Character, itemsWorkerLimit)

	sampleStringTaskFn := func(csvRow []string) entity.Character {
		character := entity.Character{}
		character.ID, _ = strconv.Atoi(csvRow[0])
		character.Name = csvRow[1]
		character.Status = csvRow[2]
		character.Gender = csvRow[3]
		character.Image = csvRow[4]
		character.Url = csvRow[5]
		character.Created = csvRow[6]
		return character
	}

	go create(idType, items, *csv, jobs, sampleStringTaskFn)
	finished := make(chan interface{})
	go func() {
		for d := range characters {
			responseBody.Results = append(responseBody.Results, d)
		}
		finished <- true
	}()

	makeWP(workers, jobs, characters)
	<-finished
	return &responseBody, nil
}
