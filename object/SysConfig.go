package object

type SysConfig struct {
	Total total `toml:"total"`
	Db    db    `toml:"db"`
}

type total struct {
	Port    int  `toml:"port"`
	IsDebug bool `toml:"isDebug"`
	MisType int  `toml:"misType"`
}

type db struct {
	Server string `toml:"server"`
	Port   int    `toml:"port"`
	DbName string `toml:"dbName"`
	Uid    string `tom;:"uid"`
	Pwd    string `toml:"pwd"`
}
