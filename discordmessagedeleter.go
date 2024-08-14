package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

var token string

var channelID string

var userID string

type messages struct {
    ID      string `json:"id"`
    Content string `json:"content"`
    Author  struct {
        ID            string `json:"id"`
        Username      string `json:"username"`
        Discriminator string `json:"discriminator"`
        Avatar        string `json:"avatar"`
    } `json:"author"`
}

func main() {
    fmt.Print("Enter your token: ")

    fmt.Scanln(&token)

    fmt.Print("Enter target channelID: ")

    fmt.Scanln(&channelID)

    fmt.Print("Enter your userID: ")

    fmt.Scanln(&userID)

    url := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages", channelID)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    req.Header.Add("Authorization", token)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    var msgs []messages
    json.Unmarshal(body, &msgs)

    for _, msg := range msgs {
        if msg.Author.ID == userID {
            deleteURL := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages/%s", channelID, msg.ID)
            deleteReq, err := http.NewRequest("DELETE", deleteURL, nil)
            if err != nil {
                fmt.Println("Error creating delete request:", err)
                continue
            }

            deleteReq.Header.Add("Authorization", token)

            _, err = http.DefaultClient.Do(deleteReq)
            if err != nil {
                fmt.Println("Error deleting message:", err)
                continue
            }
            fmt.Println("Deleted message:", msg.ID)
        }
    }
}
