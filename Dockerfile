FROM alpine:latest

WORKDIR /app

COPY build/bff-gql .

RUN chmod +x bff-gql

EXPOSE 8000

CMD ["./bff-gql"]
