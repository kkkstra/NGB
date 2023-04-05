package param

type ReqSignUp struct {
	Username string `json:"username" binding:"required,min=5,max=32"`
	Email    string `json:"email" binding:"required,max=128,email"`
	Password string `json:"password" binding:"required,min=6,max=64"`
	Intro    string `json:"intro" binding:"max=512"`
	Github   string `json:"github" binding:"max=39"`
	School   string `json:"school" binding:"max=32"`
	Website  string `json:"website" binding:"max=128,url"`
}

type ReqSignIn struct {
	Username string `json:"username" binding:"required,min=5,max=32"`
	Password string `json:"password" binding:"required,min=6,max=64"`
}

type ReqEditProfile struct {
	Intro   string `json:"intro,omitempty" binding:"max=512"`
	Github  string `json:"github,omitempty" binding:"max=39"`
	School  string `json:"school,omitempty" binding:"max=32"`
	Website string `json:"website,omitempty" binding:"max=128,url"`
}

type ReqEditPassword struct {
	OldPassword string `json:"old-password" binding:"required,max=64"`
	NewPassword string `json:"new-password" binding:"required,max=64"`
}

type ReqEditEmail struct {
	Email string `json:"email" binding:"required,max=128,email"`
}
