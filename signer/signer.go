package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

func ExecutePipeline(jobs ...job) {
	start := make(chan interface{}, MaxInputDataLen)
	close(start)
	out := make(chan interface{}, MaxInputDataLen)
	wg := sync.WaitGroup{}
	wg.Add(len(jobs))
	go func(in, out chan interface{}) {
		jobs[0](in, out)
		close(out)
		wg.Done()
	}(start, out)
	for i := 1; i < len(jobs); i++ {
		tempOut := out
		out = make(chan interface{}, MaxInputDataLen)
		go func(i int, in, out chan interface{}) {
			jobs[i](in, out)
			close(out)
			wg.Done()
		}(i, tempOut, out)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()
}

func SingleWorker(in string, out chan interface{}, wg *sync.WaitGroup){
	defer wg.Done()
	var dataMd5 string
	dataCrc32 := make(chan string)
	result32 := make(chan string)
	go func(data5 string) {
		mutex.Lock()
		dataMd5 = DataSignerMd5(data5)
		mutex.Unlock()
		result32 <- DataSignerCrc32(dataMd5)

	}(in)
	go func(data32 string) {
		dataCrc32 <- DataSignerCrc32(data32)
	}(in)
	out <- <-dataCrc32 + "~" + <-result32
}
func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for i := range in {
		wg.Add(1)
		num := strconv.Itoa((i).(int))
		go SingleWorker(num, out, wg)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()
}

func MultiWorker(in string, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	result := make([] chan string, 6)
	data := make([]string, 6)
	var res string
	for i := 0; i < 6; i++ {
		res = strconv.Itoa(i) + in
		result[i] = make(chan string)
		go func(data string, ind int) {
			temp := DataSignerCrc32(data)
			result[ind] <- temp
		}(res, i)
	}
	for i :=0; i < 6; i++{
		data[i] = <-result[i]
	}
	out <- strings.Join(data,"")
}
func MultiHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for i := range in {
		wg.Add(1)
		hash := (i).(string)
		go MultiWorker(hash, out, wg)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()
}

func CombineWorker(in chan interface{}, wg *sync.WaitGroup) string{
	defer wg.Done()
	var res []string
	for i := range in{
		res = append(res, i.(string))
	}
	sort.Strings(res)
	return strings.Join(res,"_")
}
func CombineResults(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	out <- CombineWorker(in, wg)
}
