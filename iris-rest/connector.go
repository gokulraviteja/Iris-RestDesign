package main

import (
	"fmt"

	"github.com/gocql/gocql"
	. "github.com/gocql/gocql"
	"gopkg.in/mgo.v2"
)

var booksTableConnector *mgo.Collection
var cassandraSession *gocql.Session
var dbconnected bool = dbconnect()

func dbconnect() bool {
	mongoInit()
	cassandraInit()
	return true
}

func mongoInit() {
	session, err := mgo.Dial(string("localhost"))
	if err != nil {
		fmt.Print("Error in setting up Mongo %%%% ", err)
	}

	booksTableConnector = session.DB(string("Store")).C("Books")
	fmt.Println("Mongo Connection set!!!!")
}
func cassandraInit() {
	cluster := NewCluster("127.0.0.1")
	cluster.Keyspace = string("store")
	cluster.Consistency = Consistency(1)
	cluster.Port = 29042

	fallbackSession, err := cluster.CreateSession()
	cassandraSession = fallbackSession
	if err != nil {
		fmt.Println("Cassandra connection problem", err)
		return
	}

	fmt.Println("Cassandra Connection set!!!!")
}
