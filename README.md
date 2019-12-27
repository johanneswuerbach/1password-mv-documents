# Move documents between 1password vaults

Currently 1Password does not support to move items in the Documents category directly between 1Password vaults on different servers:

> Documents can't be moved from an account on one server to another.
> To move to the destination account, save the file from the document item and recreate it in the destination vault.

This tool automates those steps and can be used to move documents from the US to EU instance.

## Requirements

- 1Password CLI
- Golang 1.13+

## Usage example

```
op signin my.1password.com <email> <secret-key> --shorthand=my-com
op signin my.1password.eu <email> <secret-key> --shorthand=my-eu
```


```
go run main.go --origin-shorthand my-com --origin-vault Personal --target-shorthand my-eu --target-vault Private
```
