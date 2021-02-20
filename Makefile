.PHONY: clean

build:
	docker-compose run go ./scripts/build.sh

clean:
	rm -rf ./bin/*

deploy: clean build
	docker-compose run sls sls deploy --verbose
