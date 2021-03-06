docker stop hello-go-web
docker rm hello-go-web

docker load -i 1.0.0.image

docker run -d -p 80:80 -p 8080:8080 --name hello-go-web --restart always kevisong/hello-go-web:1.0.0
docker logs -f hello-go-web