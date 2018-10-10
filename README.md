# Source

## Example Image

![Example Image](https://git.darknebu.la/GalaxySimulator/Source/raw/master/example_image_cropped.png)

## Building

Simple, just:
```
$ go build .

$ ./Source
21:52:10 [+] Opening the csv
21:52:10 [+] Calculate the acting forces                                 	
starting worker nr. 0, processing 6250 stars
starting worker nr. 1, processing 6250 stars
starting worker nr. 2, processing 6250 stars
starting worker nr. 3, processing 6250 stars
starting worker nr. 4, processing 6250 stars
starting worker nr. 5, processing 6250 stars
starting worker nr. 6, processing 6250 stars
starting worker nr. 7, processing 6250 stars

Stars:  50004 / 50004 [====================================================================] 100.00% 2m2s
Done Calculating the forces! Taking the rest of the session off!
22:09:44 [+] draw the slice and save it to out_2.png
...
```