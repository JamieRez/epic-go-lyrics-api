package main

import (
	"net/http"
	"github.com/labstack/echo"
  "encoding/json"
  "io/ioutil"
  "fmt"
)


type Song struct {
  Lyrics string `json:"lyrics"`
}

type IP struct {
  Ip string `json:"ip"`
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

func getExternalIP() IP {
  resp, err := http.Get("https://api.ipify.org?format=json")
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }

  extIp := IP{}
  json.Unmarshal(body, &extIp)
  return extIp
}

func getReqLocation(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    ip := getExternalIP()
    fmt.Println("You are at Make School:")
    fmt.Println(ip.Ip == "96.78.162.189")
    c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
    return next(c)
  }
}


func main() {
	e := echo.New()
  e.Use(getReqLocation)
	e.GET("/:artist/:song", func(c echo.Context) error {
    newSong := getSongLyrics(c.Param("artist"), c.Param("song"));
		return c.String(http.StatusOK, newSong.Lyrics)
	})
	e.Logger.Fatal(e.Start(":3000"))
}
