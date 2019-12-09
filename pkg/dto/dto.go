package dto

import "github.com/wendellliu/good-search/pkg/common/dbAdapter"

type DTO interface {
	ExperienceDTO
}

type Repository struct {
	DB dbAdapter.Database
}
