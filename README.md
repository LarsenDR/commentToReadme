# commentToReadme
---
This is a program will create a **Readme Markdown** file from the tagged comments and code in that source file.

by Dave Larsen KV0S

code at <github.com/kv0s/commentToReadme>

Licensed under the GPL3

This README.md generated by this program.  Current testing includes golang sources.

In the Release directory there are compiled executables for windows64, linux64, linux64arm (nano, odroid), linux32arm RPI

To run use this commandline on our computer

  > commentToReadme -filename sourcecode >README.md


---

     '''
	
	//MD # commentToReadme
	//MD ---
	//MD This is a program will create a **Readme Markdown** file from the tagged comments and code in that source file.
	//MD
	//MD by Dave Larsen KV0S
	//MD
	//MD code at <github.com/kv0s/commentToReadme>
	//MD
	//MD Licensed under the GPL3
	//MD
	//MD This README.md generated by this program.  Current testing includes golang sources.
	//MD
	//MD In the Release directory there are compiled executables for windows64, linux64, linux64arm (nano, odroid), linux32arm RPI
	//MD
	//MD To run use this commandline on our computer
	//MD
	//MD   > commentToReadme -filename sourcecode >README.md
	//MD
	//MD
	//MD ---
	
     '''

	This is a block comment 1
	more comment 2

