# polyllama
## About
polyllama is FOSS cli tool for local translation writen in golang.
It uses ollama to load and interact with LLMs such as Mistral or Llama 2.

## Getting started
### Installation
1. Install ollama
2. Install go
3. Clone this repository
```sh
git clone https://github.com/mykykh/polyllama.git
```
4. Cd to folder with polyllama
```
cd polyllama
```
5. Install polyllama
```sh
go install
```
6. Add go/bin to your PATH

### Usage
You can translate text by using translate command:
```sh
polyllama translate --target-lang french --text "Hello world!"
```

You can translate from any language your model knows to any other:
```sh
polyllama translate --source-lang ukrainian --target-lang french --text "Слава Україні!"
```

By default polyllama uses mistral, but
you can use any model that ollama supports:
```sh
polyllama t --tl french --model llama2
```

To translate file, specify path to it at the end:
```sh
polyllama t --sl ukrainian --tl french -m mistral path/to/file
```

To save translation to file, just pipe polyllama output:
```sh
polyllama t --sl ukrainian --tl french path/to/source > path/to/save
```

## Licence
Distributed under GPLv3 Licence. See `LICENSE` for more information.
