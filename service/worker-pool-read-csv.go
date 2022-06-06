package service

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
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
		queue: make(chan work),
	}

	gp.AddWorkers(workerSize, itemsPerWorkers)
	return gp
}

func (gp *GoroutinePool) Close() {
	close(gp.queue)
	gp.wg.Wait()
}

func (gp *GoroutinePool) ScheduleWork(fn WorkFunc) {
	gp.queue <- work{fn}
}

func (gp *GoroutinePool) AddWorkers(numWorkers, itemsWorker int) {
	gp.wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(workerId int) {
			count := 0
			for job := range gp.queue {
				job.fn.Run()
				count++
			}
			fmt.Println(fmt.Sprintf("Worker %d executed %d tasks", workerId, count))
			gp.wg.Done()
		}(i)
	}
}

type testTask struct {
	CsvLine       int
	TaskProcessor func(...interface{})
}

func (t testTask) Run() {
	t.TaskProcessor(t.CsvLine)
}

func (c *csvService) ReadCsvWorkerPool(data [][]string, items, itemsWorkerLimit int) (*entity.ResponseBody, error) {
	responseBody := entity.ResponseBody{}
	if len(data) == 0 {
		return nil, repository.ErrorCsvEmpty
	}

	pool := NewGoroutinePool(5, itemsWorkerLimit)
	taskSize := items

	wg := &sync.WaitGroup{}
	wg.Add(taskSize)
	sampleStringTaskFn := func(dm ...interface{}) {
		if row, ok := dm[0].(int); ok {
			var rec entity.Character = entity.Character{}
			rec.ID, _ = strconv.Atoi(data[row][0])
			rec.Name = data[row][1]
			rec.Status = data[row][2]
			rec.Gender = data[row][3]
			rec.Image = data[row][4]
			rec.Url = data[row][5]
			rec.Created = data[row][6]
			responseBody.Results = append(responseBody.Results, rec)
			wg.Done()
		}
	}

	var tasks []testTask
	for v := 0; v < taskSize; v++ {
		tasks = append(tasks, testTask{
			CsvLine:       v,
			TaskProcessor: sampleStringTaskFn,
		})
	}

	for _, task := range tasks {
		pool.ScheduleWork(task)
	}
	pool.Close()
	wg.Wait()
	return &responseBody, nil
}
