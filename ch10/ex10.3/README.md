Exercise 10.3: Using fetch http://gopl.io/ch1/helloworld?go-get=1, find out which service hosts the code samples for this book. (HTTP requests from go get include the go-get parameter so that servers can distinguish them from ordinary browser requests.)


```
go run . 'http://gopl.io/ch1/helloworld?go-get=1'
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io">
</head>
<body>
</body>
</html>

```

Answer: `https://github.com/adonovan/gopl.io`

