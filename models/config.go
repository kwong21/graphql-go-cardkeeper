type Config struct {
	Server   server
	Database database
}

type database struct {
	Server   string
	User     string
	Password string
	Database string
	Port     string
}

type server struct {
	Mode string
}
