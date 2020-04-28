mkdir -p release
KO_DOCKER_REPO=mkmik ko resolve -f app.yaml >release/app.yaml
