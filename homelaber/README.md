## Homelaber
### Notes

Due to changing to using Silverblue it was easier to use `/usr/local/bin/homelaber` as a path, instead of `/usr/bin/` as the latter is mounted on a RO partition.
However, for this path to work in a systemd service, it needs to be relabled.
```
semanage fcontext -a -t bin_t "/usr/local/bin/"
restorecon -r -v /usr/local/bin/
```

Will move to using a container eventually.
