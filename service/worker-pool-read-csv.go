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
	ExecFn func(ctx context.Context, csvRow []string) (entity.Character, error)
	Args   []string
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
	value, err := j.ExecFn(ctx, j.Args)
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

func testJobs(idType string, poolSize int, data [][]string) []Job {
	jobs := make([]Job, poolSize)
	count := 0
	id := 0
	for _, row := range data {
		for count < poolSize {
			id, _ = strconv.Atoi(row[0])
			if idType == "even" {
				if id%2 == 0 {
					jobs[count] = Job{
						ExecFn: execFn,
						Args:   row,
					}
					count++
				}
			}

			if idType == "odd" {
				if id%2 == 1 {
					jobs[count] = Job{
						ExecFn: execFn,
						Args:   row,
					}
				}
			}
		}
		break
	}

	return jobs
}

func execFn(ctx context.Context, csvRow []string) (entity.Character, error) {
	result := entity.Character{}
	var err error
	result.ID, _ = strconv.Atoi(csvRow[0])
	result.Name = csvRow[1]
	result.Status = csvRow[2]
	result.Gender = csvRow[3]
	result.Image = csvRow[4]
	result.Url = csvRow[5]
	result.Created = csvRow[6]
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

	wp := NewWorkerPool(5, itemsWorkerLimit)
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	go wp.GenerateFrom(testJobs(idType, items, *data))
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
