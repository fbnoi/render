package render

import (
	"sync"
	"text/template"

	"github.com/pkg/errors"
)

var _cache = &cache{
	store: make(map[string]*template.Template),
	lock:  &sync.RWMutex{},
}

type cache struct {
	store map[string]*template.Template
	lock  *sync.RWMutex
}

func (c *cache) get(name string) (*template.Template, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if t, ok := c.store[name]; ok {
		return t, nil
	}

	return nil, errors.Errorf("template %s do not exist", name)
}

func (c *cache) set(name string, t *template.Template) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if t == nil {
		return errors.New("nil template pointer")
	}
	c.store[name] = t

	return nil
}

func Parse(path string) (t *template.Template, err error) {
	t = template.Must(template.New(path).Funcs(template.FuncMap{
		"extends": extends,
		"include": include,
	}).ParseFiles(path))

	err = _cache.set(path, t)

	return
}

func extends(path string) string {
	return path
	// return template.Must(template.New(path).Funcs(template.FuncMap{
	// 	"extends": extends,
	// 	"include": include,
	// }).ParseFiles(path))
}

func include(path string) string {
	return path
	// return template.Must(template.New(path).Funcs(template.FuncMap{
	// 	"extends": extends,
	// 	"include": include,
	// }).ParseFiles(path))
}
