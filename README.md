Aports API
==========

[![Build Status](https://travis-ci.org/myhro/aports-api.svg?branch=master)](https://travis-ci.org/myhro/aports-api)

Alpine Linux package database API written in Go, inspired by [aports-turbo][aports-turbo].

## Database

Right now there's no way to import the Alpine Linux package data straight to PostgreSQL. The recommended approach is to use [PGLoader][pgloader] to migrate the aports-turbo SQLite database to PostgreSQL.


[aports-turbo]: https://github.com/alpinelinux/aports-turbo
[pgloader]: https://github.com/dimitri/pgloader
