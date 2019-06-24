package configure

type Configure struct {
	App    App
	Server Server
	MySQL  MySQL
	Redis  Redis
}

type App struct {
	Title   string
	Group   string
	Sys     string
	Inst    string
	Type    string
	Env     string
	Desc    string
	Version string
}
type Server struct {
	Host           string
	Port           string
	ReadTimeout    string
	WriteTimeout   string
	maxHeaderBytes string
}
type MySQL struct {
	Host     string
	Port     string
	Database string
	UserName string
	Password string
}
type Redis struct {
	Host     string
	Port     string
	Database string
	Auth     string
}
