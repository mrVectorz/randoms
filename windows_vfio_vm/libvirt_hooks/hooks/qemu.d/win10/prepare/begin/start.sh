#!/bin/bash
set -x
source "/etc/libvirt/hooks/kvm.conf"

# Stop your display manager. If youre on kde it ll be sddm.service. Gnome users should use killall gdm-x-session instead
logger "start hook - killing display"
systemctl stop display-manager.service
killall gdm-x-session
kill $(pgrep Xorg)

pulse_pid=$(pgrep -u unknown pulseaudio)
pipewire_pid=$(pgrep -u unknown pipewire)
kill $pulse_pid
kill $pipewire_pid

sleep 10

# Unbind VTconsoles
logger "start hook - unbinding"
echo 0 > /sys/class/vtconsole/vtcon0/bind
echo 0 > /sys/class/vtconsole/vtcon1/bind
echo efi-framebuffer.0 > /sys/bus/platform/drivers/efi-framebuffer/unbind 

# Avoid a race condition by waiting a couple of seconds. This can be calibrated to be shorter or longer if required for your system
sleep 10

logger "start hook - unloading mods"
if $(lspci -s $(awk 'BEGIN{FS="_"};{print $(NF-2)":"$(NF-1)"."$NF; exit}' /etc/libvirt/hooks/kvm.conf) | grep -q NVIDIA) ; then
	# Unload all nvidia gpu drivers
	modprobe -r nvidia_drm
	modprobe -r nvidia_modeset
	modprobe -r nvidia_uvm
	modprobe -r nvidia
else
	# Unload all amd gpu drivers
	modprobe -r amdgpu
fi
 
# Unbind the GPU from display drive
logger "start hook - nodedev-detach"
virsh nodedev-detach $VIRSH_GPU_VIDEO
virsh nodedev-detach $VIRSH_GPU_AUDIO
 
# Load VFIO kernel module
modprobe vfio
modprobe vfio_pci
modprobe vfio_iommu_type1
modprobe vfio_virqfd
