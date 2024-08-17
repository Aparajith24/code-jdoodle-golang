package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type JDoodleRequest struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Script       string `json:"script"`
	Stdin        string `json:"stdin"`
	Language     string `json:"language"`
	VersionIndex string `json:"versionIndex"`
	CompileOnly  bool   `json:"compileOnly"`
}

type JDoodleResponse struct {
	Output  string      `json:"output"`
	Status  interface{} `json:"statusCode"`
	Memory  string      `json:"memory"`
	CpuTime string      `json:"cpuTime"`
}

func readMultilineInput() string {
	reader := bufio.NewReader(os.Stdin)
	var lines []string
	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimRight(line, "\r\n")
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func mapLanguage(language string) string {
	switch language {
	case "python":
		return "python3"
	default:
		return language
	}
}

func main() {
	var code, input, language string
	fmt.Println("Enter the code:")
	code = readMultilineInput()
	fmt.Println("Enter the input (if any):")
	input = readMultilineInput()
	fmt.Println("Enter the programming language (e.g., python, java):")
	fmt.Scanln(&language)

	jdoodleReq := JDoodleRequest{
		ClientID:     "YOUR_KEY",
		ClientSecret: "YOUR_SECRET",
		Script:       code,
		Stdin:        input,
		Language:     mapLanguage(language),
		VersionIndex: "0",
		CompileOnly:  false,
	}
	log.Printf("Request: %s", jdoodleReq.Script)

	jdoodleReqBody, err := json.Marshal(jdoodleReq)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	resp, err := http.Post("https://api.jdoodle.com/v1/execute", "application/json", bytes.NewBuffer(jdoodleReqBody))
	if err != nil {
		log.Fatalf("Failed to communicate with JDoodle API: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read JDoodle API response: %v", err)
	}
	log.Printf("Response: %s", string(body))

	var jdoodleResp JDoodleResponse
	if err := json.Unmarshal(body, &jdoodleResp); err != nil {
		log.Fatalf("Failed to parse JDoodle API response: %v", err)
	}

	fmt.Println("Output:")
	fmt.Println(jdoodleResp.Output)
}
