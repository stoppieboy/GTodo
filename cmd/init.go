package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
)

func initCmd(){
	file, err := os.Create("tasks.csv")

	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	entry := []string{"0", "", "", "", ""} // {No. of tasks,,,,}
	writer.Write(entry)
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error writing CSV data: ", err)
		return
	}
	return
}
