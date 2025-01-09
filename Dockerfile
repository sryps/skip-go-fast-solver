FROM golang:1.23-bookworm AS build-env

ARG TAG=main
WORKDIR /root
RUN git clone -b $TAG https://github.com/skip-mev/skip-go-fast-solver.git
RUN cd skip-go-fast-solver && make build

RUN wget -P /lib https://github.com/CosmWasm/wasmvm/releases/download/v2.1.0/libwasmvm.x86_64.so

FROM debian:bookworm-slim
RUN apt-get update && \
    apt-get install -y ca-certificates sqlite3 && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN useradd -m skip_go_fast_solver -s /bin/bash
WORKDIR /home/skip_go_fast_solver
USER skip_go_fast_solver:skip_go_fast_solver

COPY --chown=0:0 --from=build-env /root/skip-go-fast-solver/build/skip_go_fast_solver /usr/local/bin/skip_go_fast_solver
COPY --chown=0:0 --from=build-env /lib/libwasmvm.x86_64.so /lib/libwasmvm.x86_64.so
COPY --chown=1000:1000 --from=build-env /root/skip-go-fast-solver/db/migrations /home/skip_go_fast_solver/data/migrations

# mount docker volume here
VOLUME [ "/home/skip_go_fast_solver/data" ]

# `-migrations-path` flag can be set to /home/skip_go_fast_solver/data/migrations
# `-sqlite-db-path` flag can be set to /home/skip_go_fast_solver/data/skip_go_fast.db
# `-keys` flag can be set to /home/skip_go_fast_solver/keys.json and mount keys here (outside of the mounted volume)
# `-config` flag can be set to /home/skip_go_fast_solver/config.json and mount config here (outside of the mounted volume)

ENTRYPOINT ["/usr/local/bin/skip_go_fast_solver"]
