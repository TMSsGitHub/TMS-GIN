package config

type Database struct {
	Host      string `yml:"host"`
	Port      int    `yml:"port"`
	User      string `yml:"user"`
	Password  string `yml:"password"`
	Table     string `yml:"table"`
	Charset   string `yml:"charset"`
	ParseTime bool   `yml:"parse_time"`
	Loc       string `yml:"loc"`

	Prefix        string `yml:"prefix"`
	SingularTable bool   `yml:"singular_table"`
}
