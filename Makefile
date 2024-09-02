.PHONY: dockerRun

dockerRun:
	docker build -t ddd-api-backend .
	docker rm -f ddd-api-backend
	docker run -d --name ddd-api-backend -p 8080:8080 ddd-api-backend