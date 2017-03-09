# GOLANG

### Directory structure
```
bin/
    hello                          # command executable
    outyet                         # command executable
pkg/
    linux_amd64/
        github/GO/hello/
            stringutil.a           # package object
src/
    github/GO/
        .git/                      # Git repository metadata
	      hello/
	          hello.go               # command source
	      outyet/
	          main.go                # command source
	          main_test.go           # test source
	      stringutil/
	          reverse.go             # package source
	          reverse_test.go        # test source
```
### To build (This will create respective file in bin)
      
      $ cd $GOPATH/src/github/GO/hello
      $ go install
