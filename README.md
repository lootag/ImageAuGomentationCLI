# ImageAuGomentationCLI
ImageAuGomentationCLI is a simple command-line tool that allows users to conveniently analyze and augment annotated and not-annotated images. It's writeen in Go (Golang), and leverages its beautiful concurrency model to speed up computations.

### Setup
DISCLAIMER: this utility has only been tested on GNU/Linux and macOS. If you're running Windows, it's strongly recommended that you install Windows Subsystem for Linux.
Below are the steps to get the CLI.
#### 1) Get the Go SKD
On Debian-based systems: 
```
$> sudo apt-get install golang
```
On macOs:
```
$> brew install go
```
If you're running some other GNU/Linux distribution check out https://golang.org/doc/install.
#### 2) Clone this repository
Simply go into the folder where you intend to clone your code and run ```$> git clone https://github.com/lootag/ImageAuGomentationCLI.git ```. 
#### 3) Build the utility
Now run ``` $> go build ```. This is going to create a binary called ImageAuGomentationCLI in the folder where you have your code.
#### 4) Put the utility in your /usr/local/bin folder
Simply run ``` $> sudo mv ImageAugomentationCLI /usr/local/bin/augoment ```. 
If you don't already have a ```.bashrc``` create it with ```$> touch ~/.bashrc```. 
Now run ```$> source ~/.bashrc```. 
At this point the utility should have been installed successfully. Try out your installation by running ```$> augoment -h```. If you get a list with the utility's command-line arguments, the process was successful.

### Usage
The augoment command needs to be run in a directory that has two subdirectories:
1) ```Images```, containing all your images.
2) ```Annotations```, containing all your annotations .
By default, the command will run in the current directory, that is ```.```. If you wish to run it in a directory other than ```.``` you'll need to specify a value for the ```-folder``` argument.
#### 1) Scan your dataset
The first thing you might want to do is get a picture of your dataset. You can do this running 
```$> augoment -scan -folder="/path/to/folder/```. 
This will print a list of all the classes in your dataset, with their corresponding number of instances.
