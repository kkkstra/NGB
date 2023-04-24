package param

type ReqAddPost struct {
	Title      string `json:"title" binding:"required,min=1,max=128"`
	Content    string `json:"content" binding:"required,min=1"`
	CategoryID uint   `json:"category_id" binding:"required"`
}
