package dsp

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
)

func ReadSignalStringValuesFile(path string, sampleRate float64) (*Signal, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ReadSignalStringValues(file, sampleRate)
}

func ReadSignalBytes(dataBytes []byte, sampleRate float64) (*Signal, error) {
	reader := bytes.NewReader(dataBytes)
	return ReadSignalStringValues(reader, sampleRate)
}

func ReadSignalStringValues(reader io.Reader, sampleRate float64) (*Signal, error) {
	signal := Signal{
		SampleRate: sampleRate,
		Signal:     make([]float64, 0),
	}

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		v, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return nil, err
		}

		signal.Signal = append(signal.Signal, v)
	}
	return &signal, scanner.Err()
}
