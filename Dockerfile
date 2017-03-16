FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/iducate-services /opt/bin/iducate-services
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/iducate-services"]

