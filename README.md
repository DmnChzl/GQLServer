# GraphQL Server

## Explanation

This is an example of modern API made with **Go** and **GraphQL** based on _book store_ idea.

### Mandatory

_Install **MongoDB** and run it before lauching this program._

```
mkdir db
mongo --dbpath db
```

## Process

Repository:

```
git clone https://github.com/mrdoomy/gqlserver.git
```

Dependencies:

```
go get -v github.com/graphql-go/graphql
go get -v github.com/graphql-go/handler
go get -v go.mongodb.org/mongo-driver/mongo
```

Build:

```
go build main.go
```

Enjoy =)

## Docs

- The **Go** Programming Language : [https://golang.org](https://golang.org)
- **GraphQL** : [https://github.com/graphql-go/graphql](https://github.com/graphql-go/graphql)
- **MongoDB** Go Driver : [https://github.com/mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver)

## License

```
"THE BEER-WARE LICENSE" (Revision 42):
<phk@FreeBSD.ORG> wrote this file. As long as you retain this notice you
can do whatever you want with this stuff. If we meet some day, and you think
this stuff is worth it, you can buy me a beer in return. Damien Chazoule
```
