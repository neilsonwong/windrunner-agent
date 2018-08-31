#!/bin/bash
# adds agent to osx user profile startup

# AGENT="$(pwd)/../bin/agent_osx"
# uninstall old one first
# sudo launchctl remove "WindrunnerAgent"
# sudo launchctl submit -l "WindrunnerAgent" -- $AGENT 

# ^^^ is old one, plist should be easier to use
# we want a launchagent

# copy into launchAgents
# use cp command as bash one might be aliased
/bin/cp -rf ./windrunnerAgent.plist ~/Library/LaunchAgents/windrunnerAgent.plist

#fix file permission
chmod 644 ~/Library/LaunchAgents/windrunnerAgent.plist
