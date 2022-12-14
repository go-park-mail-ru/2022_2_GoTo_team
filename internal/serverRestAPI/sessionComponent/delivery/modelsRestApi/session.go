package modelsRestApi

type UserInfoBySession struct {
	Username      string `json:"username"`
	Login         string `json:"login"`
	AvatarImgPath string `json:"avatar_img_path"`
}

type SessionCreate struct {
	UserData UserData `json:"user_data"`
}

type UserData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
