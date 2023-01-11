FROM node:16.19

RUN npm install -g nodemon

COPY --from=golang:1.19.5 /usr/local/go /usr/local/go

ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /home/node/app

CMD ["nodemon", "-e", "go,mod,sum", "--signal", "SIGTERM", "--exec", "go run ./cmd/going-live || exit 1"]
