go run main.go web
curl http://localhost:8000/ --output 1.gif
curl http://localhost:8000/ --output 2.gif
curl http://localhost:8000/ --output 3.gif
curl http://localhost:8000/ --output 4.gif
go run main.go > out.gif