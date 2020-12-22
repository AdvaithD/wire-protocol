package main

type client struct {
	conn       net.Conn
	outbound   chan<- command // send only channel with type command (client <-> hub)
	register   chan<- *client // send only channel with type client
	deregister chan<- *client // send only channel b/w hub and client
	username   string
}

func (c *client) read() error {
	for {
		msg, err := bufio.NewReader(c.conn).ReadBytes('\n')
		if err == io.EOF {
			// deregister client, connection closed
			c.deregister <- c
			return nil
		}
		if err != nil {
			return err
		}
		c.handle(msg)
	}

}
