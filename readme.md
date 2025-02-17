# Folderizer

**Folderizer** is a simple CLI application to organize files in a folder based on their types.

## Installation

Make sure you have **Go** installed and set up properly. Then, clone this repository and build the application:

```sh
# Clone repository
$ git clone https://github.com/riziqalbab/folderizer
$ cd folderizer

# Build the application
$ go build -o folderizer
```

## Usage

Folderizer has several main commands:

### 1. List Files (`list` / `ls`)

Displays the list of files in the current working directory or a specified directory.

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

Creates folders based on file types and moves files into the appropriate folders.

```sh
$ folderizer start
# or
$ folderizer s
```

#### Flags (Optional):

- `--folder_prefix`, `-f`, `-fnp` → Adds a prefix to the folder name.
- `--folder_suffix`, `-s`, `-snp` → Adds a suffix to the folder name.

**Example Usage:**

Suppose there are files in a folder:
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

Will result in:
```
important_documents_files/
  ├── document1.pdf
  ├── document2.docx

important_images_files/
  ├── image1.png
  ├── image2.jpg
```

## Contribution
If you want to contribute, feel free to create a pull request or open an issue in the repository.
