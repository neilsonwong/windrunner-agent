#!/bin/bash
# update windrunner on linux

# ------------------------------------------------------------------------------
# touch the updated file
# ------------------------------------------------------------------------------
touch updated

# ------------------------------------------------------------------------------
# set the windrunner root dir
# ------------------------------------------------------------------------------
WINDRUNNER_ROOT=$PWD

# ------------------------------------------------------------------------------
# set the update logging
# ------------------------------------------------------------------------------
LOGFILE=$WINDRUNNER_ROOT/update.log
echo "starting to update" >> $LOGFILE

# ------------------------------------------------------------------------------
# extract the zip in the update dir
# ------------------------------------------------------------------------------
echo "extracting windrunner update $WINDRUNNER_ROOT" >> $LOGFILE
unzip -o "$WINDRUNNER_ROOT/updates/*.zip" -d "$WINDRUNNER_ROOT/updates/"

# ------------------------------------------------------------------------------
# stop windrunner agent
# ------------------------------------------------------------------------------
#echo "stopping windunnner service" >> $LOGFILE
#systemctl --user stop windrunnerAgent.service

# ------------------------------------------------------------------------------
# copy binaries and configs into install dir
# ------------------------------------------------------------------------------
echo "remove windrunner agent" >> $LOGFILE
rm "$WINDRUNNER_ROOT/agent"
echo "backup config" >> $LOGFILE
mv "$WINDRUNNER_ROOT/config.json" "$WINDRUNNER_ROOT/config.json.old"
echo "copying new files" >> $LOGFILE
cp "$WINDRUNNER_ROOT/updates/agent" "$WINDRUNNER_ROOT/agent"
cp "$WINDRUNNER_ROOT/updates/config.json" "$WINDRUNNER_ROOT/config.json"
#merge is handled by windrunner internally

# ------------------------------------------------------------------------------
# clear out update dir
# ------------------------------------------------------------------------------
echo "clearing directory $WINDRUNNER_ROOT" >> $LOGFILE
rm -rf $WINDRUNNER_ROOT/updates/*

# ------------------------------------------------------------------------------
#start windrunner agent
# ------------------------------------------------------------------------------
echo "restarting windunnner service" >> $LOGFILE
systemctl --user restart windrunnerAgent.service
