package repository

var Repo MysqlRepository

type MySqlRepositoryRepo struct{}

func MySqlInit() {
	Repo = &MySqlRepositoryRepo{}
}
