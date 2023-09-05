package main

import (
	"github.com/pulumi/pulumi-rancher2/sdk/v3/go/rancher2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"log"
	"rancher-projects/pkg/parser"
	"rancher-projects/pkg/util"
)

func main() {
	config, err := parser.ParseYAML("./projects/oat/cp.yaml")

	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	pulumi.Run(func(ctx *pulumi.Context) error {
		for _, project := range config.RancherProjects {
			newProject, err := rancher2.NewProject(ctx, project.ProjectName, &rancher2.ProjectArgs{
				Name:      pulumi.String(project.ProjectName),
				ClusterId: pulumi.String("c-bmdb2"),
			})

			for _, namespace := range project.Namespaces {
				_, err := rancher2.NewNamespace(ctx, namespace.Name, &rancher2.NamespaceArgs{
					Name:      pulumi.String(namespace.Name),
					Labels:    util.ConvertToMapInput(namespace.Labels),
					ProjectId: newProject.ID(),
				})
				if err != nil {
					return err
				}
			}
			if err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}

		return nil
	})
}
