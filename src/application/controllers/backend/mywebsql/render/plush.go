package render

import (
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"sync"

	"github.com/gobuffalo/plush"
	"github.com/xiusin/pine"
)

type Plush struct {
	cache map[string]*plush.Template
	l     sync.RWMutex
	dir   string
	debug bool
}

func New(dir string, debug bool) *Plush {
	t := &Plush{debug: debug, dir: dir}

	t.cache = map[string]*plush.Template{}
	return t
}

func (c *Plush) AddFunc(funcName string, funcEntry interface{}) {
	c.l.Lock()
	_ = plush.Helpers.Add(funcName, funcEntry)
	c.l.Unlock()
}

func (c *Plush) Ext() string {
	return ".php"
}

func (c *Plush) Exec(name string, binding pine.H) (data []byte, err error) {
	defer func() {
		if err := recover(); err != nil {
			err = fmt.Errorf("%s", err)
		}
	}()
	var (
		tpl *plush.Template
		ok  bool
	)
	c.l.RLock()
	tpl, ok = c.cache[name]
	c.l.RUnlock()
	if !ok || c.debug {
		c.l.Lock()
		defer c.l.Unlock()
		s, err := ioutil.ReadFile(filepath.Join(c.dir, name))
		if err != nil {
			return nil, err
		}

		tpl, err = plush.NewTemplate(string(s))
		if err != nil {
			return nil, err
		}
		c.cache[name] = tpl
	}
	ctx := plush.NewContext()
	for k, v := range binding {
		ctx.Set(k, v)
	}
	html, err := tpl.Exec(ctx)
	if err != nil {
		return nil, err
	}
	data = []byte(html)
	return data, err
}

func (c *Plush) HTML(writer io.Writer, name string, binding map[string]interface{}) error {
	if byts, err := c.Exec(name, pine.H(binding)); err != nil {
		return err
	} else {
		_, err = writer.Write(byts)
		return err
	}
}
