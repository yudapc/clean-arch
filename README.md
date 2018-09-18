# clean-arch


![Build Status](https://travis-ci.com/yudapc/clean-arch.svg?branch=master)

Install glide

`$ brew install glide`

Install packages

`$ glide install`

Build Project

`$ go build`

Run Project

`$ ./clean-arch`

Please import file `.json` on folder `postman` to your postman app.

Test RESTFull using postman (`users` and `posts`). 


### Testing
Test RESTFull using `curl`:

#### Users

Get All Data Users:

`$ curl localhost:1323/users`

Create new data user:

```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"name":"Jhon Smith"}' \
  localhost:1323/users
```
Get one data user by id

`curl localhost:1323/users/1`

Update data user by id:

```
curl -X PUT \
  -H 'Content-Type: application/json' \
  -d '{"name":"[Update] Joe"}' \
  localhost:1323/users/1
```

Delete data user by id:

`curl -X DELETE localhost:1323/users/1`



#### Posts

Get All Data Posts:

`$ curl localhost:1323/posts`

Create new data post:

```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"title":"Motto Keseharian", "content": "Tiada Hari Tanpa Belajar dan harus kerja smart"}' \
  localhost:1323/posts
```
Get one data post by id

`curl localhost:1323/posts/1`

Update data post by id:

```
curl -X PUT \
  -H 'Content-Type: application/json' \
  -d '{"title":"[Update] Motto Keseharian", "content": "[Update] Tiada Hari Tanpa Belajar dan harus kerja smart"}' \
  localhost:1323/posts/1
```

Delete data post by id:

`curl -X DELETE localhost:1323/posts/1`
