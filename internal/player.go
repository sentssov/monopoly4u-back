package internal

type Player struct {
	Id           string `bson:"id"`
	Email        string `bson:"email"`
	Nickname     string `bson:"nickname"`
	PasswordHash []byte `bson:"password"`
}
