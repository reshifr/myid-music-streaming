FROM debian:bookworm

# Initial
RUN apt-get update
RUN apt-get install -y wget

COPY docker-entrypoint.sh /usr/bin/docker-entrypoint.sh
RUN chmod +x /usr/bin/docker-entrypoint.sh

# PostgreSQL client
RUN echo 'deb https://apt.postgresql.org/pub/repos/apt bookworm-pgdg main' > /etc/apt/sources.list.d/pgdg.list
RUN wget -q -O /etc/apt/trusted.gpg.d/postgresql.asc https://www.postgresql.org/media/keys/ACCC4CF8.asc
RUN apt-get update
RUN apt-get -y install postgresql-client

COPY postgresql/init.sh /opt/play/postgresql/init.sh
RUN chmod +x /opt/play/postgresql/init.sh

# Minio client


# End
ENTRYPOINT [ "docker-entrypoint.sh" ]
