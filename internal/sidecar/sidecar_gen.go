// Code generated by gonstructor -type Sidecar -constructorTypes builder; DO NOT EDIT.

package sidecar

import (
	"sync"

	"github.com/reconquest/snake-runner/internal/cloud"
	"github.com/reconquest/snake-runner/internal/sshkey"
)

type SidecarBuilder struct {
	cloud          *cloud.Cloud
	name           string
	pipelinesDir   string
	slug           string
	promptConsumer cloud.PromptConsumer
	outputConsumer cloud.OutputConsumer
	sshKey         sshkey.Key
	sshAgent       sync.WaitGroup
}

func NewSidecarBuilder() *SidecarBuilder {
	return &SidecarBuilder{}
}

func (b *SidecarBuilder) Cloud(cloud *cloud.Cloud) *SidecarBuilder {
	b.cloud = cloud
	return b
}

func (b *SidecarBuilder) Name(name string) *SidecarBuilder {
	b.name = name
	return b
}

func (b *SidecarBuilder) PipelinesDir(pipelinesDir string) *SidecarBuilder {
	b.pipelinesDir = pipelinesDir
	return b
}

func (b *SidecarBuilder) Slug(slug string) *SidecarBuilder {
	b.slug = slug
	return b
}

func (b *SidecarBuilder) PromptConsumer(promptConsumer cloud.PromptConsumer) *SidecarBuilder {
	b.promptConsumer = promptConsumer
	return b
}

func (b *SidecarBuilder) OutputConsumer(outputConsumer cloud.OutputConsumer) *SidecarBuilder {
	b.outputConsumer = outputConsumer
	return b
}

func (b *SidecarBuilder) SshKey(sshKey sshkey.Key) *SidecarBuilder {
	b.sshKey = sshKey
	return b
}

func (b *SidecarBuilder) SshAgent(sshAgent sync.WaitGroup) *SidecarBuilder {
	b.sshAgent = sshAgent
	return b
}

func (b *SidecarBuilder) Build() *Sidecar {
	return &Sidecar{
		cloud:          b.cloud,
		name:           b.name,
		pipelinesDir:   b.pipelinesDir,
		slug:           b.slug,
		promptConsumer: b.promptConsumer,
		outputConsumer: b.outputConsumer,
		sshKey:         b.sshKey,
		sshAgent:       b.sshAgent,
	}
}
