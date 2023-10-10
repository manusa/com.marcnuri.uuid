FROM azul/zulu-openjdk-alpine:11-jre

MAINTAINER Marc Nuri <marc@marcnuri.com>
LABEL MAINTAINER="Marc Nuri <marc@marcnuri.com>"

EXPOSE 8080

COPY ./build/libs /opt

CMD java -jar /opt/uuid-0.0.1-SNAPSHOT.jar
