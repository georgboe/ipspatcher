package main

import "os"

func readNextBytes(file *os.File, number int) ([]byte, error) {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func getIntValue(arr []byte) int {
	value := 0
	for i, b := range arr {
		shiftValue := uint(len(arr) - 1 - i)
		value += int(b) << (8 * shiftValue)
	}
	return value
}

func readIntValue(file *os.File, number int) (int, error) {
	value, err := readNextBytes(file, number)
	if err != nil {
		return 0, err
	}
	return getIntValue(value), nil
}

func getByteArrayWithValue(length int, value byte) []byte {
	replacement := make([]byte, length)
	for i := range replacement {
		replacement[i] = value
	}
	return replacement
}
