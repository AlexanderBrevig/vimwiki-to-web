#!/bin/bash

# configure for your server
VIMWIKI=~/vimwiki
GMIOUT=~/gemini
HOST_USER=root
HOST_IP=10.10.10.10
HOST_GMI_PATH=/root/

# generate gemini files from your vimwiki folder
vwtw -vimwiki $VIMWIKI -file / -gmiout $GMIOUT

# keep in mind, the folder gemini will be copied
# so your export folder should be named the same
# as the target folder on your host machine
scp -r $GMIOUT $HOST_USER@$HOST_IP:$HOST_GMI_PATH
