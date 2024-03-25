NAME=fizzbuzz

all: $(NAME)

$(NAME):
	docker-compose -f docker-compose.yml -f docker-compose-prod.yml up -d --build

dev:
	docker-compose up -d --build
	make logs

logs:
	docker-compose logs -f api

ps:
	docker-compose ps

re: fclean dev

fclean: clean
	docker image rm fizzbuzz-api-img

clean:
	docker-compose down