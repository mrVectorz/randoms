# Single GPU VM VFIO Passthrough
This is notes and configurations on getting a Windows libvirt VM with a gpu passedthrough working
on a host running Fedora 34 with only one GPU.

## Hardware
Getting all this working is very hardware dependant.

CPU: Ryzen 2700x
GPU: Nvidia MSI 1070Ti
MEMORY: 32Gb DDR4 Memory
Motherboard: ASUSTeK PRIME X470-PRO (BIOS version 5216)

## Guest Specifications
The libvirt hooks operate on specific guest names (this is due to the main hook script that I utilized), thus
in this example the guest needs to be named "win10". If you wish to have multiple guests using these hooks
either duplicate the hooks directory with the other names or the main hook script.

Baseline is that it's a 8Gb, 12 vpcu guest. Using the official Windows ISO.

Refer to the guest xml file for all the information.
Of note:
- I had to use the i440fx chipset and UEFI `/usr/share/edk2/ovmf/OVMF_CODE.fd` firmware
- For Nvidia it is required to change the hyperv feature `vendor_id` with any other 12 character string
- Another fun Nvidia specific thing is that you have to specify a ROM file for the card being passthrough

### Finding ROM
In my case I have an MSI 1070Ti on this system. Find your card and download the ROM file.

This is where I downloaded mine:
- [https://www.techpowerup.com/vgabios/195978/msi-gtx1070ti-8192-171013-1](https://www.techpowerup.com/vgabios/195978/msi-gtx1070ti-8192-171013-1)

Once downloaded you will have to put in a path where qemu can access it and have the correct permissions for it.
```
[root@localhost ~]# semanage fcontext -a -t virt_content_t /media/images/MSI.GTX1070Ti.8192.171013_1.rom
[root@localhost ~]# chcon -u system_u /media/images/MSI.GTX1070Ti.8192.171013_1.rom
[root@localhost ~]# ll -Z /media/images/MSI.GTX1070Ti.8192.171013_1.rom 
-rw-r--r--. 1 qemu qemu system_u:object_r:virt_content_t:s0 266240 Jan 19 10:05 /media/images/MSI.GTX1070Ti.8192.171013_1.rom
```

## OS Configurations
Some configurations required.
This is the used cmdline argument to start:
```
BOOT_IMAGE=(hd3,gpt3)/vmlinuz-5.13.6-200.fc34.x86_64 root=/dev/mapper/NVMe-root ro rd.driver.blacklist=nouveau modprobe.blacklist=nouveau nvidia-drm.modeset=1 rd.lvm.lv=NVMe/root rhgb quiet iommu=pt amd_iommu=on video=vesafb:off video=efifb:off
```

- `iommu=pt` is required to enable IOMMU bypass.
- `amd_iommu=on` is to enable the AMD IOMMU driver in the system.
- `video=vesafb:off` this is an old remnant of my previous Intel config, it is to disable the generic Intel graphic framebuffer driver.
- `video=efifb:off` this is to disable the EFI platform driver (on UEFI firmware) for both GOP and UGA displays.

## BIOS Configurations
Not including the OC configurations, just the boot configs.
**NOTE**: I had to try multiple BIOS versions, latest at the time had issues. Settled on version 5216.

- UEFI boot mode
- Secure boot enabled, uses the 'Fedora Secure Boot CA'

## Hooks

I mostly used and customized the classic scrpit from [Sebastiaan Meijer](https://github.com/PassthroughPOST/VFIO-Tools/blob/master/libvirt_hooks/qemu).
Didn't add much other than logging (helped in debugging), then my own steps.
Of note:
- Had to add some sleep timers as I was hitting a race condition when stopping the display manager
- Force unbind the EFI framebuffer (might be fine now that I disabled loading it in the cmdline, haven't retested)

For more information check the scripts.


## Issues

AMD cards may hit the vendor reset issue, and then requires an out of tree kernel module:
- [https://github.com/gnif/vendor-reset](https://github.com/gnif/vendor-reset)

## References
- [https://wiki.archlinux.org/title/PCI_passthrough_via_OVMF](https://wiki.archlinux.org/title/PCI_passthrough_via_OVMF)
- [https://www.reddit.com/r/Fedora/comments/o73i4b/single_gpu_passthrough_guide_fedora_34/](https://www.reddit.com/r/Fedora/comments/o73i4b/single_gpu_passthrough_guide_fedora_34/)
