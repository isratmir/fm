### Installation

Download and install it:
```bash
$ go get -u github.com/isratmir/fm
$ go install github.com/isratmir/fm
```

### Usage

```bash
$ $GOPATH/bin/fm
Take rest; a field that has rested gives a bountiful crop. 
- Ovid
```

For convenience, add the workspace's bin subdirectory to your PATH:
```bash
$ export PATH=$PATH:$(go env GOPATH)/bin
```

and you can get quotes with short call:
```bash
$ fm
We must embrace pain and burn it as fuel for our journey. 
- Kenji Miyazawa
```

Use `-l ru` flag to fetch quote in russian language

```bash
$ fm -l ru
Ничего не происходит просто так, нам просто не известны мотивы. 
Доктор Хаус
```
