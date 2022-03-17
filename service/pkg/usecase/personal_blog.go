package usecase

import (
	"github.com/wonpanu/personal-blog/service/pkg/entity"
	"github.com/wonpanu/personal-blog/service/pkg/repo"
)

type IBlogUsecase interface {
	Create(blog entity.Blog) (entity.Blog, error)
	GetAll() ([]entity.Blog, error)
	GetByBlogID(ID string) (entity.Blog, error)
	UpdateByBlogID(ID string, blog entity.Blog) (entity.Blog, error)
	DeleteByBlogID(ID string) error
}

type BlogUsecase struct {
	Repo repo.IBlogRepo
}

func (uc BlogUsecase) Create(blog entity.Blog) (entity.Blog, error) {
	response, err := uc.Repo.Create(blog)
	return response, err
}

func (uc BlogUsecase) GetAll() ([]entity.Blog, error) {
	return uc.Repo.GetAll()
}

func (uc BlogUsecase) GetByBlogID(ID string) (entity.Blog, error) {
	response, err := uc.Repo.GetByBlogID(ID)
	return response, err
}

func (uc BlogUsecase) UpdateByBlogID(ID string, blog entity.Blog) (entity.Blog, error) {
	response, err := uc.Repo.UpdateByBlogID(ID, blog)
	return response, err
}

func (uc BlogUsecase) DeleteByBlogID(ID string) error {
	return uc.Repo.DeleteByBlogID(ID)
}

func NewBlogUsecase(blogRepo repo.IBlogRepo) BlogUsecase {
	return BlogUsecase{
		Repo: blogRepo,
	}
}
