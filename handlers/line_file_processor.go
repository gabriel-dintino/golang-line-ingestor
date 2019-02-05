package handlers

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Processor() {
	file, err := os.Open("./files/SRV_LINEAS_DESA.DAT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		fmt.Println(fmt.Sprintf("%d - %s", i, scanner.Text()))
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
