# taylz.io/env

Package env provides runtime environment, using flags, file, and/or environment variables

## Environment config

```go
func main() {
	// most common case
	env := env.Default() // overwrite package name why not
	server := server.New(env) // clients use env
	...
	// multiple env #1
	serviceA := env.File(".private1.env") // new env with custom file name
	serviceB := env.File(".private2.env") // new env with custom file name
	handler1 := newXHandler(e1) // clients use env
	handler2 := newYHandler(e2) // clients use env
	server := server.New(handler1, handler2) // separate environments
	...
	// multiple env #2
	env := env.Default() // load global environment
	dbenv := env.Match("DB_") // extract env from context
	pay := env.Match("PAY_") // extract env from context
	handler1 := newXHandler(dbenv) // clients use env
	handler2 := newYHandler(pay) // clients use env
	server := server.New(handler1, handler2) // separate environments
}
```

## `cmd/dotenv`

Executable: print all values in the default environment

```sh
$ dotenv
env: open .env: The system cannot find the file specified.
env is empty
$ dotenv -help
env: open .env: The system cannot find the file specified.
help=true
$ touch .env
$ dotenv
env is empty
$ echo ENV=pro > .env
$ dotenv
ENV=pro
$ dotenv -a ENV=dev -flag=x flag=y
ENV=dev
a=true
flag=y
```
