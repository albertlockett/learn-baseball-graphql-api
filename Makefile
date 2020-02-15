default: build

clean:
	rm learn-baseball-graphql-api

build:
	go build
	docker build . -t albertlockett2/learnbaseball-graphql-api:latest

push:
	docker push albertlockett2/learnbaseball-graphql-api:latest