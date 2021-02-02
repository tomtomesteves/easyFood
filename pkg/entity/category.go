package entity

type Category struct {
	Id   int    `db:"id"`
	Name string `db:"nome"`
}
