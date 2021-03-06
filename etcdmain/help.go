// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package etcdmain

var (
	usageline = `usage: etcd [flags]
       start an etcd server

       etcd --version
       show the version of etcd

       etcd -h | --help
       show the help information about etcd
	`
	flagsline = `
member flags:

	--name 'default'
		human-readable name for this member.
	--data-dir '${name}.etcd'
		path to the data directory.
	--snapshot-count '10000'
		number of committed transactions to trigger a snapshot to disk.
	--heartbeat-interval '100'
		time (in milliseconds) of a heartbeat interval.
	--election-timeout '1000'
		time (in milliseconds) for an election to timeout.
	--listen-peer-urls 'http://localhost:2380,http://localhost:7001'
		list of URLs to listen on for peer traffic.
	--listen-client-urls 'http://localhost:2379,http://localhost:4001'
		list of URLs to listen on for client traffic.
	-cors ''
		comma-separated whitelist of origins for CORS (cross-origin resource sharing).


clustering flags:

	--initial-advertise-peer-urls 'http://localhost:2380,http://localhost:7001'
		list of this member's peer URLs to advertise to the rest of the cluster. 
	--initial-cluster 'default=http://localhost:2380,default=http://localhost:7001'
		initial cluster configuration for bootstrapping.
	--initial-cluster-state 'new'
		initial cluster state ('new' or 'existing').
	--initial-cluster-token 'etcd-cluster'
		initial cluster token for the etcd cluster during bootstrap.
	--advertise-client-urls 'http://localhost:2379,http://localhost:4001'
		list of this member's client URLs to advertise to the rest of the cluster.
	--discovery ''
		discovery URL used to bootstrap the cluster.
	--discovery-fallback 'proxy'
		expected behavior ('exit' or 'proxy') when discovery services fails.
	--discovery-proxy ''
		HTTP proxy to use for traffic to discovery service.
	--discovery-srv ''
		dns srv domain used to bootstrap the cluster.


proxy flags:

	--proxy 'off'
		proxy mode setting ('off', 'readonly' or 'on').


security flags:

	--ca-file ''
		path to the client server TLS CA file.
	--cert-file ''
		path to the client server TLS cert file.
	--key-file ''
		path to the client server TLS key file.
	--client-cert-auth 'false'
		enable client cert authentication.
	--trusted-ca-file ''
		path to the client server TLS trusted CA key file.
	--peer-ca-file ''
		path to the peer server TLS CA file.
	--peer-cert-file ''
		path to the peer server TLS cert file.
	--peer-key-file ''
		path to the peer server TLS key file.
	--peer-client-cert-auth 'false'
		enable peer client cert authentication.
	--peer-trusted-ca-file ''
		path to the peer server TLS trusted CA file.


unsafe flags:

Please be CAUTIOUS to use unsafe flags because it will break the guarantee given 
by consensus protocol. 
	
	--force-new-cluster 'false'
		force to create a new one-member cluster.
`
)
