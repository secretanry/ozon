package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"ozontest/api/payload/requests"
	"ozontest/api/payload/responses"
	"ozontest/database/models"
	"ozontest/database/repositories"
)

func writeErrorToResponse(w http.ResponseWriter, code int, message string, logger *logrus.Logger) {
	w.WriteHeader(code)
	jsonResponse, err := json.Marshal(&responses.RespModel{Message: message})
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
}

func generateShortLink() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func LongHandler(w http.ResponseWriter, r *http.Request, linkRepo repositories.Repo, logger *logrus.Logger) {
	var request requests.LongLinkRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Infoln("Request body is incorrect")
		logger.Debugln(err)
		writeErrorToResponse(w, 400, "Request body is incorrect", logger)
		return
	}
	link := request.Link
	logger.Infof("Received: %s", link)
	fromDb := linkRepo.GetShortByFull(link)
	if fromDb != "" {
		w.WriteHeader(200)
		jsonResponse, err := json.Marshal(&responses.ShortLinkResp{ShortLink: fromDb})
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
	var shortLink string
	c := 0
	for shortLink == "" {
		temp := generateShortLink()
		if !linkRepo.ContainsShortLink(temp) {
			shortLink = temp
		}
		if c == 1000 {
			logger.Infoln("Request body is incorrect")
			logger.Debugln("Request body is incorrect")
			writeErrorToResponse(w, 500, "Cannot generate short link, please try again later", logger)
			return
		}
		c++
	}
	err = linkRepo.AddLink(&models.Links{ShortLink: shortLink, FullLink: link})
	if err != nil {
		logger.Infoln("Cannot add new link")
		logger.Debugln(err)
		return
	}
	w.WriteHeader(200)
	jsonResponse, err := json.Marshal(&responses.ShortLinkResp{ShortLink: shortLink})
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
}
