FROM golang:1.8
RUN useradd -ms /bin/bash dockeruser
USER dockeruser
RUN mkdir -p /go/src/github.com/

EXPOSE 8080

CMD ["bash"]
