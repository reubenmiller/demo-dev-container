
NAME?="myapp"

start-minikube:
	nohup bash -c 'minikube status || minikube start &' > minikube.log 2>&1

build:
	rm -f dist/app
	mkdir -p dist
	CGO_ENABLED=0 go build -ldflags '-w -extldflags "-static"' -o dist/app main.go
	chmod +x dist/app

start:
	@go run main.go

package: build
	docker build -t app -f container/Dockerfile .

start-docker: package
	docker run --rm -p 9000:9000 app

start-docker-compose: package
	docker compose -f container/docker-compose.yaml up

package-helm: build
	minikube image build -t custom/app:1.16.0 -f container/Dockerfile .

start-helm: package-helm
	helm install $(NAME) myapp/ --values myapp/values.yaml --values myapp/override-local.yaml --set 'app.settings[0].value=$(NAME)'

remove-helm:
	helm uninstall $(NAME)
