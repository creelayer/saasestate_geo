package core

type Application struct {
	Config map[string]string
	Pgx    *Pgx
	Gorm *Gorm
}

var App = Application{}
