# Gofio

Command line version of [Fily]("https://github.com/hfogelberg/fily") to quickly reduce a jpg image to a web friendly size.

## Installation
````
$ go build
$ go install gofio
````

## Usage
Basic usage is just to type `$ gofio -i my-file.jpg`. The file will be reduced to a default width of 700 px and the output file will be named with a timestamp.

## Flags
-i: Name of input file. Obligatory<br>
-o: Name of output file. Default is timestamp<br>
-w: Width of output file in pixels. Default is 700.<br>
-r: Remove original file. Default is false.<br>
-h: Show help<br>

So, the command `$ gofio -i big.jpg -o small.jpg -w 1200 -r` will create a 1200 px wide copy of the file big.jpg, name it small.jpg and remove the file big.jpg. 

-Get it?<br>
-Got it!<br>
-Good!<br>

## Improvements?
Please contribute!

## Todo
- Also handle png files.
