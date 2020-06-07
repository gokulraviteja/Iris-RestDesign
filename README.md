# Iris-RestDesign
Sample Rest Design with Iris a Framework of Go , With Cassandra and Mongo Connections


Requirements
1) Go
2) Cassandra
3) Mongodb

How to run:

1) Start Mongo, Cassandra(at 29042) setups locally
2) go build 
3) ./iris-rest

Apis:

/v1/getbooks (get)

/v1/createbook (post)

/v1/book/{bookname} (get)

/v1/book/{bookname} (put)

/v1/book/{bookname} (delete)
