package service

import (
	"api/pkg/blog"
)

type Service struct {
	repo blog.Repo
}

// New creates, inits and returns a Service struct
func New(repo blog.Repo) *Service {
	return &Service{repo: repo}
}
