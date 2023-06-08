package structures

type PostWithLike struct {
	Post
	Username string
	Likes    int
}

type LikeSentByTheFront struct {
	IdPost string
	IdUser string
}
