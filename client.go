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


// parse raw messages from the socket
// message - bytes
func (c *client) handle(message []byte) {
  cmd := bytes.ToUpper(bytes.TrimSpace(bytes.split(message, []byte(" "))[0]))
  args := bytes.TrimSpace(bytes.TrimPrefix(message, cmd))

  switch string(cmd) {
  case "REG":
    if err := c.reg(args); err != nil {
      c.err(err)
    }
  case "JOIN":
    if err := c.join(args); err != nil {
      c.err(err)
    }
  case "LEAVE":
    if err := c.leave(args); err != nil {
      c.err(err)
    }
  case "MSG":
    if err := c.msg(args); err != nil {
      c.err(err)
    }
  case "CHNS":
    c.chns()
  case "USRS":
    c.usrs()
  default:
    c.err(fmt.Errorf("Invalid command %S", cmd))

  }

}
