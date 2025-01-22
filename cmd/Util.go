package cmd

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"
)

func getRecord(id int) (Task, error) {
	fin, err := os.Open("tasks.csv")
	defer fin.Close()
	if err != nil {
		return Task{}, err
	}

	r := csv.NewReader(fin)
	for i := 0; ;i++ {
		if i == id {
			break
		}
		r.Read()
	}
	record, err := r.Read()
	var (
		ID, _ = strconv.Atoi(record[0])
		desc = record[1]
		createdAt, _ = time.Parse(time.RFC1123, record[2])
		completedAt, _ = time.Parse(time.RFC1123, record[3])
		completed, _ = strconv.ParseBool(record[4])
	)
	result := Task{ID, desc, createdAt, completedAt, completed}
	return result, nil
}

func updateRecord(id int, task string, createdAt time.Time, completedAt time.Time, completed bool) error {
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

	for i := 0; ;i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if i == id {
			record[1] = task
			record[2] = createdAt.Format(time.RFC1123)
			record[3] = completedAt.Format(time.RFC1123)
			record[4] = strconv.FormatBool(completed)
		}
		w.Write(record)
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

func updateData(nt int) error {
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

	for i := 0; ;i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if i == 0 {
			record[0] = strconv.Itoa(nt)
		}
		w.Write(record)
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