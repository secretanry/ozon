package handlers

import (
	"context"
	"github.com/sirupsen/logrus"
	"math/rand"
	"ozontest/database/models"
	"ozontest/database/repositories"
)

func generateShortLink() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

type LinksServer struct {
	UnimplementedLinksServiceServer
	logger *logrus.Logger
	repo   repositories.Repo
}

func NewLinksServer(logger *logrus.Logger, repo repositories.Repo) *LinksServer {
	return &LinksServer{logger: logger, repo: repo}
}

func (s *LinksServer) PostFullLink(ctx context.Context, req *LongLinkRequest) (*PostFullLinkResponse, error) {
	link := req.LongLink

	s.logger.Infof("Received: %s", link)
	fromDb := s.repo.GetShortByFull(link)
	if fromDb != "" {
		return &PostFullLinkResponse{
			Result: &PostFullLinkResponse_ShortLink{
				ShortLink: &ShortLinkResponse{ShortLink: fromDb},
			},
		}, nil
	}

	var shortLink string
	c := 0
	for shortLink == "" {
		temp := generateShortLink()
		if !s.repo.ContainsShortLink(temp) {
			shortLink = temp
		}
		if c == 1000 {
			s.logger.Infoln("Request body is incorrect")
			s.logger.Debugln("Request body is incorrect")
			return &PostFullLinkResponse{
				Result: &PostFullLinkResponse_Message{
					Message: &MessageResponse{Message: "Cannot generate short link, please try again later"},
				},
			}, nil
		}
		c++
	}
	err := s.repo.AddLink(&models.Links{ShortLink: shortLink, FullLink: link})
	if err != nil {
		s.logger.Infoln("Cannot add new link")
		s.logger.Debugln(err)
		return &PostFullLinkResponse{
			Result: &PostFullLinkResponse_Message{
				Message: &MessageResponse{Message: "Cannot add new link"},
			},
		}, nil
	}
	return &PostFullLinkResponse{
		Result: &PostFullLinkResponse_ShortLink{
			ShortLink: &ShortLinkResponse{ShortLink: shortLink},
		},
	}, nil
}
