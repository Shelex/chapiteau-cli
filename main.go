package main

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Define command-line flags
	htmlPath := flag.String("src", "playwright-report", "Path to test HTML report")
	// apiToken := flag.String("apiToken", "", "API token")
	// projectId := flag.String("project id", "", "Project id")
	// url := flag.String("url", "", "hosted playwright report html url")
	// buildId := flag.String("buildId", "", "CI Build ID")
	// buildUrl := flag.String("buildUrl", "", "CI Build URL")

	flag.Parse()

	// Read the HTML file

	if htmlPath == nil {
		fmt.Println("Missing report path")
		return
	}

	htmlContent, err := os.ReadFile(filepath.Join(*htmlPath, "index.html"))
	if err != nil {
		fmt.Println("Error reading HTML file:", err)
		return
	}

	// Extract the base64 string
	htmlString := string(htmlContent)
	start := strings.Index(htmlString, "window.playwrightReportBase64 = \"") + len("window.playwrightReportBase64 = \"")
	end := strings.Index(htmlString[start:], "\";") + start

	base64Header := "data:application/zip;base64,"
	base64String := strings.TrimPrefix(htmlString[start:end], base64Header)

	// Decode the base64 string into zip
	decodedData, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		return
	}

	// Unzip the decoded data
	zipReader, err := zip.NewReader(bytes.NewReader(decodedData), int64(len(decodedData)))
	if err != nil {
		fmt.Println("Error unzipping data:", err)
		return
	}

	// find report.json data in decoded zip archive
	var reportJson []byte
	for _, file := range zipReader.File {
		if file.Name != "report.json" {
			continue
		}
		rc, err := file.Open()
		if err != nil {
			fmt.Println("Error opening report.json:", err)
			return
		}
		reportJson, err = io.ReadAll(rc)
		rc.Close()
		if err != nil {
			fmt.Println("Error reading report.json:", err)
			return
		}
		break
	}

	if reportJson == nil {
		fmt.Println("report.json not found in the zip archive")
		return
	}

	log.Println(string(reportJson))

	// TODO parse []byte into proper Report struct
	// TODO
	// Prepare the POST request
	// postUrl := "https://chapiteau.shelex.dev/api/report"
	// req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(reportJson))
	// if err != nil {
	// 	fmt.Println("Error creating POST request:", err)
	// 	return
	// }
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", "Bearer "+*apiToken)

	// // Send the POST request
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("Error sending POST request:", err)
	// 	return
	// }
	// defer resp.Body.Close()

	// // Print the response
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response:", err)
	// 	return
	// }
	// fmt.Println("Response:", string(body))
}
