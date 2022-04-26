FROM gitpod/workspace-full

RUN sudo apt-get update && \
    sudo apt-get install gettext -y

USER gitpod

RUN brew tap suborbital/subo && \
    brew install subo && \
    brew install httpie && \
    brew install superfly/tap/flyctl

 