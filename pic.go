package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	//"time"
)

/*
type Items struct {
	Fact   string
	Length string
}

	func GetCatFact() {
		apiUrl := "https://catfact.ninja/fact"

		reqst, err := http.NewRequest("GET", apiUrl, nil)

		if err != nil {
			fmt.Print("err: ", err)
		}

		client := &http.Client{}
		response, error := client.Do(reqst)

		if error != nil {
			fmt.Println(error)
		}

		responseBody, error := io.ReadAll(response.Body)

		if error != nil {
			fmt.Println(error)
		}

		//formattedData := formatJSON(responseBody)
		fmt.Println("Status: ", response.Status)
		fmt.Println("Response body: ", string(responseBody))

		Data := []byte(responseBody)

		var data1 Items

		errr := json.Unmarshal(Data, &data1)

		if errr != nil {
			fmt.Println(err)
		}

		fmt.Println("Stuct is: ", data1)
		fmt.Println("Fact is: ", data1.Fact)
		// clean up memory after execution
		defer response.Body.Close()
	}
*/

type JsonItems struct {
	CopyRight      string
	Date           string
	Explanation    string
	Hdurl          string
	MediaType      string
	ServiceVersion string
	Title          string
	Url            string
}

var Data1 JsonItems

func Geturl() string {
	var apiURL = "https://api.nasa.gov/planetary/apod?api_key=***"

	reqst, err := http.NewRequest("GET", apiURL, nil)

	if err != nil {
		fmt.Print("err: ", err)
	}

	client := &http.Client{}
	response, error := client.Do(reqst)

	if error != nil {
		fmt.Println(error)
	}

	responseBody, erroR := io.ReadAll(response.Body)

	if erroR != nil {
		fmt.Println(erroR)
	}

	//formattedData := formatJSON(responseBody)
	//fmt.Println("Status: ", response.Status)
	//fmt.Println("Response body: ", string(responseBody))

	Data := []byte(responseBody)

	errr := json.Unmarshal(Data, &Data1)

	if errr != nil {
		fmt.Println(err)
	}

	// clean up memory after execution
	defer response.Body.Close()
	return Data1.Url
}

func DownLoadPic() {
	url := Geturl()
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	datePath := "pic/" + "nasa" + Data1.Date + ".jpg"

	//open a file for writing
	file, err := os.Create(datePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Success!")
}

//todo Add a setings so you can pick wither or not to go and download the high-res pic or to copy the lower-res pic

func SavePic() {
	datePath := "pic/" + "nasa" + Data1.Date + ".jpg"

	savedPic := Data1.Date + ".jpg"
	if !fs.ValidPath(datePath) {
		DownLoadPic()
	}

	file, err := os.Create("savedPics/" + savedPic)
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()

	srcPath, e := os.Open(datePath)
	if e != nil {
		log.Fatal(e)
	}

	_, err = io.Copy(file, srcPath)
	if err != nil {
		log.Fatal(err)

	}

}
func CleanUp() {
	err := os.RemoveAll("pic/")
	if err != nil {
		log.Fatal(err)
	}
}
