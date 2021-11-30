package main

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/wailsapp/wails"
)

type Customers struct {
	filename string
	runtime  *wails.Runtime
}

func NewCustomers() (*Customers, error) {
	result := &Customers{}
	return result, nil
}

func (c *Customers) LoadList() (string, error) {
	bytes, err := ioutil.ReadFile(c.filename)
	if err != nil {
		err = fmt.Errorf("Unable to open list: %s", c.filename)
	}
	return string(bytes), err
}

func (c *Customers) SaveList(customers string) error {
	return ioutil.WriteFile(c.filename, []byte(customers), 0600)
}

func (c *Customers) WailsInit(runtime *wails.Runtime) error {

	c.runtime = runtime
	homedir, err := c.runtime.FileSystem.HomeDir()
	if err != nil {
		return err
	}
	c.filename = path.Join(homedir, "mylist.json")

	return nil
}
