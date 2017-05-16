#!/bin/sh

CONF_FILE=/app/config.yaml

if [ ! -z "$HOST" ];then
    echo "Replacing host with $HOST"
    sed -i "s/host:.*/host: '$HOST'/" $CONF_FILE
fi

if [ ! -z "$PORT" ];then
    echo "Replacing port with $PORT"
    sed -i "s/port:.*/port: $PORT/" $CONF_FILE
fi

if [ ! -z "$SHOW_URL" ];then
    echo "Replacing show_url with $SHOW_URL"
    sed -i "s/show_url:.*/show_url: '$SHOW_URL'/" $CONF_FILE
fi

if [ ! -z "$RESET_URL" ];then
    echo "Replacing reset_url with $RESET_URL"
    sed -i "s/reset_url:.*/reset_url: '$RESET_URL'/" $CONF_FILE
fi


/app/api_recorder -config $CONF_FILE
