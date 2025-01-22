package cmd

import (
	// "bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID int
	desc string
	createdAt time.Time
	completedAt time.Time
	completed bool
}

// Add a task to the list 
// 1. determine the number of tasks in the list
// 2. add the task with the correct ID
func Add(task string) error {
	file, err := os.OpenFile("tasks.csv", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	if err != nil {
		return err
	}
	reader := csv.NewReader(file)
	data, err := reader.Read()
	num, err := strconv.Atoi(data[0])

	writer := csv.NewWriter(file)

	entry := []string{strconv.Itoa(num+1), task, time.Now().Format(time.RFC1123), "",strconv.FormatBool(false)} // {ID, Task, createdAt, completedAt, Completed}
	writer.Write(entry)
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error writing CSV data: ",err)
		return err
	}
	file.Close()

	updateData(num+1)
	

	return nil
}

// Delete a task from the list
// 1. find the task with the given ID
// 2. delete the task
// 3. update the IDs of the remaining tasks
func Delete(id int) error {
	fin, err := os.Open("tasks.csv")
	defer fin.Close()
	if err != nil {
		return err
	}
	fout, err := os.Create("tasks.temp.csv")
	defer fout.Close()
	if err != nil {
		return err
	}

	r := csv.NewReader(fin)
	w := csv.NewWriter(fout)

	var index int = 1
	for i := 0; ;i++ {
		record, err := r.Read()
		if i == 0 {
			currentNo, _ := strconv.Atoi(record[0])
			record[0] = strconv.Itoa(currentNo-1)
			w.Write(record)
			continue
		}
		if err == io.EOF {
			break
		}
		if i != id {
			record[0] = strconv.Itoa(index)
			w.Write(record)
			index++
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}

	fin.Close()
	err = os.Remove("tasks.csv")
	if err != nil {
		return err
	}
	fout.Close()
	err = os.Rename("tasks.temp.csv", "tasks.csv")
	if err != nil {
		return err
	}

	return nil
}

// Complete a task
// 1. find the task with the given ID
// 2. update the task's completed field to true
func Complete(id int) error {
	task, err := getRecord(id)
	if err != nil {
		return err
	}
	task.completed = true
	err = updateRecord(id, task.desc, task.createdAt, time.Now(), task.completed)
	if err != nil {
		return err
	}
	return nil
}

func List() ([]string, error) {

	file, err := os.Open("tasks.csv")
	defer file.Close()
	if err != nil {
		return nil, err
	}

	// data, err := io.ReadAll(file)
	// if err != nil {
	// 	return "", err
	// }

	reader := csv.NewReader(file)
	formattedData := make([]string, 0)
	for i := 0; ;i++ {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}else if err != nil {
			fmt.Println("Error reading CSV data: ",err)
			break
		}else if i == 0 {
			continue
		}
		var status string = "Pending"
		if record[4] == "true" {
			status = "Completed"
		}
		formattedData =  append(formattedData, record[0]+"\t"+record[1]+"\t"+status+"\n")
	}
	return formattedData, nil
}