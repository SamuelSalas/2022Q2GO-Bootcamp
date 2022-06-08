package service

import (
	"context"
	"strconv"
	"sync"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/err"
)

type Result struct {
	Character entity.Character
	Err       error
}

type Job struct {
	TaskId int
	ExecFn func(ctx context.Context, workerId int, csvRow [][]string) (entity.Character, error)
	Args   [][]string
}

type WorkerPool struct {
	workersCount int
	jobs         chan Job
	results      chan Result
	Done         chan struct{}
}

func NewWorkerPool(wcount, itemPerWorker int) WorkerPool {
	return WorkerPool{
		workersCount: wcount,
		jobs:         make(chan Job, itemPerWorker),
		results:      make(chan Result, itemPerWorker),
		Done:         make(chan struct{}),
	}
}

func (j Job) execute(ctx context.Context) Result {
	value, err := j.ExecFn(ctx, j.TaskId, j.Args)
	if err != nil {
		return Result{
			Err: err,
		}
	}

	return Result{
		Character: value,
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			results <- job.execute(ctx)
		case <-ctx.Done():
			results <- Result{
				Err: ctx.Err(),
			}
			return
		}
	}
}

func (wp WorkerPool) Run(ctx context.Context) {
	var wg sync.WaitGroup
	for i := 0; i < wp.workersCount; i++ {
		wg.Add(1)
		go worker(ctx, &wg, wp.jobs, wp.results)
	}

	wg.Wait()
	close(wp.Done)
	close(wp.results)
}

func (wp WorkerPool) Results() <-chan Result {
	return wp.results
}

func (wp WorkerPool) GenerateFrom(jobsBulk []Job) {
	for i := range jobsBulk {
		wp.jobs <- jobsBulk[i]
	}
	close(wp.jobs)
}

func testJobs(poolSize int, data [][]string) []Job {
	jobs := make([]Job, poolSize)
	for i := 0; i < poolSize; i++ {
		jobs[i] = Job{
			TaskId: i,
			ExecFn: execFn,
			Args:   data,
		}

	}

	return jobs
}

func execFn(ctx context.Context, taskId int, csvRow [][]string) (entity.Character, error) {
	result := entity.Character{}
	var err error
	result.ID, _ = strconv.Atoi(csvRow[taskId][0])
	result.Name = csvRow[taskId][1]
	result.Status = csvRow[taskId][2]
	result.Gender = csvRow[taskId][3]
	result.Image = csvRow[taskId][4]
	result.Url = csvRow[taskId][5]
	result.Created = csvRow[taskId][6]
	return result, err
}

func (s *csvService) ReadCsvWorkerPool(idType string, items, itemsWorkerLimit int) (entity.ResponseBody, error) {
	responseBody := entity.ResponseBody{}
	data, errs := s.csvRepo.ExtractCsvData()
	if errs != nil {
		return responseBody, errs
	}

	if len(*data) == 0 {
		return responseBody, err.ErrorCsvEmpty
	}

	processedData := mapData(idType, data)

	wp := NewWorkerPool(5, itemsWorkerLimit)
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	go wp.GenerateFrom(testJobs(items, *processedData))
	go wp.Run(ctx)

	select {
	case r, _ := <-wp.Results():
		responseBody.Results = append(responseBody.Results, r.Character)
	case <-wp.Done:
		return responseBody, errs
	default:
	}
	return responseBody, errs
}

func mapData(idType string, data *[][]string) *[][]string {
	id := 0
	result := [][]string{}
	for _, row := range *data {
		id, _ = strconv.Atoi(row[0])
		if idType == "odd" {
			if id%2 == 1 {
				result = append(result, row)
			}
		}

		if idType == "even" {
			if id%2 == 0 {
				result = append(result, row)
			}
		}
	}
	return &result
}
