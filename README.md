# Bcrypt.fun

A simple website to generate Bcrypt hashes from strings.

See the project live at https://bcrypt.fun

## Run for development

Use `go-bindata` as described below and then type `go run main.go app.go api.go bindata.go` and go to `localhost:8005` in your browser

## Run for production

Use go-bindata to include the template files in your binary

```
go-bindata assets/... views/...
```

Then build

```
go build
```

And then simply run using `./bcrypt_fun` and go to `localhost:8005` in your browser

The following arguments are available


```
-cookiename - The name to be used for cookies
-siteurl - The site url
-sitename - The name of the site
-host - The host and port (localhost:8005 or :80)
-sessionexpiry - Time in seconds that sessions should last (3600 * 24 * 365)
```

## Contributors

Created by [Markus Tenghamn](https://ma.rkus.io)

Other contributors welcome!

