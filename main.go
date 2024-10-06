//import "gopkg.in/yaml.v3"

package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Manifest struct {
	API      string     `yaml:"apiVersion"`
	Kind     string     `yaml:"kind"`
	Metadata ObjectMeta `yaml:"metadata"`
	Spec     PodSpec    `yaml:"spec"`
}

type ObjectMeta struct {
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace,omitempty"`
	Labels    map[string]string `yaml:"labels,omitempty"`
}

type PodSpec struct {
	PodOS      string      `yaml:"os,omitempty"`
	Containers []Container `yaml:"containers"`
}

type Container struct {
	Name          string              `yaml:"name"`
	Image         string              `yaml:"image"`
	Ports         []ContainerPort     `yaml:"ports,omitempty"`
	ReadnessProbe Probe               `yaml:"readinessProbe,omitempty"`
	LivenessProbe Probe               `yaml:"livenessProbe,omitempty"`
	Resources     ResourceRequrements `yaml:"resources,omitempty"`
}
type ContainerPort struct {
	Port     int    `yaml:"containerPort"`
	Protocol string `yaml:"protocol,omitempty"`
}

type Probe struct {
	HTTPGet HTTPGetAction `yaml:"httpGet"`
}

type HTTPGetAction struct {
	Path string `yaml:"path"`
	Port int    `yaml:"port"`
}
type ResourceRequrements struct {
	Requests Requests `yaml:"requests"`
	Limits   Limits   `yaml:"limits"`
}

type Requests struct {
	CPU    int    `yaml:"cpu"`
	Memory string `yaml:"memory"`
}
type Limits struct {
	CPU    int    `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

func main() {
	var manifest Manifest
	yamlFile, err := os.ReadFile("test.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &manifest)
	if err != nil {
		fmt.Println(err)
	}
}
