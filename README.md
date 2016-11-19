# api_example

To run commands, start the server:

```sh
go run api.go
```

## GET:

```sh 
go run get.go {id}
```
Where {id} is uuid, for example:

```sh 
go run get.go 4e10e0ee-3e4b-4c73-9de4-492592998445
```
Or use curl directly

```sh
curl localhost:8080/posts/{id}
```

## GET all:
```sh 
go run collect.go
```

Or use curl directly

```sh
curl localhost:8080/posts
```

## DELETE:
```sh 
go run delete.go {id}
```

Or use curl directly

```sh
curl -X "DELETE" localhost:8080/posts/{id}
```

## POST:

```sh 
go run post.go "{post}"
```

Where {post} must be in quotation marks, for example:

```sh 
go run post.go "This is a new post"
```

Or use curl directly

```sh
curl -H "Content-Type: application/json" -X "POST" -d {Body:post} localhost:8080/posts
```
Where {post} is in json format, for example:

```sh
curl -H "Content-Type: application/json" -X "POST" -d '{"Body":"This is a new post"}' localhost:8080/posts
```


# blog_example

```sh
go run goblog.go
```
runs the first template website

![alt tag](http://i.imgur.com/n5zhgXr.png)

```sh
go run goblog2.go
```
runs the second template

![alt tag](http://i.imgur.com/kUb4Q8M.png)
