package handlers

import (
	"context"
)

func (s *LinksServer) GetFullLink(ctx context.Context, req *ShortLinkRequest) (*GetFullLinkResponse, error) {
	link := req.ShortLink
	s.logger.Infof("Recieved: %s", link)
	fullLink := s.repo.GetFullByShort(link)
	if fullLink == "" {
		s.logger.Infoln("Invalid short link!")
		return &GetFullLinkResponse{
			Result: &GetFullLinkResponse_Message{
				Message: &MessageResponse{Message: "Invalid short link!"},
			},
		}, nil
	}
	return &GetFullLinkResponse{
		Result: &GetFullLinkResponse_FullLink{
			FullLink: &FullLinkResponse{FullLink: fullLink},
		},
	}, nil
}
