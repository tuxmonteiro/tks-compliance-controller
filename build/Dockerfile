FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/tks-compliance-controller \
    USER_UID=1001 \
    USER_NAME=tks-compliance-controller

# install operator binary
COPY build/_output/bin/tks-compliance-controller ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
