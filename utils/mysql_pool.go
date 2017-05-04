/*
 * @Author: dongtao
 * @Date: 2017-04-21 17:55:04
 * @Desc: mysql连接池
 * @Last Modified by: dongtao
 * @Last Modified time: 2017-04-21 18:19:09
 */
package utils

import "go.app.dt.com/conf"

// mysql连接池
type DbStore struct {
	pool chan *Mysql
}

// 创建连接池
func NewDbStore() {
	store := new(DbStore)
	store.pool = make(chan *Mysql, 50)
	for i := 0; i < 50; i++ {
		db, err := Connect(conf.XMmall)
		checkErr(err)
		store.pool <- db
	}
}

// 获取连接池
func (d *DbStore) GetConn() (db *Mysql) {
	db = <-d.pool
	return
}

// 返回连接池
func (d *DbStore) RetConn(db *Mysql) {
	d.pool <- db
	return
}
