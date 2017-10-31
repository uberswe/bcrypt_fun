# Bcrypt.fun

Created by [Markus Tenghamn](https://ma.rkus.io)

The project uses the Revel framework.

See the project live at https://bcrypt.fun

Credit to this blog post for helping me get Revel running in Docker http://jbeckwith.com/2015/05/08/docker-revel-appengine/

## How to run this locally

Install docker

Clone this project

Then run the following commands in the project directory

`docker build -t bcrypt-fun .`

`docker run -it -p 8005:8005 bcrypt-fun`

The app should now be running on http://127.0.0.1:8080