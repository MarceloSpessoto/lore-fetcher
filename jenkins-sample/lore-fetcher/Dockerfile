FROM golang

WORKDIR /root/
RUN git clone https://github.com/MarceloSpessoto/lore-fetcher.git

WORKDIR /root/lore-fetcher
RUN ./install.sh

COPY lore-fetcher.toml /etc/lore-fetcher/lore-fetcher.toml

ENTRYPOINT [ "lore-fetcher",  "--fetch" ]
