package webssh

import (
	"github.com/xiusin/pinecms/src/application/controllers/backend"
)

type SshController struct {
	backend.BaseController
}

func (c *SshController) Index() {
	c.Render().HTML("index.html")
}

func (c *SshController) GetLogin() {
	c.Render().HTML("index.html")
}

func (c *SshController) GetConsole()  {
	c.Render().HTML("console.html")
}

func (c *SshController) GetServers()  {
	c.Render().HTML("s_list.html")
}

func (c *SshController) GetAdd()  {
	c.Render().HTML("add.html")
}

func (c *SshController) GetSetpass()  {
	c.Render().HTML("reset.html")
}

func (c *SshController) GetOpenterm()  {
	c.Render().HTML("open_term.html")
}

func (c *SshController) GetTerm()  {
	c.Render().HTML("term.html")
}
