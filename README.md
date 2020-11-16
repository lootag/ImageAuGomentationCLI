# ImageAuGomentationCLI
ImageAuGomentationCLI is a simple command-line tool that allows users to conveniently analyze and augment annotated and not-annotated images. It's written in Go (Golang), and leverages its beautiful concurrency model to speed up computations.

### Setup
DISCLAIMER: this utility has only been tested on GNU/Linux and macOS. If you're running Windows, it's strongly recommended that you install Windows Subsystem for Linux.
Below are the steps to get the CLI.
#### 1) Get the Go SDK
On Debian-based systems: 
```
$> sudo apt-get install golang
```
On macOs:
```
$> brew install go
```
If you're running some other GNU/Linux distribution check out https://golang.org/doc/install.
#### 2) Clone this repository and setup environment vars
Enter the folder where you intend to clone the code and run 
```sh
$> git clone https://github.com/lootag/ImageAuGomentationCLI.git 
```

Then:
```sh
$> export GOPATH=${HOME}/go
$> export GOBIN=${GOPATH}bin
$> export PATH=${PATH}:${GOBIN}
$> go env
[...]
```

#### 3) Make the utility
Prepare test environment:
```sh
$> make build_test
```

Run test:
```sh
$> make test
```

### 4) Install

Option 1: Install augoment in /usr/local/bin. 
Fix _build_ target in Makefile, in order to set your preferred location based on your $PATH env
```sh
$> make build
```

Option 2: Install augoment in $GOBIN
```sh
$> make install
```

At this point the utility should have been installed successfully. Try out your installation by running ```$> augoment -h```. If you get a list with the utility's command-line arguments, the process was successful.

### Usage
The augoment command needs to be run in a directory that has two subdirectories:
1) ```Images```, containing all your images;
2) ```Annotations```, containing all your annotations.

By default, the command will run in the current directory, that is ```.```. If you wish to run it in a directory other than ```.``` you'll need to specify a value for the ```-folder``` argument. 
#### 1) Scan your dataset
The first thing you might want to do is get a picture of your dataset. You can do this by running 
```$> augoment -scan -folder="/path/to/folder/"```. 
This will print a list of all the classes in your dataset, with their corresponding number of instances. 
It's important to note that at the moment the utility only supports PASCAL_VOC. However, implementing readers and writers for your own custom format is pretty straightforward.
1) Add an annotation type in entities/AnnotationType.go;
2) In the annotationReaders folder, create a structure for your custom reader;
3) Implement the ```Read``` method for the structure (check out PascalVocReader.go for an example);
4) Implement the ```ReadSync``` method for the structure (check out PascalVocReader.go for an example);
5) Add the structure to annotationReaders/AnnotationReadersFactory.go;
6) In the annotationWriters folder, create a structure for your custom writer;
7) Implement the ```Write``` method for the structure (check out PascalVocWriter.go for an example);
5) Add the structure to annotationWriters/AnnotationWritersFactory.go;
6) Now add your command-line argument to converStringToAnnotationType.go.

Now you can use the utility with your custom annotation format by simply specifying the arguments ```-in_annotationtype``` and ```-out_annotationtype```.

#### 2) Augment your dataset
The first thing to understand, is that augoment splits up the data into batches, and processes them in parallel. In order to specify a batch size, you need to assign a value to the ```-batch_size``` command-line argument (the default value is 50).  
As of right now, the utility will allow you to: 
1) rotate your images 90 degrees left and right;
2) rotate your images 180 degrees;
3) blur your images.

It's important to note that all images will also be resized. You can specify the size through the ```-size``` command-line-argument (height and width will be the equal). The default value is 464. 
These actions are controlled by the  ```-rotate``` and ```-blur``` command-line arguments. If you don't specify any value for these arguments,  all augmentations will be performed.


By setting the ```-exlusion_threshold``` argument, you can exclude from your augmented dataset the images which contain classes whose number of instances is less than a certain threshold. 


If you want to manually exclude some classes from the augmented dataset, you can specify the ```-user_defined_exclusions``` argument (Ex. ```-  user_defined_exclusions="class1;class2;"```).


By default, the utility will assume that your dataset is annotated, and will therefore augment your annotations too. If you only wish to augment your images, you can simply set ```-annotations=false```.


