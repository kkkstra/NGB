package param

type ReqAddPost struct {
	Title      string `json:"title" binding:"required,min=1,max=128"`
	Content    string `json:"content" binding:"required,min=1"`
	CategoryID uint   `json:"category_id" binding:"required"`
}

type ResThumbsUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type ResPost struct {
	Title      string `json:"title" binding:"required,min=1,max=128"`
	Content    string `json:"content" binding:"required,min=1"`
	CategoryID uint   `json:"category_id" binding:"required"`
	Category   string `json:"category" binding:"required"`
	UserID     uint   `json:"user_id" binding:"required"`
	User       string `json:"user" binding:"required"`
}

type ResUserThumbs struct {
	PostID uint   `json:"post_id" binding:"required"`
	Title  string `json:"title" binding:"required,min=1,max=128"`
}
