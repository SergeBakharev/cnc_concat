# cnc_concat

This utility is designed to solve a specific problem encountered by users of the Autodesk Fusion 360 personal use license: the inability to combine multiple toolpaths into a single G-code (.cnc) file. The CNC File Concatenator automatically combines multiple .cnc files into a single file, ordering them accendingly based on the filename.

## Problem Overview
Autodesk Fusion 360, under its personal use license, restricts users from exporting multiple toolpaths within a single G-code file. This limitation can be cumbersome for users who require complex, multi-step or multi-tool machining processes. This utility addresses that problem by enabling users to concatenate multiple .cnc files into a single file that can be processed in a sequence determined by the filenames' alphabetical order.

## Features
- **Folder Concatenation**: Place all your `.cnc` files in a single folder and let the tool concatenate them into one seamless G-code file.
- **File Selection Concatenation**: Manually select multiple `.cnc` files, regardless of their directory, and combine them into one file.
- **Filename Order**: Files are concatenated in ascending alphabetical order based on their filenames to maintain a logical machining sequence.

## Usage
The easiest way to use this utility is through Windows File Explorer by using drag and drop functionality. Below are the detailed instructions:

### Installation
1. Download the latest release binary for your system or compile the source directly if you have Go installed.
2. If downloaded, unzip the contents to a preferred directory.

### Running the Tool
- **Using a Folder**:
  1. Drag and drop a folder containing `.cnc` files onto the compiled executable. The tool will automatically find all `.cnc` files within the folder and concatenate them.
  2. The result will be a new `.cnc` file named after the folder.

- **Using Multiple Files**:
  1. Select multiple `.cnc` files (from one or various locations), and drag and drop them onto the executable.
  2. The tool will concatenate them in the order they were selected.

### Output
The program will generate a `.cnc` file with the concatenated contents of the individual files, named after the parent folder. 

## License
This software is offered under the MIT License, permitting free use, modification, redistribution, and private or commercial use as long as the original copyright and license notice are retained.

## Contributions
Contributions are welcome. If you find bugs or would like to suggest enhancements, please open an issue or a pull request on this repository.

## Disclaimer
This tool is provided "as is", without warranty of any kind. Users should verify the output file before use in a production environment to ensure it meets all safety and operational standards.

## Contact
For support or to provide feedback, please create an issue in the project's GitHub repository.