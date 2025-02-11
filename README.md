# Example use of age in Go code

taken from 

https://pkg.go.dev/filippo.io/age#example-Encrypt

https://pkg.go.dev/filippo.io/age#example-Decrypt

this is an example that is a little more complete than the examples above and that use an input file and public and private keys that are stored in an `.env` file

the input and output files, be they plain text, Encrypted or decrypted are still hard coded in the go code and how to handle these with configurable values is left to the reader to implement with environment variables, command line parameters, database queries, external configuration etc or whatever takes your fancy, fill your boots

for convenience and to create a set of keys, initially install `age` with your package manager, for example

```bash
sudo apt update && sudo apt install age
```

so that a key pair may be created with

```bash
❯ age-keygen -o age_key.txt 
Public key: ....
```

NB: if you do this more than once it will not over write an existing key file if one already exists with this filename

then create an .env file that has the public and private keys that this file contains :

```bash
❯ cat .env 
PUBLIC_KEY=YOUR PUBLIC KEY GOES HERE
PRIVATE_KEY=YOUR PRIVATE_KEY GOES HERE
```

to run the go code examples 

```bash
❯ go run encrypt/main.go 
Encrypted file size: 247
```

then decrypt with

```bash
❯ go run decrypt/main.go 
Decryption successful, data written to decrypted_file.txt
```

edit the file `input.txt` replacing it with your message, re-run the above to prove to yourself things are working as expected

to build executables

```bash
go build -o enc encrypt/main.go 
go build -o dec decrypt/main.go 
```

and copy the binaries `enc` and `dec` together with the `.env` file created earlier and an `input.txt` file to run elsewhere 

```bash
$ ./enc 
Encrypted file size: 266
$ ./dec 
Decryption successful, data written to decrypted_file.txt
$ cat decrypted_file.txt 
I'm forever blowing bubbles,  

pretty bubbles in the air !

.fin
```

on the remote system, we dont need age installed now as we have created our own 'mini apps' with age built in