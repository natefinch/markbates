# Mark Bates

This is an example of https://github.com/gobuffalo/packr/issues/11

First get this code:

```
go get -d github.com/natefinch/markbates
```

To reproduce, first the common case, which works:

from $GOPATH/src/github.com/natefinch/markbates 

run:
```
packr build
```

This will produce a markbates executable in the current directory, which when run will state `embedded: true`

now move two directories to $GOPATH/src/github.com/
 
run:
```
packr build github.com/natefinch/markbates
```

This will produce a markbates executable in the current directory, which when run will state `embedded: false`


