package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type AppConfigProperties map[string]string

func ReadPropertiesFile(filename string) (AppConfigProperties, error) {
	config := AppConfigProperties{}

	if len(filename) == 0 {
		return config, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}

func main() {
	props, err := ReadPropertiesFile("properties/k8s_example.properties")
	service := os.Getenv("service")

	if err != nil && props["text"] != "" && service != "" {
		fmt.Println("Error while reading properties file")
		return
	}

	url := "http://" + service + "/?text=" + props["text"]
	for {
		resp, err := http.Get(url)
		fmt.Println(resp)
		fmt.Println(err)
		time.Sleep(5 * time.Second)
	}
}
