package service

import (
	"strconv"
	"sync"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
)

type WorkFunc interface {
	Run()
}

type GoroutinePool struct {
	queue chan work
	wg    sync.WaitGroup
}

type work struct {
	fn WorkFunc
}

func NewGoroutinePool(workerSize, itemsPerWorkers int) *GoroutinePool {
	gp := &GoroutinePool{
		queue: make(chan work, itemsPerWorkers),
	}

	gp.AddWorkers(workerSize)
	return gp
}

func (gp *GoroutinePool) Close() {
	close(gp.queue)
	gp.wg.Wait()
}

func (gp *GoroutinePool) ScheduleWork(fn WorkFunc) {
	gp.queue <- work{fn}
}

func (gp *GoroutinePool) AddWorkers(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go func() {
			gp.wg.Add(1)
			for job := range gp.queue {
				job.fn.Run()
			}
			gp.wg.Done()
		}()
	}
}

type Job struct {
	TaskId int
	ExecFn func(csvRow []string)
	Args   []string
}

func testJobs(idType string, poolSize int, data [][]string, execFn func(csvRow []string)) []Job {
	jobs := []Job{}
	count := 0
	id := 0
	for _, row := range data {
		if count < poolSize {
			id, _ = strconv.Atoi(row[0])
			if idType == "even" {
				if id%2 == 0 {
					jobs = append(jobs, Job{
						ExecFn: execFn,
						Args:   row,
					})
					count++
				}
			}

			if idType == "odd" {
				if id%2 == 1 {
					jobs = append(jobs, Job{
						ExecFn: execFn,
						Args:   row,
					})
					count++
				}
			}
		}
	}

	return jobs
}

func (t Job) Run() {
	t.ExecFn(t.Args)
}

func (s *csvService) ReadCsvWorkerPool(idType string, items, itemsWorkerLimit int) (*entity.ResponseBody, error) {
	data, errs := s.csvRepo.ExtractCsvData()
	if errs != nil {
		return nil, errs
	}

	responseBody := entity.ResponseBody{}
	wp := NewGoroutinePool(5, itemsWorkerLimit)
	wg := &sync.WaitGroup{}

	sampleStringTaskFn := func(csvRow []string) {
		result := entity.Character{}
		result.ID, _ = strconv.Atoi(csvRow[0])
		result.Name = csvRow[1]
		result.Status = csvRow[2]
		result.Gender = csvRow[3]
		result.Image = csvRow[4]
		result.Url = csvRow[5]
		result.Created = csvRow[6]
		responseBody.Results = append(responseBody.Results, result)
		wg.Done()
	}

	tasks := testJobs(idType, items, *data, sampleStringTaskFn)
	for _, task := range tasks {
		wg.Add(1)
		wp.ScheduleWork(task)
	}

	wp.Close()
	wg.Wait()

	return &responseBody, nil
}
