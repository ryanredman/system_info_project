#!/bin/bash
if [ $(which go 2&> /dev/null) -eq 0 ]; then
    # We Go 

else
    # We no Go
    if [ $(uname -a | grep -qi ubuntu) -eq 0]; then
        echo "Adding repo: ppa:ubuntu-lxc/lxd-stable - $(date)"
        sudo add-apt-repository -y ppa:ubuntu-lxc/lxd-stable 
        echo "Uptating - $(date)"
        sudo apt-get -y update
        echo "Installing Go - $(date)"
        sudo apt-get -y install golang

        if [ $? -eq 0 ]; then

        else
            echo "There was a problem install Go. Exiting script."
            exit 1
        fi
    else
        echo "This script is designed to be ran on a machine running Ubuntu Linux."
        echo "You are running:"
        echo "\t$(uname -a)"
        exit 1
    fi
fi
