#!/bin/bash

export VCAP_SERVICES='{"p-redis":[{"name":"product-info","label":"p-redis","tags":["pivotal","redis"],"plan":"shared-vm","credentials":{"host":"localhost","password":"","port":6379}}]}'
export sell_margin="1.3"
export whole_sell_margin="1.4"
