# go-vk-groups-loader

[![Build Status](https://travis-ci.com/vnkrtv/go-vk-groups-loader.svg?branch=master)](https://travis-ci.com/vnkrtv/go-vk-groups-loader)

### Description

Loads posts from vk groups into PostgreSQL DB.

### Installation

- Install app:
  - ```git clone https://github.com/vnkrtv/go-vk-groups-loader.git```
- Set list of vk groups screen names in config/groups.json. Example of config/groups.json:
  -  ```["meduzaproject", "ria", "kommersant_ru", "tj", "rbc"]```
- Run 'deploy_service' script for managing PostgreSQL configuration and running app instance:
  - ```./deploy/deploy_service```