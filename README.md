# Bcrypt.fun

Created by [Markus Tenghamn](https://ma.rkus.io)

See the project live at https://bcrypt.fun

## Run for development

Simply type `go run main.go app.go api.go` and go to `localhost:8005` in your browser

## Run for production

Optionally use bindata if you want to include the template files in your binary

```
go-bindata assets/... views/...
```

And then simply run using `./bcrypt_fun` and go to `localhost:8005` in your browser


