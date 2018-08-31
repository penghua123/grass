#!/bin/bash
usage()
{
    filename=$(basename $0)
    echo "Usage:   $filename -c \"local\""
    echo "         $filename -h | --help"
    echo "         -c --client  : which client to build the environment"
    echo "         -h --help    : Help usage"
    exit 1
}
VERSION="0.0.1"
PARSE=`/usr/bin/getopt -q  'c' --long client "$@"`

if [[ $? != 0 ]] || [[ -z $2 ]]; then
    usage
fi

CLIENT=d
while [ -n "$1" ] ; do
    case "$1" in
        -c | --client) CLIENT=$2; shift 2 ;;
        -h | --help) usage;;
        --) shift; break ;;
        *) echo "Parameter error"; usage ;;
    esac
done


DIR=$(pwd)
SQL_FILE=$DIR"/usgmtr.sql"
 if [ $CLIENT == l ]; then
    pg_statue=$(psql -U postgres -c "SELECT u.datname  FROM pg_catalog.pg_database u where u.datname='usgmtr';")
    pg_num=$(echo $pg_status |cut -c 27)
    if [ $pg_num==1 ]; then
        pg_dump -U usgmtr -d usgmtr -f um-updates-$(date +'%m%d%H%M%S').tar
        psql -U postgres -c "drop database usgmtr;"
        echo "Remove database usgmtr."
    fi
    psql -U postgres -c "create database usgmtr;"
    psql -U usgmtr < $SQL_FILE
else
    IMG="postgres:9.6.8"
    CONTAINER=usgmtr

    POSTGRE_USER=usgmtr
    POSRGRES_PASSWORD=$PGPASSWORD
    POSTGRES_DB=usgmtr
    POSTGRES_HOST=localhost
    POSTGRES_PORT=5433

    if cid=$(docker ps -a|grep -o -E $CONTAINER); then
    echo "Docker container $cid is removing...."
    pg_dump -h localhost -p 5433 -U usgmtr -F t -d usgmtr  -f um-updates-$(date +'%m%d%H%M%S').tar
    docker rm -f $cid
    fi
    PG_CID=$(docker create -p $POSTGRES_PORT:5432  --name $CONTAINER \
                       -e POSTGRES_USER=$POSTGRE_USER \
                       -e POSTGRES_PASSWORD=$POSRGRES_PASSWORD \
                       $IMG)
    docker start $PG_CID
    echo "$LOG_PREFIX Container $CONTAINER has been built." 
    # Postgres container need extra 5 seconds to accomplish initialization.
    sleep 5
    docker cp $SQL_FILE $CONTAINER:/usr/share/
    docker exec -it $CONTAINER /bin/sh -c 'psql -U usgmtr < /usr/share/usgmtr.sql'
    docker exec -it $CONTAINER /bin/sh -c 'rm /usr/share/usgmtr.sql'
fi
echo "Import tables structure to database $POSTGRES_DB successfully." 