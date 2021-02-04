/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	clientset "k8s.io/client-go/kubernetes"
	sshutil "k8s.io/kubernetes/pkg/ssh"
)

type KindProvider struct {
	features Features
}

func NewKindProvider(_ map[string]string) Provider {
	return &KindProvider{
		features: Features{
			SupportProbe:                        true,
			SupportSSHToMaster:                  false,
			SupportImagePreload:                 true,
			SupportEnablePrometheusServer:       true,
			SupportGrabMetricsFromKubelets:      true,
			SupportAccessAPIServerPprofEndpoint: true,
		},
	}
}

func (p *KindProvider) Name() string {
	return KindName
}

func (p *KindProvider) Features() *Features {
	return &p.features
}

func (p *KindProvider) GetComponentProtocolAndPort(componentName string) (string, int, error) {
	return getComponentProtocolAndPort(componentName)
}

func (p *KindProvider) GetConfig() Config {
	return Config{}
}

func (p *KindProvider) RunSSHCommand(cmd, host string) (string, string, int, error) {
	// TODO(#1693): To maintain compatibility with the use of SSH to scrape measurements,
	// we can SSH to localhost then run `docker exec -t <masterNodeContainerID> <cmd>`.
	signer, err := sshSignerFromKeyFile("LOCAL_SSH_KEY", "id_rsa")
	if err != nil {
		return "", "", 0, err
	}
	user := defaultSSHUser()
	return sshutil.RunSSHCommand(cmd, user, host, signer)
}

func (p *KindProvider) Metadata(client clientset.Interface) (map[string]string, error) {
	return nil, nil
}
