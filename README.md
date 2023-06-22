# Encrypt-Seeds
Encrypt your Mnemonics with PIN

You can easily encrypt or decrypt the mnemonic of your metamask or other decentralized wallets using 4 to 8 digits pin.

#
## Steps for setup:

### 1. Clone the respository

```shell
$ git clone https://github.com/knackroot-technolabs-llp/encrypt-seeds.git

```

### 2. Get into the folder

```shell
$ cd encrypt-seeds

```

### 3. Install the dependencies

```shell
$ go get github.com/urfave/cli/v2
$ go get gitlab.com/david_mbuvi/go_asterisks

```
#
## Steps to run the project

### 1. Build the project

```shell
$ go build encrypt-seeds.go

``` 

### 2. Run the project

```shell
$ ./encrypt-seeds [commands]

```
#

## To encrypt the mnemonics

### Run the encrypt command

```shell
$ ./encrypt-seeds encrypt

```

1. Enter the number of words, your seed-phrase has.
2. Enter the passphrase of 4 to 8 digits.
3. Confirm your passphrase.
4. Enter the seed phrase one-by-one and press enter after each word.
5. Store the encrypted seed phrase somewhere and remember your passphrase.

#

## To decrypt the mnemonics

### Run the decrypt command

```shell
$ ./encrypt-seeds decrypt

```

1. Enter the number of words, your seed-phrase has.
2. Enter the passphrase of 4 to 8 digits
3. Enter the encrypted seed phrase one-by-one and press enter after each word.
4. You will get your original seed phrase.


