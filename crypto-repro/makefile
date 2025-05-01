build:
	go build -o bin/crypto-repro

run-fips-scenario:
	docker build . -f Dockerfile-fips -t esrey/crypto-repro:v1 && docker run esrey/crypto-repro:v1

run-normal-scenario:
	docker build . -f Dockerfile-normal -t esrey/crypto-repro:v1 && docker run esrey/crypto-repro:v1