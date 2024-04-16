package models

type Teacher struct {
	Id          int
	Fullname    string
	Birthdate   string
	Subject     string
	Category    string
	Phone       string
	Img         string
	IsImgPublic bool
}

type User struct {
	Id          int
	Username    string
	Password    string
	Role        string
	Status      bool
	LastLoginAt string
}

type News struct {
	Id             int
	Title          string
	Overview       string
	Body           string
	IsNews         bool
	CreatedAt      string
	CreatedAtMilli int64
	CreatedBy      string
	Accepted       bool
}

type Pagination struct {
	Index int
	Value int
}

type NewsPage struct {
	TitleName  string
	Path       string
	Page       int
	NewsArray  []News
	Pagination []Pagination
}

type TeachersPage struct {
	Page       int
	Teachers   []Teacher
	Pagination []Pagination
}

type NewsCreate struct {
	TitleName string
	Path      string
}

type NewsUpdate struct {
	TitleName string
	Path      string
	Title     string
	Overview  string
	Body      string
}

type MainPage struct {
	News               []News
	NewsPagination     []Pagination
	Article            []News
	ArticlePagination  []Pagination
	Teachers           []Teacher
	TeachersPagination []Pagination
}

