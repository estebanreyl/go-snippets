FROM mcr.microsoft.com/oss/go/microsoft/golang:1.23-fips-azurelinux3.0

RUN tdnf update -y && \
    tdnf install ca-certificates-microsoft -y && \
    tdnf install golang -y

RUN tdnf update -y && \
    tdnf install ca-certificates-microsoft -y && \
    tdnf install shadow-utils -y && \
    tdnf install net-tools -y && \
    tdnf install tcpdump -y && \
    tdnf install libxcrypt-devel -y && \
    tdnf install awk -y