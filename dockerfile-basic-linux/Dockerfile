FROM nginx:1.14.2

RUN apt-get update

RUN apt-get install tcpdump -y

RUN apt install iputils-ping -y

RUN apt-get install netcat -y

RUN apt-get install net-tools -y

RUN apt-get install wget -y

RUN wget https://aka.ms/downloadazcopy-v10-linux

RUN tar -xvf downloadazcopy-v10-linux

RUN mv azcopy_linux_amd64_10.13.0/azcopy /usr/local/bin

RUN rm -rf downloadazcopy-v10-linux

RUN apt-get install -y procps

ENTRYPOINT ["sleep", "infinity"]