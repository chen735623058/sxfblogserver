package domain

type Article struct {
	Id string
	Title string
	Desc string
	Content  string
	Coverimage string
	Anthologyid string
	Tags []string
	Createtime string
	Updatetime string
	Readnumbers int
	State string
}