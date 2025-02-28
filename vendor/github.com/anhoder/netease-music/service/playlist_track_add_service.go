package service

import (
	"encoding/json"
	"net/http"

	"github.com/anhoder/netease-music/util"
)

type PlaylistTrackAddService struct {
	Id      string   `json:"id" form:"id"`
	SongIds []string `json:"songIds" form:"songIds"`
}

func (service *PlaylistTrackAddService) AddTracks() (float64, []byte) {
	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}
	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	data["id"] = service.Id

	tracks := make([]map[string]interface{}, 0, len(service.SongIds))
	for _, id := range service.SongIds {
		tracks = append(tracks, map[string]interface{}{
			"type": 3,
			"id":   id,
		})
	}
	if d, err := json.Marshal(tracks); err == nil {
		data["tracks"] = string(d)
	}

	code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/playlist/track/add`, data, options)

	return code, reBody
}
