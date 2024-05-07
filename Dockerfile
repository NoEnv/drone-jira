FROM golang:1.21 as builder

WORKDIR /go/src/github.com/noenv/drone-jira
ADD . /go/src/github.com/noenv/drone-jira

RUN go vet ./... \
  && go test -cover ./... \
  && CGO_ENABLED=0 go build -v -o release/plugin

FROM plugins/base:multiarch

LABEL maintainer="Lukas Prettenthaler <lukas@noenv.com>" \
  org.label-schema.name="Drone Jira Plugin" \
  org.label-schema.vendor="NoEnv" \
  org.label-schema.schema-version="1.0"

COPY --from=builder /go/src/github.com/noenv/drone-jira/release/plugin /bin/

ENTRYPOINT ["/bin/plugin"]
