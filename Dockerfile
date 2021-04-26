FROM alpine:latest

EXPOSE 8080
ADD ./exec_bin/server-alpine /server
CMD /server


