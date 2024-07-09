package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func send(input ChapiteauArguments, contentType string, body *bytes.Buffer, isFile bool) error {
	lastPart := "report"
	if isFile {
		lastPart = "file"
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/project/%s/%s", input.ServiceUrl, input.Project, lastPart), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Auth))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload file: %s", resp.Status)
	}

	return nil
}

func uploadIndexHtml(input ChapiteauArguments) error {
	file, err := os.Open(input.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	if file.Name() != "index.html" {
		return fmt.Errorf("file should be html file of playwright report")
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return err
	}

	if _, err := io.Copy(part, file); err != nil {
		return err
	}

	writer.WriteField("buildId", input.BuildId)
	writer.WriteField("buildName", input.BuildName)
	writer.WriteField("reportUrl", input.ReportURL)

	err = writer.Close()
	if err != nil {
		return err
	}

	if err := send(input, writer.FormDataContentType(), body, true); err != nil {
		return err
	}

	fmt.Println("Index.html uploaded successfully")
	return nil
}

func uploadReport(input ChapiteauArguments) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	hasIndexHtml := false

	err := filepath.Walk(input.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			if file.Name() == "index.html" {
				hasIndexHtml = true
			}

			part, err := writer.CreateFormFile("files", filepath.Join(filepath.Base(input.Path), info.Name()))
			if err != nil {
				return err
			}
			_, err = io.Copy(part, file)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if !hasIndexHtml {
		return fmt.Errorf("index.html file not found in the provided folder")
	}
	if err != nil {
		return err
	}

	writer.WriteField("buildId", input.BuildId)
	writer.WriteField("buildName", input.BuildName)

	err = writer.Close()
	if err != nil {
		return err
	}

	if err := send(input, writer.FormDataContentType(), body, false); err != nil {
		return err
	}

	fmt.Println("Playwright report uploaded successfully")
	return nil
}
