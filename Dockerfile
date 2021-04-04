FROM golang:1.16.3

LABEL "com.github.actions.name"="beetle-actions"
LABEL "com.github.actions.description"="Deploy to Kubernetes Cluster Using Beetle"
LABEL "com.github.actions.icon"="package"
LABEL "com.github.actions.color"="red"

LABEL "repository"="https://github.com/Clivern/beetle-actions"
LABEL "homepage"="http://github.com/clivern"
LABEL "maintainer"="Clivern <hello@clivern.com>"

ARG ACTION_VERSION=0.1.0
ARG BEETLE_VERSION=1.0.2

ENV GO111MODULE=on

RUN apt-get update

RUN mkdir -p /app

WORKDIR /app

RUN curl -sL https://github.com/Clivern/Beetle/releases/download/v${BEETLE_VERSION}/beetle_${BEETLE_VERSION}_Linux_x86_64.tar.gz | tar xz
RUN rm LICENSE
RUN rm README.md

RUN curl -sL https://github.com/Clivern/beetle-actions/releases/download/v${ACTION_VERSION}/beetle-actions_${ACTION_VERSION}_Linux_x86_64.tar.gz | tar xz
RUN rm LICENSE
RUN rm README.md

COPY entrypoint.sh /app

RUN chmod +x /app/entrypoint.sh
RUN chmod +x /app/beetle-actions
RUN chmod +x /app/beetle

ENTRYPOINT ["/app/entrypoint.sh"]