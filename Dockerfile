FROM php:7.1
ENV VERSION=6.0.3
RUN apt-get update && \
    apt-get install -y git zip wget && \
    apt-get clean &&\
    rm -rf /var/lib/apt/lists/*
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" && \
    php composer-setup.php && \
    php -r "unlink('composer-setup.php');" && \
    mv composer.phar /usr/local/bin/composer
WORKDIR /tmp
RUN wget https://github.com/sensiolabs/security-checker/archive/v${VERSION}.zip && \
    unzip v${VERSION}.zip && \
    mv security-checker-${VERSION} /opt/security-checker
WORKDIR /opt/security-checker
RUN composer install
ARG WORK_DIR=/
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]