package mysqlsrc

type MysqlSource struct {
}

func (ms *MysqlSource) Read() ([]byte, error) {
	return nil, nil
}

func (ms *MysqlSource) Close() error {
	return nil
}
