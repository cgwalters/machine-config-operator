package daemon

// This file provides changes that we make to the control plane
// only.

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
)

// setRootDeviceSchedulerBFQ switches to the `bfq` I/O scheduler
// for the root block device to better share I/O between etcd
// and other processes.  See
// https://github.com/openshift/machine-config-operator/issues/1897
func setRootDeviceSchedulerBFQ() error {
	bfq := "bfq"

	rootDevSysfs, err := getRootBlockDeviceSysfs()
	if err != nil {
		return err
	}

	schedulerPath := filepath.Join(rootDevSysfs, "/queue/scheduler")
	schedulerContentsBuf, err := ioutil.ReadFile(schedulerPath)
	schedulerContents := string(schedulerContentsBuf)
	if strings.Contains(schedulerContents, fmt.Sprintf("[%s]", bfq)) {
		return nil
	}

	f, err := os.OpenFile(schedulerPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(bfq))
	if err != nil {
		return err
	}
	glog.Infof("Set root blockdev %s to use scheduler %v", rootDevSysfs, bfq)

	return nil
}

// synchronizeControlPlaneState contains any special dynamic tweaks
// we want to make to control plane nodes.
func (dn *Daemon) synchronizeControlPlaneState() error {
	return setRootDeviceSchedulerBFQ()
}
