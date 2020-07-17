package daemon

import (
	"context"
	"time"

	"github.com/golang/glog"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/pkg/transport"
	"google.golang.org/grpc"
)

// This file contains code to find the etcd leader.

// type etcdCtlHeader struct {
// 	ClusterID string `json:"cluster_id"`
// 	MemberID  string `json:"member_id"`
// }

// type etcdCtlMember struct {
// 	ID   string
// 	Name string `json:"name"`
// }

// type etcdctlOutput struct {
// 	header  etcdCtlHeader
// 	members []etcdCtlMember
// }

// WatchCurrentEtcdLeader creates a channel that signals changes in the etcd leader
func WatchCurrentEtcdLeader() (<-chan string, error) {
	dialOptions := []grpc.DialOption{
		grpc.WithBlock(), // block until the underlying connection is up
	}

	// TODO - do this in a better way; it might just be having the etcd pod write out
	// a file in /run when it becomes the leader.
	tlsInfo := transport.TLSInfo{
		CertFile:      "/etc/kubernetes/static-pod-resources/kube-apiserver-pod-2/secrets/etcd-client/tls.crt",
		KeyFile:       "/etc/kubernetes/static-pod-resources/kube-apiserver-pod-2/secrets/etcd-client/tls.key",
		TrustedCAFile: "/etc/kubernetes/static-pod-resources/kube-apiserver-pod-2/configmaps/etcd-serving-ca/ca-bundle.crt",
	}
	tlsConfig, err := tlsInfo.ClientConfig()

	cfg := clientv3.Config{
		DialOptions: dialOptions,
		Endpoints:   []string{"https://localhost:2379"},
		DialTimeout: 15 * time.Second,
		TLS:         tlsConfig,
	}

	c, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}

	_, err = c.MemberList(context.TODO())
	if err != nil {
		return nil, err
	}

	r := make(chan string)
	go func() {
		lastLeader := ""
		for {
			glog.Infof("getting member list")
			memberResp, err := c.MemberList(context.TODO())
			if err != nil {
				glog.Errorf("Failed to get etcd MemberList: %v", err)
				time.Sleep(10 * time.Second)
				continue
			}
			leaderID := memberResp.Header.MemberId
			glog.Infof("leader %d", leaderID)
			leaderName := ""
			for _, member := range memberResp.Members {
				if member.ID != leaderID {
					continue
				}
				leaderName = member.Name
			}
			if leaderName == "" {
				glog.Warningf("Failed to find leader id %d in member list", leaderID)
			} else {
				if leaderName != lastLeader {
					r <- leaderName
					lastLeader = leaderName
				}
			}
			time.Sleep(30 * time.Second)
		}
	}()

	return r, nil
}
