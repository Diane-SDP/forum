package models

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
)

func Upload(img multipart.File, header *multipart.FileHeader) string {

	tempFile, err := ioutil.TempFile("./Assets/image/", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(img)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	print(tempFile.Name())

	return tempFile.Name()
}
