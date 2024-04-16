package driver

import "log"

type ConnectionDetail struct {
	Hostname string
	Username string
	Password string
	Database string
	ConId    int
}

type Driver interface {
	Open() (Conn string, err error)
	Execute(conn int, statement string) (result []string, err error)
	Close(conn int)
}

func (c *ConnectionDetail) Open() (Conn string, err error) {
	log.Println("connection details converted to connection string")
	Conn = c.Username + ":" + c.Password + "@" + c.Hostname + "/" + c.Database
	c.ConId = 1
	return Conn, nil
}

func (c *ConnectionDetail) Execute(conn int, statement string) (result []string, err error) {
	log.Println("execute the SQL", statement)
	result = []string{"I", "tell", "you"}
	return result, nil
}

func (c *ConnectionDetail) Close(conn int) {
	log.Println("close the connection", conn)
	c = new(ConnectionDetail)
	log.Println("closed connection", c)
}
