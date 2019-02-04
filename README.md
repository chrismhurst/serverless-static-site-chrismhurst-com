###Adding a function

1. Create a subdirectory of the name of the application in the root of the project
2. Make a main.go file there - this is where the function code will go
3. Add a 'env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go' line to the build steps in makefile
4. Add the scriptblock (see below) to the serverless.yml configuration
```
functions:
  hello:
    handler: bin/hello
    events:
      - http:
          path: hello
          method: get
```
5. 'make build'
6. 'sls deploy'
7. test (sls invoke -f hello)