#!/usr/bin/env bash
openssl req -x509 -nodes -newkey rsa:2048 -keyout tls.key -out tls.crt -days 3650
