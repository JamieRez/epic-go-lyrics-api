package main

import (
	"net/http"
	"github.com/labstack/echo"
  "encoding/json"
  "io/ioutil"
)


type Song struct {
  Lyrics string `json:"lyrics"`
}

func getSongLyrics(artist string, title string) Song {
  resp, err := http.Get("https://api.lyrics.ovh/v1/" + artist + "/" + title)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }

  newSong := Song{}
  json.Unmarshal(body, &newSong)

  return newSong
}


func main() {
	e := echo.New()
	e.GET("/:artist/:song", func(c echo.Context) error {
    newSong := getSongLyrics(c.Param("artist"), c.Param("song"));
		return c.String(http.StatusOK, newSong.Lyrics)
	})
	e.Logger.Fatal(e.Start(":3000"))
}
