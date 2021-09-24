build:
	docker build -t userwallet .
	docker run -t -i --env-file .env -d userwallet .
run:
	docker-compose up --build
down:
	 docker-compose down --remove-orphans --volumes

