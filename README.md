# Pass-ssh 0.1

A simple script to manage ssh keys with pass.
The ssh-key will be stored in pass with the key name as the file name.
While accessing the key, the script will decrypt the key and store it in a temporary file.
The temporary file will be deleted after the command is executed.
This is nessesary because ssh does not support reading keys from stdin.
At least go my knowledge.

## Installation

```bash
git clone
cd pass-ssh
sudo go install
```

alternatively

```bash
go get github.com/fenge0/pass-ssh
```

## Usage

```bash
pass-ssh
```


### Example

Usage:
  SSH:    pass-ssh <key_name> ssh     <user@host>   [ssh_options]
  Git:    pass-ssh <key_name> git     <git_command> [git_args]
  Create: pass-ssh <git/ssh>  create  <key_name>
  Copy:   pass-ssh <key_name> copy    <user@host>

Examples:
  SSH:    pass-ssh work_laptop ssh user@example.com
  Git:    pass-ssh github_key git clone git@github.com:user/repo.git
          pass-ssh github_key git push origin main
  Create: pass-ssh new_key_name create
  Copy:   pass-ssh work_laptop copy user@example.com
Environment:
  SSH_KEY_ENCRYPTION_LEVEL - Key encryption level (default: 4096)


## Future plans

- [ ] Add to generete ssh keys with other then RSA (DSA, ECDSA, ED25519)
- [ ] add config file to store default values
- [ ] add option to store the key in a different location
- [ ] add to package managers



