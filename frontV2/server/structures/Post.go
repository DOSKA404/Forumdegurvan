package structures

type Post struct {
	Username      string
	Content       string
	Likes         int
	IsLikedByUser bool
	Id_user       int
	Id_post       int
	Date_post     string
}
