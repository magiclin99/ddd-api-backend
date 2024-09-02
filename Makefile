.PHONY: init

init:
	go install github.com/golang/mock/mockgen@v1.6.0

dockerRun:
	docker build -t ddd-api-backend .
	docker rm -f ddd-api-backend
	docker run -d --name ddd-api-backend -p 8080:8080 ddd-api-backend