# Pixel Challenge
This repo contains my entry for the pixel challenge, as part of the Extended Golang Training. When provided with an image and a directory, the program will output a JSON containing a similarity score between 0 and 1.  

## Usage
To use the tool, navigate to the directory and run:
```bash
go run main.go [image-filename] [directory-relative-path]
 ```

 The image you wish to use as the basis of your comparison **must be in the same directory provided as the second argument**
 
 ## Output
 The similarity result will be output as a JSON file in the ```/tmp``` directory. It will have the same filename as the image used as the basis for comparison.

## Future Work
- More testing
- Funcitoning bencmark test for the entire process
- Performance improvements. Performance drops off a cliff when analysing large numbers of images. 