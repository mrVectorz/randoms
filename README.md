For making use of `anaconda-ks.cfg`, host the file on a webserver
Then modify the boot param (example while in Grub menu) append to it:
`inst.ks=http://{HOSTNAME OR IP}/filename` example `inst.ks=http://10.19.20.190/anaconda.ks`
