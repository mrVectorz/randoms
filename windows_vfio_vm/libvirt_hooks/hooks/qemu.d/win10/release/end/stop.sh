#!/bin/bash
# Helpful to read output when debugging
set -x

# Load the config file with our environmental variables
source "/etc/libvirt/hooks/kvm.conf"

# Unload all the vfio modules
modprobe -r vfio_virqfd
modprobe -r vfio_pci
modprobe -r vfio_iommu_type1
modprobe -r vfio

# Reattach the gpu
virsh nodedev-reattach $VIRSH_GPU_VIDEO
virsh nodedev-reattach $VIRSH_GPU_AUDIO

# Load all drivers
if $(lspci -s $(awk 'BEGIN{FS="_"};{print $(NF-2)":"$(NF-1)"."$NF; exit}' /etc/libvirt/hooks/kvm.conf) | grep -q NVIDIA) ; then
	modprobe  nvidia_drm
	modprobe  nvidia_modeset
	modprobe  nvidia_uvm
	modprobe  nvidia
else
	modprobe  amdgpu
fi
modprobe  gpu_sched
modprobe  ttm
modprobe  drm_kms_helper
modprobe  i2c_algo_bit
modprobe  drm
modprobe  snd_hda_intel

# Start you display manager
systemctl start display-manager.service
