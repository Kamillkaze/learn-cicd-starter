FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

ADD notely /usr/bin/notely
ADD .env /usr/bin/.env

CMD ["sh", "-c", "env && notely"]
