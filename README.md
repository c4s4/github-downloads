# Github Downloads

This tool prints downloads for given Github repo.

## Installation

### Unix users (Linux, BSDs and MacOSX)

Unix users may download and install latest *github-downloads* release with command:

```bash
sh -c "$(curl http://sweetohm.net/dist/github-downloads/install)"
```

If *curl* is not installed on you system, you might run:

```bash
sh -c "$(wget -O - http://sweetohm.net/dist/github-downloads/install)"
```

**Note:** Some directories are protected, even as *root*, on **MacOSX** (since *El Capitan* release), thus you can't install *github-downloads* in */usr/bin* for instance.

### Binary package

Otherwise, you can download latest binary archive at <https://github.com/c4s4/github-downloads/releases>. Unzip the archive, put the binary of your platform somewhere in your *PATH* and rename it *github-downloads*.

## Usage

To print downloads for repository *bar* of account *foo*, you would type in a terminal:

```bash
$ github-downloads foo bar
Release: 0.2.0
- bar-0.2.0.tar.gz: 12
Release: 0.1.0
- bar-bin-0.1.0.tar.gz: 2
Total: 14
```

This will print, for each release, the list of its assets and the number of times they were downloaded and the total downloads for this repository.

*Enjoy:*
