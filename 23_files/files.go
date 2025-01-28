package main

import (
	"fmt"
	"os"
)

func main() {

	// // Reading File Info
	// file, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo, err := file.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("File Name:", fileInfo.Name())
	// fmt.Println("Is Folder:", fileInfo.IsDir())
	// fmt.Println("File Size:", fileInfo.Size())
	// fmt.Println("Last Modified At:", fileInfo.ModTime())

	// Read File - Manually
	// file, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close() // closing the file

	// buffer := make([]byte, 12) // ideally instead of 12 file size should be passed

	// _, err = file.Read(buffer)

	// if err != nil {
	// 	panic(err)
	// }

	// for i := range buffer {
	// 	fmt.Println(string(buffer[i]))

	// }

	// Direct way to read file
	// data, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// READING FOLDER

	// dir, err := os.Open("../")

	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// info, err := dir.ReadDir(-1) // passing -1 to get all the file

	// for _, fi := range info {
	// 	fmt.Println(fi.Name())
	// }

	// CREATE A FILE
	// file, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// // Writing to file
	// file.WriteString("Hi GoLang ")
	// file.WriteString("Nice Language")

	// bytes := []byte("Hello GoLang")
	// file.Write(bytes)

	// Read from one file and write to another in streaming fashion

	// sourceFile, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example3.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	by, err := reader.ReadByte()
	// 	if err != nil {

	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(by)
	// 	if e != nil {
	// 		panic(e)
	// 	}

	// }

	// writer.Flush() // write to file whatever is present in the writer

	// fmt.Println("Written to new file successfully")

	// DELETE A FILE

	e := os.Remove("example3.txt")

	if e != nil {
		panic(e)
	}

	fmt.Println("File deleted successfully")

}
