# Folderizer

**Folderizer** is a simple CLI application for organizing files into folders based on their types.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [List Files (`list` / `ls`)](#1-list-files-list--ls)
  - [Create Folders and Organize Files (`start` / `s`)](#2-create-folders-and-organize-files-start--s)
  - [Move Files to Destination Folder (`set` / `st`)](#3-move-files-to-destination-folder-set--st)
- [Contribution](#contribution)
- [License](#license)

## Installation

Make sure you have **Go** installed and a proper environment set up. Then, clone this repository and build the application:

```sh
# Clone repository
$ git clone https://github.com/username/folderizer.git
$ cd folderizer

# Build the application
$ go build -o folderizer
```

## Usage

Folderizer provides several main commands:

### 1. List Files (`list` / `ls`)

Used to display the list of files in the current working directory or a specified directory.

```sh
$ folderizer list
# or
$ folderizer ls
```

To list files in a specific directory:

```sh
$ folderizer list /path/to/directory
```

### 2. Create Folders and Organize Files (`start` / `s`)

Used to create folders based on file types and move files into their respective folders.

```sh
$ folderizer start
# or
$ folderizer s
```

#### Optional Flags:

- `--folder_prefix`, `-f`, `-fnp` → Adds a prefix to folder names.
- `--folder_suffix`, `-s`, `-snp` → Adds a suffix to folder names.

**Example Usage:**

Assume the following files are in a folder:
```
document1.pdf
document2.docx
image1.png
image2.jpg
```

Running the following command:
```sh
$ folderizer start --folder_prefix important --folder_suffix files
```

Will produce:
```
important_documents_files/
  ├── document1.pdf
  ├── document2.docx

important_images_files/
  ├── image1.png
  ├── image2.jpg
```

### 3. Move Files to Destination Folder (`set` / `st`)

Used to move files matching the specified flags into a target folder.

```sh
$ folderizer set
# or
$ folderizer st
```

#### Flags:

- `--file_prefix`, `-fp` → Detects a prefix in the file name.
- `--file_suffix`, `-fs` → Detects a suffix in the file name.
- `--to_folder`, `-tf` → Specifies the destination folder for matching files.

**Example Usage:**

Assume the following files exist:
```
report_2023.pdf
report_2024.pdf
data_final.csv
data_raw.csv
```

Running the following command:
```sh
$ folderizer set --file_prefix report --to_folder reports
```

Will produce:
```
reports/
  ├── report_2023.pdf
  ├── report_2024.pdf

data_final.csv
data_raw.csv
```

## Contribution
If you want to contribute, feel free to submit a pull request or open an issue in the repository.
