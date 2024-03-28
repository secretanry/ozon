package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"ozontest/api/payload/responses"
	"ozontest/database/repositories"
)

func ShortLinkHandler(w http.ResponseWriter, r *http.Request, linkRepo repositories.Repo, logger *logrus.Logger) {
	shortLink := r.URL.Query().Get("link")
	logger.Infof("Recieved: %s", shortLink)
	fullLink := linkRepo.GetFullByShort(shortLink)
	if fullLink == "" {
		logger.Infoln("Invalid short link!")
		w.WriteHeader(400)
		jsonResponse, err := json.Marshal(&responses.RespModel{Message: "Invalid short link"})
		if err != nil {
			logger.Infoln("Cannot marshall json")
			logger.Debugln(err)
			return
		}
		_, err = w.Write(jsonResponse)
		if err != nil {
			logger.Infoln("Cannot send a response")
			logger.Debugln(err)
			return
		}
		return
	}
	w.WriteHeader(200)
	jsonResponse, err := json.Marshal(&responses.FullLinkResp{FullLink: fullLink})
	if err != nil {
		logger.Infoln("Cannot marshall json")
		logger.Debugln(err)
		return
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Infoln("Cannot send a response")
		logger.Debugln(err)
		return
	}
	return
}
