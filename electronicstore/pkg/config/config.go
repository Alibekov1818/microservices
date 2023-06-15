package config

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
}

func GetConfig() Config {
	var myConfig Config

	myConfig.Db.Dsn = "postgres://estore_l4nw_user:eZ0QjO58OfiWD7n1RM1HMMEsRMKFESfL@dpg-ci3q10mnqqlbd9kjaneg-a.oregon-postgres.render.com/estore_l4nw"
	myConfig.Port = 8000
	myConfig.Env = "dev"
	myConfig.Db.MaxOpenConns = 15
	myConfig.Db.MaxIdleConns = 15
	myConfig.Db.MaxIdleTime = "15m"
	return myConfig
}
