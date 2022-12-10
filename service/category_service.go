package service

import (
	"final_project_4/models"
	"final_project_4/repositories"
)

type CategoryService struct {
	rr repositories.CategoryRepoApi
}

func NewCategoryService(rr repositories.CategoryRepoApi) *CategoryService { //provie service
	return &CategoryService{rr: rr}
}

type CategoryServiceApi interface {
	CreateCategoryService(input models.Category) (models.Category, error)
	GetAllCategoryService() ([]models.Category, error)
	UpdateCategoryService(categoryId int, input models.Category) (models.Category, error)
	DeleteCategoryService(categoryId int) error
}

func (cs CategoryService) CreateCategoryService(input models.Category) (models.Category, error) {
	var category models.Category
	category.Type = input.Type

	category, err := cs.rr.CreateCategory(category)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (cs CategoryService) GetAllCategoryService() ([]models.Category, error) {
	var category []models.Category

	category, err := cs.rr.GetAllCategory(category)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (cs CategoryService) UpdateCategoryService(categoryId int, input models.Category) (models.Category, error) {
	// get category
	category, err := cs.rr.GetCategoryById(categoryId)
	if err != nil {
		return category, err
	}

	// Update user
	category, err = cs.rr.UpdateCategory(categoryId, input)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (cs CategoryService) DeleteCategoryService(categoryId int) error {
	// get category
	_, err := cs.rr.GetCategoryById(categoryId)
	if err != nil {
		return err
	}
	// delete category
	err = cs.rr.DeleteCategory(categoryId)
	if err != nil {
		return err
	}
	return nil

}
