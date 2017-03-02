#!/bin/bash
install_system_info() {
    if [ -d $HOME/system_info_project ]; then 
        GOPATH=$HOME/system_info_project/system_info/
        GOBIN=$GOPATH/bin
        
        if [ $(go install $GOPATH/src/system_info) ]; then 
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

if [ $(which go 2&> /dev/null) -eq 0 ]; then
    install_system_info
else
    if [ $(uname -a | grep -qi ubuntu) -eq 0]; then
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
        echo "\t$(uname -a)"
        exit 1
    fi
fi
