FRONT_END_BINARY = frontApp
BACK_END_BINARY = backApp

build_front:
	@echo "Building the front-end..."
	cd ../front-end && env GOOS=linux CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} .
	#cd ../front-end && docker build -f front-end.dockerfile -t longcat123/muscoach-front-end:1.0.10 . && docker push longcat123/muscoach-front-end:1.0.10 && cd ../project
	@echo "Front-end built."

build_backend:
	@echo "Building the back-end..."
	cd ../back-end && env GOOS=linux CGO_ENABLED=0 go build -o ${BACK_END_BINARY} .
	#cd ../back-end && docker build -f back-end.dockerfile -t longcat123/muscoach-back-end:1.0.2 . && docker push longcat123/muscoach-back-end:1.0.2 && cd ../project
	@echo "Back-end built."
