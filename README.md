# Simple Archiver

This archiver tool allows you to compress `.txt` files using VLC compression.

## How It Works

### Packing a File

To pack a `.txt` file, use the following command:

```bash
./Archiver pack vlc test.txt
```
For example, given a file test.txt with the content:
```txt
Hello everyone
```
The command will generate a `test.vlc` binary file. This file will be 5 bytes smaller than the original `.txt` file.

![Снимок экрана 2024-08-09 в 13 39 59](https://github.com/user-attachments/assets/f19dd7a5-3b5b-4683-a05d-f784eebd1e31)

### Unpacking a File

To unpack a `.vlc` file back to its original `.txt` format, use the following command:
```bash
./Archiver unpack vlc test.vlc
```
This will restore the original content of the file.
