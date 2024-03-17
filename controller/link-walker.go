package controller

import (
	"encoding/json"
	"link-walker-go/types"
	"log"
	"net/http"
)

func StartTask(task *types.Task) {
	resource := make(chan types.Resource)
	go fetchResource(task.URL, task.GetToken(), resource)

	<- resource
	println("Finished fetching resource")
	println(resource)
	// Count unique requests that needs to be made

	// walk each link
}

func fetchResource(url string, token string, resource chan<- types.Resource) {
    defer close(resource)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Printf("Error creating request for URL %s: %v", url, err)
        return
    }

    req.Header.Set("Authorization", "Bearer "+token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("Error fetching URL %s: %v", url, err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode == 200 {
        var resourceResponse types.Resource

        if err := json.NewDecoder(resp.Body).Decode(&resourceResponse); err != nil {
            log.Printf("Error decoding JSON from URL %s: %v", url, err)
            return
        }

        resource <- resourceResponse
    } else {
        log.Printf("Received non-200 status code %d from URL %s", resp.StatusCode, url)
    }
}

