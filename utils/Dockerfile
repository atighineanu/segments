# Dockerfile
FROM alpine:3.13
# Get source
RUN apk update --no-cache && apk upgrade --no-cache && apk add --no-cache ca-certificates go git
#RUN ls -alh
RUN git clone https://github.com/atighineanu/segments
RUN cd segments; /usr/bin/go build . 
RUN cp /segments/megaproj1 /usr/bin/

EXPOSE 80