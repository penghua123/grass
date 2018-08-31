#!/bin/bash

pg_status=$(psql -U postgres -c "SELECT u.datname  FROM pg_catalog.pg_database u where u.datname='usgmtr';")
echo $pg_status
pg_num=$(echo $pg_status |cut -c 27)
echo $pg_num
if [ $pg_num==1 ]; then
  echo "Successfully"
fi
