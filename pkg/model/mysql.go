package model

type MySQL struct {
	Host            string `envconfig:"MYSQL_HOST" json:"host"`
	Port            string `envconfig:"MYSQL_PORT" json:"port"`
	User            string `envconfig:"MYSQL_USER" json:"user"`
	Password        string `envconfig:"MYSQL_PASSWORD" json:"password"`
	DBName          string `envconfig:"MYSQL_DBNAME" json:"dbname"`
	ConnMaxIdleTime int    `envconfig:"MYSQL_CONN_MAX_IDLE_TIME" json:"conn_max_idle_time"`
	ConnMaxLifeTime int    `envconfig:"MYSQL_CONN_MAX_LIFE_TIME" json:"conn_max_life_time"`
	DBMaxIdleConns  int    `envconfig:"MYSQL_MAX_IDLE_CONNS" json:"db_max_idle_conns"`
	DBMaxOpenConns  int    `envconfig:"MYSQL_MAX_OPEN_CONNS" json:"db_max_open_conns"`
}
