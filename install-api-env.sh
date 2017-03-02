#!/bin/bash

install_system_info() {
    if [ -d /home/vagrant/system_info_project ]; then 
        GOPATH=/home/vagrant/system_info_project/system_info/
        GOBIN=$GOPATH/bin

        if [ ! -d $GOBIN ]; then
            mkdir $GOBIN
        fi

        go install system_info
        if [ $? -eq 0 ]; then 
            $GOBIN/system_info &
            echo "system_info running in background."
            exit 0
        else
            echo "Error install system_info. Exiting"
            exit 1
        fi
    else 
        echo "Git repo missing. Exiting."
        exit 1
    fi
}

which go
if [ $? -eq 0 ]; then
    install_system_info
else
    uname -a | grep -qi ubuntu

    if [ $? -eq 0 ]; then
        echo "Adding repo: ppa:ubuntu-lxc/lxd-stable - $(date)"
        sudo add-apt-repository -y ppa:ubuntu-lxc/lxd-stable 
        echo "Uptating - $(date)"
        sudo apt-get -y update
        echo "Installing Go - $(date)"
        sudo apt-get -y install golang

        if [ $? -eq 0 ]; then
            install_system_info
        else
            echo "There was a problem install Go. Exiting script."
            exit 1
        fi
    else
        echo "This script is designed to be ran on a machine running Ubuntu Linux."
        echo "You are running:"
        echo -e "\t$(uname -a)"
        exit 1
    fi
fi
