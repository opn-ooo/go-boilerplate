package database

type connection struct {
	host     string
	port     string
	dbname   string
	user     string
	password string
}

func (c *connection) filterPassword() {
	if c.password != "" {
		c.password = "[FILTERED]"
	}
}

func (c *connection) string() string {
	s := "host=" + c.host +
		" port=" + c.port +
		" dbname=" + c.dbname +
		" user=" + c.user +
		" sslmode=disable"

	if c.password != "" {
		s = s + " password=" + c.password
	}

	return s
}
