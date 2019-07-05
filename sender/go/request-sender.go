package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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

func PrintFile(filename string, msg string, err string) {
	f, err1 := os.OpenFile("storage/"+filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err1 != nil {
		panic(err1)
	}

	defer f.Close()

	if _, err1 = f.WriteString(msg + " | " + err + "\n"); err1 != nil {
		panic(err1)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	log_filename := strconv.Itoa(rand.Int())
	props, err := ReadPropertiesFile("properties/k8s-example.properties")
	text := os.Getenv("text")

	if err != nil && props["service"] != "" && text != "" {
		fmt.Println("Error while reading properties file")
		return
	}

	url := "http://" + props["service"] + "/?text=" + text
	for {
		resp, err := http.Get(url)
		fmt.Println(resp)
		fmt.Println(err)
		if err == nil {
			PrintFile(log_filename, resp.Status, "")
		}
		time.Sleep(5 * time.Second)
	}
}
