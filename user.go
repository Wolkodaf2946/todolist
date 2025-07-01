package todolist

//все сущности аналогично бд

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
