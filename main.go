package main

import (
	"fmt"
	"io"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	patchFile = kingpin.Arg("patch", "The patch file").Required().ExistingFile()
	nesFile   = kingpin.Arg("rom", "The rom file (.nes)").Required().ExistingFile()
)

func main() {
	kingpin.Parse()
	patches := getPatches(*patchFile)
	cp(*nesFile, *nesFile+".orig")
	writePatchesToFile(patches, *nesFile)
}

func writePatchesToFile(patches []Patch, file string) {
	f, err := os.OpenFile(file, os.O_WRONLY, os.FileMode(0666))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, patch := range patches {
		_, err := f.WriteAt(patch.bytes, int64(patch.address))
		if err != nil {
			fmt.Println("Write failed for patch", patch)
			fmt.Println("Err: ", err)
			os.Exit(1)
		}
	}
}

func getPatches(file string) []Patch {

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	headerBytes, err := readNextBytes(f, 5)
	if err != nil {
		fmt.Println("Could not read patch header.")
		os.Exit(1)
	}
	header := string(headerBytes)

	if header != "PATCH" {
		fmt.Println("File is missing patch header.")
		os.Exit(1)
	}

	var patches []Patch

	for {
		address, _ := readIntValue(f, 3)
		noOfBytes, err := readIntValue(f, 2)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			fmt.Println("Error reading from file")
			os.Exit(1)
		}

		var replacement []byte
		if noOfBytes == 0 {
			noOfBytes, _ = readIntValue(f, 1)
			singleByteArr, _ := readNextBytes(f, 1)
			replacement = getByteArrayWithValue(noOfBytes, singleByteArr[0])

		} else {
			replacement, _ = readNextBytes(f, noOfBytes)
		}

		p := Patch{
			address: address,
			bytes:   replacement}

		patches = append(patches, p)
	}

	return patches
}

type Patch struct {
	address int
	bytes   []byte
}

func cp(src, dst string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	// no need to check errors on read only file, we already got everything
	// we need from the filesystem, so nothing can go wrong now.
	defer s.Close()
	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}
