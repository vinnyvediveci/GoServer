FROM golang
COPY Httpserver /opt/http-server-go/Httpserver
WORKDIR /opt/http-server-go/
ENTRYPOINT ["/opt/http-server-go/Httpserver"]