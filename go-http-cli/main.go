package main

import (
	//"crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
	"io"
	"net/http"
	"sync"
)

func get(client *http.Client, filename string) error {
	url := fmt.Sprintf("https://localhost:10443/img/%s", filename)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
	/*
			outputFile, err := os.Create("img/" + filename)
			if err != nil {
				return err
			}
			_, err = io.Copy(outputFile, resp.Body)
		  if err != nil {
			  return err
		  }
		  return nil
	*/
}

func main() {
	tr := &http.Transport{}
	err := http2.ConfigureTransport(tr)
	if err != nil {
		fmt.Println("Error configuring HTTP/2 transport:", err)
		return
	}
	tr.TLSClientConfig.InsecureSkipVerify = true
	client := &http.Client{Transport: tr}

	wg := &sync.WaitGroup{}

	for i := 1; i < 20; i++ {
		wg.Add(1)
		filename := fmt.Sprintf("%d.webp", i)
		go func(filename string, client *http.Client) {
			defer wg.Done()
			err := get(client, filename)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%s downloadded successfully\n", filename)
		}(filename, client)
	}

	wg.Wait()
}
