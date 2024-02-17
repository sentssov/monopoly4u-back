package internal

type Player struct {
	Id       string `bson:"id"`
	Nickname string `bson:"nickname"`
	Password string `bson:"password"`
}
