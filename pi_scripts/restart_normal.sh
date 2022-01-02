#!/bin/bash
echo "killing Chromium"
ps -ef | grep chromium-browser | grep -v grep | awk '{print $2}' | xargs kill

echo "restarting....."
DISPLAY=:0 "/usr/bin/chromium-browser" --start-fullscreen --start-maximized https://videowall-kb9x9.ondigitalocean.app &
