// vessel-service/datastore.go
package main

import mgo "gopkg.in/mgo.v2"

// CreateSession creates the main session to our mongodb instance
func CreateSession(host string) (*mgo.Session, error) {
	dialInfo, err := mgo.ParseURL("mongodb://" + host)
	session, err := mgo.DialWithInfo(dialInfo)
	dialInfo.Direct = true
	dialInfo.FailFast = true

	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
