package controller

import (
	"NGB/internal/controller/param"
	"NGB/internal/controller/response"
	"NGB/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	m := model.GetModel()
	res, err := m.GetAllCategories()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get categories. ", err.Error())
		return
	}
	categories := []param.ResCategory{}
	for _, category := range res {
		categories = append(categories, param.ResCategory{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	response.Success(c, http.StatusOK, response.Data{"categories": categories}, "Get all categories successfully. ")
}

func AddCategory(c *gin.Context) {
	t, _ := c.Get("userdata")
	userData := t.(map[string]string)

	// check authorization
	if userData["role"] != "admin" {
		response.Error(c, http.StatusUnauthorized, "Insufficient permission. ")
		return
	}

	var req param.ReqAddCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
		return
	}

	m := model.GetModel()
	category := model.Category{
		Name:        req.Name,
		Description: req.Description,
	}
	categoryID, err := m.CreateCategory(&category)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create category. ", err.Error())
		return
	}
	response.Success(c, http.StatusOK, response.Data{"category_id": categoryID}, "Create category successfully. ")
}

func GetCategory(c *gin.Context) {
	categoryID := c.Param("category_id")
	m := model.GetModel()
	category, err := m.GetCategory(categoryID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get category. ", err.Error())
		return
	}
	categoryData := response.Data{
		"id":          category.ID,
		"name":        category.Name,
		"description": category.Description,
	}
	response.Success(c, http.StatusOK, categoryData, "Get category successfully. ")
}

func DeleteCategory(c *gin.Context) {
	t, _ := c.Get("userdata")
	userData := t.(map[string]string)

	// check authorization
	if userData["role"] != "admin" {
		response.Error(c, http.StatusUnauthorized, "Insufficient permission. ")
		return
	}

	categoryID := c.Param("category_id")
	m := model.GetModel()
	_, err := m.GetCategory(categoryID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get category. ", err.Error())
		return
	}

	err = m.DelCategory(categoryID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete the category. ", err.Error())
		return
	}
	response.Success(c, http.StatusOK, response.Data{}, "Delete the category successfully. ")
}

func GetPostsFromCategory(c *gin.Context) {

}
