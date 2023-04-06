FROM ubuntu:22.04

RUN apt update && apt install --no-install-recommends -y librsvg2-bin imagemagick icnsutils
RUN mkdir /workspace

ARG binary_location=target/ui-graphics-tools

COPY ${binary_location} /usr/bin
COPY ./scripts/launcher.sh /usr/bin/launcher
RUN chmod 755 /usr/bin/ui-graphics-tools && chmod 755 /usr/bin/launcher

WORKDIR /workspace
ENTRYPOINT [ "/usr/bin/launcher" ]