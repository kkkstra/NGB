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

func GetAllThumbs(c *gin.Context) {
	postID := c.Param("post_id")
	m := model.GetModel()
	userIDs, err := m.GetThumbsByPost(postID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get thumbs.", err.Error())
		return
	}
	userDatas := []param.ResThumbsUser{}
	for _, userID := range userIDs {
		u, err := m.FindUserById(strconv.Itoa(int(userID)))
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to get thumbs.", err.Error())
			return
		}
		userData := param.ResThumbsUser{
			ID:       userID,
			Username: u.Username,
		}
		userDatas = append(userDatas, userData)
	}
	response.Success(c, http.StatusOK, response.Data{"thumbs": userDatas}, "Get all thumbs successfully. ")
}

func AddThumbs(c *gin.Context) {
	updateThumbs(c, false)
}

func DeleteThumbs(c *gin.Context) {
	updateThumbs(c, true)
}

func updateThumbs(c *gin.Context, delete bool) {
	t, _ := c.Get("userdata")
	userData := t.(map[string]string)
	userID, _ := strconv.Atoi(userData["id"])
	postID := c.Param("post_id")
	intPostID, _ := strconv.Atoi(postID)

	m := model.GetModel()

	if delete {
		err := m.DeleteThumbs(&model.Thumbs{
			UserID: uint(userID),
			PostID: uint(intPostID),
		})
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to delete thumbs. ", err.Error())
			return
		}
		response.Success(c, http.StatusOK, response.Data{}, "Delete thumbs successfully. ")
		return
	}

	// 检查是否已经点赞
	_, err := m.GetThumbsID(&model.Thumbs{
		UserID: uint(userID),
		PostID: uint(intPostID),
	})
	if err == nil {
		response.Error(c, http.StatusBadRequest, "Already liked the post. ")
		return
	}
	_, err = m.CreateThumbs(&model.Thumbs{
		UserID: uint(userID),
		PostID: uint(intPostID),
	})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to add thumbs. ", err.Error())
		return
	}
	response.Success(c, http.StatusOK, response.Data{}, "Add thumbs successfully. ")
}

func GetPostsByUser(c *gin.Context) {
	username := c.Param("username")

	m := model.GetModel()
	// get user id
	u, err := m.FindUserByUsername(username)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get the user. ", err.Error())
		return
	}
	posts, err := m.GetPostsByUser(strconv.Itoa(int(u.ID)))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get the posts. ", err.Error())
		return
	}
	resPosts := []param.ResPost{}
	for _, post := range posts {
		category, err := m.GetCategory(strconv.Itoa(int(post.CategoryID)))
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to get the category of the post. ", err.Error())
			return
		}
		resPost := param.ResPost{
			Title:      post.Title,
			Content:    post.Content,
			CategoryID: post.CategoryID,
			Category:   category.Name,
			UserID:     post.UserID,
			User:       u.Username,
		}
		resPosts = append(resPosts, resPost)
	}
	response.Success(c, http.StatusOK, response.Data{"posts": resPosts}, "Get all posts successfully. ")
}

func GetThumbsByUser(c *gin.Context) {
	username := c.Param("username")

	m := model.GetModel()
	// get user id
	u, err := m.FindUserByUsername(username)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get the user. ", err.Error())
		return
	}
	thumbs, err := m.GetThumbsByUser(strconv.Itoa(int(u.ID)))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get the thumbs. ", err.Error())
		return
	}
	resThumbs := []param.ResUserThumbs{}
	for _, id := range thumbs {
		post, err := m.GetPost(strconv.Itoa(int(id)))
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to get the post. ", err.Error())
			return
		}
		thumbs := param.ResUserThumbs{
			PostID: id,
			Title:  post.Title,
		}
		resThumbs = append(resThumbs, thumbs)
	}
	response.Success(c, http.StatusOK, response.Data{"posts": resThumbs}, "Get all thumbs successfully. ")
}

func GetPostsByCategory(c *gin.Context) {
}
