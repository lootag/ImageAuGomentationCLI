#!/bin/bash
read -p "Enter your machine's password: " PASSWORD
read -p "Enter your machine's kernel(Linux/Darwin): " KERNEL
which brew
if [ $? != 0 ] && [ $KERNEL == "Darwin" ]; then
    # Install Homebrew
    ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
fi

which go    
# Install Go
if [ $? != 0 ] && [ $KERNEL == "Darwin" ]; then
    brew install golang
    touch $HOME/.bashrc
    sed -i -e '$a\export GOPATH=$HOME/go' $HOME/.bashrc
    sed -i -e '$a\export GOROOT="$(brew --prefix golang)/libexec"' $HOME/.bashrc
    sed -i -e '$a\export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin' $HOME/.bashrc
    source $HOME/.bashrc
else if [ $? != 0 ] && [ $KERNEL == "Linux" ]; then
    echo $PASSWORD | sudo -S apt-get install golang
    fi
fi

go build
echo $PASSWORD | sudo -S mv ImageAuGomentationCLI /usr/bin/augoment
