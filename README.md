# go-vk-groups-loader

[![Build Status](https://travis-ci.com/vnkrtv/go-vk-groups-loader.svg?branch=master)](https://travis-ci.com/vnkrtv/go-vk-groups-loader)

### Description

Loads posts from vk groups into PostgreSQL DB.

### Installation

- Install app:
  - ```git clone https://github.com/vnkrtv/go-vk-groups-loader.git```
- Set list of vk groups screen names in config/groups.json. Example of config/groups.json:
  -  ```["meduzaproject", "ria", "kommersant_ru", "tj", "rbc"]```
- App settings (vk token and PostgreSQL connection information) stored in 'config/config.json' file. You can fill them yourself or by running 'configure_settings' script:
  - ```./deploy/configure_settings```
- Build docker image:
  - ```docker build -t groups-loader-service .```
- Run app as docker container (running PostgreSQL required):
  - ```docker run --name groups-loader groups-loader-service ```

