package dbcache

type XormCacher struct {
	
}

func (x XormCacher) GetIds(tableName, sql string) interface{} {
	panic("implement me")
}

func (x XormCacher) GetBean(tableName string, id string) interface{} {
	panic("implement me")
}

func (x XormCacher) PutIds(tableName, sql string, ids interface{}) {
	panic("implement me")
}

func (x XormCacher) PutBean(tableName string, id string, obj interface{}) {
	panic("implement me")
}

func (x XormCacher) DelIds(tableName, sql string) {
	panic("implement me")
}

func (x XormCacher) DelBean(tableName string, id string) {
	panic("implement me")
}

func (x XormCacher) ClearIds(tableName string) {
	panic("implement me")
}

func (x XormCacher) ClearBeans(tableName string) {
	panic("implement me")
}
 
