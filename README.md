## goansible
```
A simple ansible sdk library for golang.
```

### ansible.cfg
Add items to /etc/ansible/ansible.cfg.
```
[defaults]
host_key_checking = False
stdout_callback = json
bin_ansible_callbacks = True
```

### features
```
- adhoc
- playbook
- modules
  - user
  - file
  - shell
  - script
  - synchronize
  - cron
```

### examples
```
- adhoc
- playbook
- modules
```

### install
```bazaar
go get github.com/peng0208/goansible
```
