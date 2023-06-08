package structures

type User struct {
	Id_user     int
	Email       string
	Username    string
	DateOfBirth string
	Password    string
}

type Post struct {
	Id_post   int
	Date_post string
	Content   string
	Id_user   int
}

type Comment struct {
	Id_comment   int
	Content      string
	Date_content string
	Id_post      int
	Id_user      int
}

type Like struct {
	Id_Like int
	Id_post int
	Id_user int
}
