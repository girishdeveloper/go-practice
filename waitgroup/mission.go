package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mission/visit"
	"sync"
)

type Task struct {
	Date   string
	Visits []visit.Visit
}

type DailyStats struct {
	Date   string         `json:"date"`
	ByPage map[string]int `json:"byPage"`
}

func main() {
	data, err := ioutil.ReadFile("./mission-input.json")
	if err != nil {
		fmt.Println("Error while loading data", err)
	} else {
		fmt.Println(len(data))
	}
	dayStats := make(map[string][]visit.Visit)
	err = json.Unmarshal(data, &dayStats)
	if err != nil {
		fmt.Println("Error while decoding the json", err)
	}
	var waitGroup sync.WaitGroup
	inputChan := make(chan Task, 10)
	outputChan := make(chan DailyStats, len(dayStats))
	waitGroup.Add(len(dayStats))

	for i := 0; i < len(data); i++ {
		//create worker
		go taskWorker(inputChan, i, outputChan, &waitGroup)
	}
	for date, visits := range dayStats {
		inputChan <- Task{
			Date:   date,
			Visits: visits,
		}
	}
	//won't send any more task to the input channel
	close(inputChan)
	waitGroup.Wait()
	close(outputChan)
	//collect the result
	done := make([]DailyStats, 0, len(dayStats))
	for out := range outputChan {
		done = append(done, out)
	}
	res, err := json.Marshal(done)
	if err != nil {
		log.Fatal(err)
	}
	// write result in the output file
	err = ioutil.WriteFile("results.json", res, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done!")
}

func taskWorker(inputChan chan Task, workerId int, outputChan chan DailyStats, w8 *sync.WaitGroup) {
	for received := range inputChan {
		m := make(map[string]int)
		for _, v := range received.Visits {
			m[v.Page]++
		}
		outputChan <- DailyStats{
			Date:   received.Date,
			ByPage: m,
		}
		fmt.Printf("[worker %d] finished task\n", workerId)
	}
	log.Printf("worker %d quit\n", workerId)
	if workerId < 0 {
		fmt.Printf("negative index encountered %d\n", workerId)
	} else {
		w8.Done()
	}
}
