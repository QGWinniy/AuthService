FROM postgres:16-alpine

USER root

RUN apk add --no-cache \
    python3 \
    py3-pip \
    py3-psycopg \
    curl \
    su-exec

RUN pip3 install --no-cache-dir --break-system-packages \
    'patroni[etcd3]'

# Каталог данных должен принадлежать postgres
RUN mkdir -p /var/lib/postgresql/data \
    && chown -R postgres:postgres /var/lib/postgresql

COPY init-user-servers-db.sh /usr/local/bin/init-user-servers-db.sh
RUN chmod +x /usr/local/bin/init-user-servers-db.sh

WORKDIR /var/lib/postgresql

USER postgres

CMD ["patroni", "/etc/patroni/patroni.yml"]
