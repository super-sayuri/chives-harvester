package router

import "errors"

type CRouter struct {
	Name  string `json:"name" gorm:"primaryKey;not null"`
	Value string `json:"value" gorm:"not null"`
}

func (c *CRouter) SelfCheck(needKey bool) error {
	msg := ""
	if needKey {
		if len(c.Name) == 0 {
			msg = msg + "missing name, "
		}
	}
	if len(c.Value) == 0 {
		msg = msg + "missing value, "
	}

	if len(msg) != 0 {
		return errors.New(msg)
	}
	return nil
}
