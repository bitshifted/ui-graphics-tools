FROM ubuntu:22.04

RUN apt update && apt install -y librsvg2-bin imagemagick icnsutils
RUN mkdir /workspace

ARG binary_location=target/ui-graphics-tools

COPY ${binary_location} /usr/bin

WORKDIR /workspace
