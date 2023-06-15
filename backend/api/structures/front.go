package structures

type PostWithLike struct {
	Post
	Username      string
	Likes         int
	IsLikedByUser bool
}

type LikeSentByTheFront struct {
	IdPost int
	IdUser int
}

type UpdateUsername struct {
	IdUser      string
	NewUsername string
}
