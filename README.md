# Standard Notes export to folder

This is a CLI tool used to extract info from the [Standard Notes](https://standardnotes.com) decrypted export and transform it into folder based tags and notes.
After the tool runs the tags will be folders and each note will be inside the respective folder.

## Usage

Download a decrypted backup of your Standard Notes account and extract the `zip` and locate the decrypted backup file, it's named `Standard Notes Backup and Import File.txt`

- Download the latest release for your platform [here](https://github.com/DanielNetoP/sn-export-to-folder/releases).
- Extract the file downloadable file.
- Run the CLI tool
  - Windows
    - `sn-export-folder.exe path/to/Standard Notes Backup and Import File.txt`
  - Linux/MacOS
    - `./sn-export-folder path/to/Standard Notes Backup and Import File.txt`
- Find the exported tags/notes in the `exported` directory

### Example
```
# If you have the backup file on the same folder as the sn-export-folder tool
## Windows
sn-export-folder.exe "Standard Notes Backup and Import File.txt"
## Linux/MacOS
./sn-export-folder "Standard Notes Backup and Import File.txt"
```

### MacOS

When executing the above command you might get the following message:

`"sn-export-folder" can't be opened because Apple cannot check it for malicious software`

Press `Show in Finder` and right-click on the tool and click `Open` after that you might get another message:

`macOS cannot verify the developer of "sn-export-folder". Are you sure you want to open it?`

Press `Open` and a new terminal window should open with the following:
```
sn-export-folder-1.0.6-darwin-arm64/sn-export-folder ; exit;
path for `Standard Notes Backup File` is required
```

Close the terminal window and follow the steps on [Usage](#usage).


## Badges

![Build Status](https://github.com/DanielNetoP/sn-export-to-folder/workflows/Test/badge.svg)
[![License](https://img.shields.io/github/license/DanielNetoP/sn-export-to-folder)](/LICENSE)
[![Release](https://img.shields.io/github/release/DanielNetoP/sn-export-to-folder.svg)](https://github.com/DanielNetoP/sn-export-to-folder/releases/latest)