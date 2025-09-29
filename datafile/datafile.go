package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloat(fileName string) ([3]float64, error) {
	var numers [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numers, err
	}

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numers, err
	}

	if scanner.Err() != nil {
		return numers, scanner.Err()
	}
	return numers, nil
}
