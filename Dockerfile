FROM alpine

RUN apk --no-cache add build-base git curl jq bash
RUN curl -s -k https://api.github.com/repos/JamesWoolfenden/pike/releases/latest | jq '.assets[] | select(.name | contains("linux_386")) | select(.content_type | contains("gzip")) | .browser_download_url' -r | awk '{print "curl -L -k " $0 " -o /usr/bin/pike.tar.gz"}' | sh
RUN tar -xf /usr/bin/pike.tar.gz -C /usr/bin/ && rm /usr/bin/pike.tar.gz && chmod +x /usr/bin/pike && echo 'alias pike="/usr/bin/pike"' >> ~/.bashrc
COPY entrypoint.sh /entrypoint.sh

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/entrypoint.sh"]
