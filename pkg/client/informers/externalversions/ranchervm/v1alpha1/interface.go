/*
Copyright 2018 Rancher Labs, Inc.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/rancher/vm/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ARPTables returns a ARPTableInformer.
	ARPTables() ARPTableInformer
	// Credentials returns a CredentialInformer.
	Credentials() CredentialInformer
	// MachineImages returns a MachineImageInformer.
	MachineImages() MachineImageInformer
	// Settings returns a SettingInformer.
	Settings() SettingInformer
	// VirtualMachines returns a VirtualMachineInformer.
	VirtualMachines() VirtualMachineInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ARPTables returns a ARPTableInformer.
func (v *version) ARPTables() ARPTableInformer {
	return &aRPTableInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Credentials returns a CredentialInformer.
func (v *version) Credentials() CredentialInformer {
	return &credentialInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// MachineImages returns a MachineImageInformer.
func (v *version) MachineImages() MachineImageInformer {
	return &machineImageInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Settings returns a SettingInformer.
func (v *version) Settings() SettingInformer {
	return &settingInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// VirtualMachines returns a VirtualMachineInformer.
func (v *version) VirtualMachines() VirtualMachineInformer {
	return &virtualMachineInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
