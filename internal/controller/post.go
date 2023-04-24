package controller

import (
	"NGB/internal/controller/param"
	"NGB/internal/controller/response"
	"NGB/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	postID := c.Param("post_id")
	m := model.GetModel()
	post, err := m.GetPost(postID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get post. ", err.Error())
		return
	}
	user, err := m.FindUserById(strconv.Itoa(int(post.UserID)))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get the author of the post. ", err.Error())
		return
	}
	category, err := m.GetCategory(strconv.Itoa(int(post.CategoryID)))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get the category of the post. ", err.Error())
		return
	}
	postData := response.Data{
		"title":       post.Title,
		"content":     post.Content,
		"author_id":   post.UserID,
		"author":      user.Username,
		"category_id": post.CategoryID,
		"category":    category.Name,
	}
	response.Success(c, http.StatusOK, postData, "Get post successfully. ")
}

func AddPost(c *gin.Context) {
	t, _ := c.Get("userdata")
	userData := t.(map[string]string)
	userID, _ := strconv.Atoi(userData["id"])

	var req param.ReqAddPost
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
		return
	}

	m := model.GetModel()
	// check the category id
	_, err := m.GetCategory(strconv.Itoa(int(req.CategoryID)))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to find the category. ", err.Error())
		return
	}
	post := model.Post{
		Title:      req.Title,
		Content:    req.Content,
		UserID:     uint(userID),
		CategoryID: req.CategoryID,
	}
	postID, err := m.CreatePost(&post)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create post. ", err.Error())
		return
	}
	response.Success(c, http.StatusOK, response.Data{"post_id": postID}, "Create post successfully. ")
}

func DeletePost(c *gin.Context) {
	t, _ := c.Get("userdata")
	userData := t.(map[string]string)
	userID, _ := strconv.Atoi(userData["id"])
	role := userData["role"]

	postID := c.Param("post_id")
	m := model.GetModel()
	post, err := m.GetPost(postID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get post. ", err.Error())
		return
	}

	// check authorization
	if role != "admin" && uint(userID) != post.UserID {
		response.Error(c, http.StatusUnauthorized, "Insufficient permission. ")
		return
	}

	err = m.DelPost(strconv.Itoa(int(post.ID)))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete the post. ", err.Error())
		return
	}
	response.Success(c, http.StatusOK, response.Data{}, "Delete the post successfully. ")
}
