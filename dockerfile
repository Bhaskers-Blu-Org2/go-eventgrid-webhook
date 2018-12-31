FROM alpine:latest

EXPOSE 8080

COPY eventgridwebhook .

CMD ./eventgridwebhook -logpath=/home/LogFiles/
