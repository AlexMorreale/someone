# multi-stage docker build
# build stage
FROM golang:latest AS build-env
ADD . /src/someone
ENV GOPATH=/

# check for dep binary so we dont run this
# check file || curl ....
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN cd /src/someone \
    && dep ensure \
    && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o someone

# final stage
FROM alpine
WORKDIR /app
ADD ./tmpl ./tmpl
COPY --from=build-env /src/someone/someone /app/
RUN pwd && ls -la .
ENTRYPOINT ./someone
