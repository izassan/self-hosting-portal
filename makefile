build_dev:
	@cp services_dev.json services.json
	@sudo docker build -t self-hosting-portal:dev .
	@rm services.json

build_prod:
	@cp services_prod.json services.json
	@sudo docker build -t self-hosting-portal:latest .
	@rm services.json

