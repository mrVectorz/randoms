#!/bin/bash
set -x
source "/etc/libvirt/hooks/kvm.conf"

# Stop your display manager. If youre on kde it ll be sddm.service. Gnome users should use killall gdm-x-session instead
systemctl stop display-manager.service
pulse_pid=$(pgrep -u unknown pulseaudio)
pipewire_pid=$(pgrep -u unknown pipewire)
kill $pulse_pid
kill $pipewire_pid
 
# Unbind VTconsoles
echo 0 > /sys/class/vtconsole/vtcon0/bind
echo 0 > /sys/class/vtconsole/vtcon1/bind
 
# Avoid a race condition by waiting a couple of seconds. This can be calibrated to be shorter or longer if required for your system
sleep 5
 
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
 
# Unbind the GPU from display driver
virsh nodedev-detach $VIRSH_GPU_VIDEO
virsh nodedev-detach $VIRSH_GPU_AUDIO
 
# Load VFIO kernel module
modprobe vfio
modprobe vfio_pci
modprobe vfio_iommu_type1
