package main

import (
    "fmt"
    "log"
    "os"

    "github.com/ndyakov/go-lastfm"
)

func main() {
    apiKey := os.Getenv("LASTFM_API_KEY")
    apiSecret := os.Getenv("LASTFM_API_SECRET")
    user := os.Getenv("LASTFM_USER")

    api := lastfm.New(apiKey, apiSecret)
    response, err := api.User.GetRecentTracks(user, 0, 1, 0, 0)
    if err != nil {
        log.Fatal(err)
    }

    for _, track := range response.RecentTracks {
        if track.NowPlaying == "" {
            continue
        }
        fmt.Printf("%s - %s\n", track.Artist.Name, track.Name)
        break
    }
}
