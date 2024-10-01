package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type ResultData struct {
	Items []struct {
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			PublishedAt string `json:"publishedAt"`
		} `json:"snippet"`
	} `json:"items"`
}

const apiKey = "AIzaSyBcUiZBXAMMsPi6b4UWNmAlfWkbfr3BOkE"

func main() {

	ChannelById()
	// PlaylistById()
	// VideosById()
}

// https://youtube.googleapis.com/youtube/v3/channels?part=snippet%2CcontentDetails%2Cstatistics&id=UC_x5XG1OV2P6uZZ5FSM9Ttw&key=[YOUR_API_KEY]
//https://youtube.googleapis.com/youtube/v3/playlists?part=snippet%2CcontentDetails&channelId=UC_x5XG1OV2P6uZZ5FSM9Ttw&maxResults=26&key=[YOUR_API_KEY]

func ChannelById() {
	apiUrl := "https://youtube.googleapis.com/youtube/v3/channels?part=snippet%2CcontentDetails%2Cstatistics"

	if len(os.Args) < 2 {
		log.Println("Please Provide id")
		return
	}

	parameter := url.Values{}
	parameter.Add("key", apiKey)

	for i := 1; i < len(os.Args); i++ {
		parameter.Add("id", os.Args[i])
	}

	requestUrl := fmt.Sprintf("%s&%s", apiUrl, parameter.Encode())
	//	fmt.Println(requestUrl)

	resp, er := http.Get(requestUrl)

	if er != nil {
		log.Println(er)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Error! Status: ", resp.Status)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var result ResultData

	eR := json.Unmarshal(body, &result)
	if eR != nil {
		log.Println(eR)
		return
	}
	fmt.Println()
	for _, val := range result.Items {
		fmt.Println("Title: ", val.Snippet.Title)
		fmt.Println("Description: ", val.Snippet.Description)
		fmt.Println("Published at: ", val.Snippet.PublishedAt)
		fmt.Println()
	}

}

func PlaylistById() {
	var result ResultData
	apiUrl := "https://youtube.googleapis.com/youtube/v3/playlists?part=snippet%2CcontentDetails"
	if len(os.Args) < 2 {
		log.Println("Please Provide id")
		return
	}

	parameter := url.Values{}
	maxResult := "maxResults=25"
	parameter.Add("key", apiKey)
	for _, val := range os.Args[1:] {
		parameter.Add("id", val)
	}

	requestUrl := fmt.Sprintf("%s&%s&%s", apiUrl, maxResult, parameter.Encode())
	//fmt.Println(requestUrl)

	resp, er := http.Get(requestUrl)

	if er != nil {
		log.Println(er)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Error! Status: ", resp.Status)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	eR := json.Unmarshal(body, &result)
	if eR != nil {
		log.Println(eR)
		return
	}

	for _, val := range result.Items {
		fmt.Println("Title of playlists: ", val.Snippet.Title)
		fmt.Println("Description: ", val.Snippet.Description)
		fmt.Println("Published at: ", val.Snippet.PublishedAt)
		println()
	}
}

func VideosById() {
	var result ResultData
	apiUrl := "https://youtube.googleapis.com/youtube/v3/videos?part=snippet%2CcontentDetails%2Cstatistics"
	if len(os.Args) < 2 {
		log.Println("Please Provide id")
		return
	}

	parameter := url.Values{}
	parameter.Add("key", apiKey)

	for i := 1; i < len(os.Args); i++ {
		parameter.Add("id", os.Args[i])
	}

	requestUrl := fmt.Sprintf("%s&%s", apiUrl, parameter.Encode())
	//	fmt.Println(requestUrl)

	resp, er := http.Get(requestUrl)
	if er != nil {
		log.Println(er)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Error! Status: ", resp.Status)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	eR := json.Unmarshal(body, &result)
	if eR != nil {
		log.Println(eR)
		return
	}

	for _, val := range result.Items {
		fmt.Println("Title of playlists: ", val.Snippet.Title)
		fmt.Println("Description: ", val.Snippet.Description)
		fmt.Println("Published at: ", val.Snippet.PublishedAt)
		println()
	}
}
