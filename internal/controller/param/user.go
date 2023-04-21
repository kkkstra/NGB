package param

type ReqSignUp struct {
	Username string `json:"username" binding:"required,min=5,max=32"`
	Email    string `json:"email" binding:"required,max=128,email"`
	Password string `json:"password" binding:"required,min=6,max=64"`
	Intro    string `json:"intro" binding:"max=512"`
}

type ReqSignIn struct {
	Method   string `json:"method" binding:"required"`
	Username string `json:"username" binding:"required,min=5,max=32"`
	Password string `json:"password" binding:""`
	Code     string `json:"code" binding:""`
}

type ReqGetSignInCode struct {
	Email string `json:"email" binding:"required,email"`
}

type ReqEditProfile struct {
	Username string `json:"username" binding:"required,min=5,max=32"`
	Intro    string `json:"intro" binding:"max=512"`
}

type ReqEditPassword struct {
	OldPassword string `json:"old-password" binding:"required,max=64"`
	NewPassword string `json:"new-password" binding:"required,max=64"`
}

type ReqEditEmail struct {
	Email string `json:"email" binding:"required,max=128,email"`
}

type ReqAddFollowing struct {
	Username string `json:"username" binding:"required,min=5,max=32"`
}

type ReqDeleteFollowing struct {
	Username string `json:"username" binding:"required,min=5,max=32"`
}
