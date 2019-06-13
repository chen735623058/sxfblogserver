package domain


type Anthology struct {
	Id string `bson:id`
	Name string `bson:"name"`
}
