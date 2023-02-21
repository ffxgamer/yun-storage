FROM xhofe/alist:latest
LABEL MAINTAINER="i@nn.ci"
WORKDIR /opt/alist/
COPY config.json /opt/alist/data/
EXPOSE 5244

CMD [ "./alist", "server", "--no-prefix" ]
