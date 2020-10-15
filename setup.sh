#!/bin/bash
read -p "Enter your machine's password: " PASSWORD
read -p "Enter your machine's kernel(Linux/Darwin): " KERNEL
which brew
if [ $? == 0 ] && [ $KERNEL == "Darwin" ]; then
    # Install Homebrew
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
fi

which go    
# Install Go
if [ $? == 0 ] && [ $KERNEL == "Darwin" ]; then
    echo $PASSWORD | sudo -S brew install golang
else if [ $? == 0 ] && [ $KERNEL == "Linux" ]; then
    echo $PASSWORD | sudo -S apt-get install golang
    fi
fi

go build
echo $PASSWORD | sudo -S mv ImageAuGomentationCLI /usr/bin/augoment
