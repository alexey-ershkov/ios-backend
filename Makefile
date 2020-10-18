run:
	docker build -t proj . && docker run -p 5000:5000 proj