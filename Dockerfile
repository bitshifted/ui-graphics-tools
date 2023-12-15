FROM ubuntu:22.04

ARG BUILD_DATE
ARG VCS_REF
ARG BUILD_VERSION

# Labels.
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.build-date=$BUILD_DATE
LABEL org.label-schema.name="bitshifted/ui-graphics-tools"
LABEL org.label-schema.description="Generates icons and splash screens from SVG"
LABEL org.label-schema.url="https://github.com/bitshifted/ui-graphics-tools"
LABEL org.label-schema.vcs-url="https://github.com/bitshifted/ui-graphics-tools"
LABEL org.label-schema.vcs-ref=$VCS_REF
LABEL org.label-schema.vendor="WSO2"
LABEL org.label-schema.version=$BUILD_VERSION

RUN apt update && apt install --no-install-recommends -y librsvg2-bin imagemagick icnsutils
RUN mkdir /workspace

ARG binary_location=target/ui-graphics-tools

COPY ${binary_location} /usr/bin
COPY ./scripts/launcher.sh /usr/bin/launcher
RUN chmod 755 /usr/bin/ui-graphics-tools && chmod 755 /usr/bin/launcher

WORKDIR /workspace
ENTRYPOINT [ "/usr/bin/launcher" ]
