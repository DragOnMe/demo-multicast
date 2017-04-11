FROM centos

WORKDIR /root

ADD multicast_linux_amd64 /root

EXPOSE 9999/udp

ENTRYPOINT /root/multicast_linux_amd64
