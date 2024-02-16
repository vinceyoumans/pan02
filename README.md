# Pan02 challenge


Develop a Go application that continuously monitors a specified directory, tracking the creation and
modification of files. Implement the following functionalities:



1. File Monitoring: Implement a mechanism to continuously watch a target directory for the
   creation or modification of files.

2. Data Processing: Upon detection of a new or updated file, read its contents to determine the
   total number of bytes.



3. Data Storage: Construct a JSON object where each key-value pair consists of the file path and its
   corresponding byte count. Persist this JSON object to a storage location, the path of which
   should be configurable.


4. Concurrency Management: Process files concurrently. The degree of concurrency (i.e., the
   number of files processed in parallel) must be configurable.
VY:   I did not understand Q4.  I added a parameter for concurrency Management, but is not implemented.


5. Configuration Flexibility: Allow configuration settings (such as the target directory path, storage
   location path, and concurrency level) to be specified either through command-line arguments or
   via a configuration file.


# Non Functional Requirement:
    Adhere to Go coding best practices, including but not limited to project structure, naming
   conventions, error handling, and the effective use of interfaces.
    Ensure the code is well-structured, readable, and maintainable.



#  DELIVERABLE

From root,

```
make panmon
```


From root
```
panmon02
```
then start to add and update files.



## NOTES
I believe everything requested is delivered except... 
- Item 4.  I am not sure how to resolve.
- If an file is renamed, the original name is not removed from the JSON file.
  - I know how to fix this, but ran out of time.
  

From the ./tmp file you can ...
create file
edit file
delete file

You can also rename a file, but the original file name will remain in the JSON file.
I can fix this but ran out of time.


