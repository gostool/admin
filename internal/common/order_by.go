package common

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
)

func CheckOrderByKey(set *gset.StrSet, m *gdb.Model, orderKey string, desc bool) (n *gdb.Model, err error) {
	if (set != nil) && (!set.Contains(orderKey)) {
		err = errors.New(fmt.Sprintf("orderKey must in %v", set.Slice()))
		return m, err
	}
	if desc {
		m = m.OrderDesc(orderKey)
	} else {
		m = m.OrderAsc(orderKey)
	}
	return m, nil
}
