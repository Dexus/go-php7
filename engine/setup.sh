#!/bin/bash

# Halt on errors
set -e

# Be verbose
set -x

function compile_curl() {
    pushd curl-7.54.1
    [[ -e "configure" ]] || ./buildconf
    ./configure --prefix /opt/curl --with-ssl=/opt/openssl
    make -j8
    make install
    #rm /opt/curl || true
    #ln -sf /opt/curl-7.54.1 /opt/curl
    popd
}

function compile_libmcrypt() {
    pushd libmcrypt-2.5.8
    [[ -e "configure" ]] || ./buildconf
    ./configure --prefix=/opt/libmcrypt --disable-posix-threads --enable-static
    make -j8
    make install
    #rm /opt/libmcrypt || true
    #ln -sf /opt/libmcrypt-2.5.8 /opt/libmcrypt
    #rm /opt/libmcrypt/lib/libmcrypt.so*
    #rm /opt/libmcrypt/lib/libmcrypt.la
    popd
}

function compile_zlib() {
    pushd zlib-1.2.11
    ./configure --prefix=/opt/zlib --static
    make -j8
    make install
    #rm /opt/zlib || true
    #ln -sf /opt/zlib-1.2.8 /opt/zlib
    popd
}

function compile_openssl() {
    pushd openssl-1.1.0f
    ./config --prefix=/opt/openssl --openssldir=/opt/openssl
    make -j8
    make install
    #rm /opt/openssl || true
    #ln -sf /opt/openssl-1.1.0f /opt/openssl
    popd
}

function compile_libxml2() {
    pushd libxml2-2.9.4
    ./autogen.sh --with-zlib=/opt/zlib --with-lzma --prefix=/opt/libxml2 --disable-shared --without-python
    make -j8
    make install
    #rm /opt/libxml2 || true
    #ln -sf /opt/libxml2-2.9.4 /opt/libxml2
    popd
}

[[ -e "/opt/openssl" ]] || compile_openssl
[[ -e "/opt/curl" ]] || compile_curl
[[ -e "/opt/libmcrypt" ]] || compile_libmcrypt
[[ -e "/opt/zlib" ]] || compile_zlib
[[ -e "/opt/libxml2" ]] || compile_libxml2
pushd php-7.1.7
[[ -e "configure" ]] || ./buildconf --force
./configure --prefix /opt/php7 --enable-embed=static --enable-debug \
    --enable-fpm --without-pear --without-iconv --enable-xml --with-libxml-dir=/opt/libxml2 --with-system-ciphers \
    --enable-sockets --enable-opcache --enable-redis --with-curl=/opt/curl --with-mysqli --with-mcrypt=/opt/libmcrypt \
    --enable-mbregex --enable-mbstring --enable-pdo --enable-bcmath --enable-pcntl --with-zlib=/opt/zlib \
    --with-openssl=/opt/openssl --enable-zip --enable-soap --enable-mysqlnd \
    --with-config-file-path=/opt/php7/etc
make 
make install
#rm /opt/php || true
#ln -sf /opt/php-7.1.7 /opt/php
popd
# opcache is not linked into libphp7.a, even we are doing static build
# rename to libxxx, so that we can link statically into disf-php
cp /opt/php7/lib/php/extensions/debug-non-zts-20160303/opcache.a /opt/php7/lib/php/extensions/debug-non-zts-20160303/libopcache.a
