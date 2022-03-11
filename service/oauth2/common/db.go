/*
 * @Author: lwnmengjing
 * @Date: 2022/3/9 10:53
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/9 10:53
 */

package common

import (
	"gopkg.in/mgo.v2"
	"time"
)

var DB *mgo.Database

func MakeDB(url, name string, timeout int) error {
	session, err := mgo.DialWithTimeout(
		url,
		time.Second*time.Duration(timeout))
	if err != nil {
		return err
	}
	DB = session.DB(name)
	return nil
}
