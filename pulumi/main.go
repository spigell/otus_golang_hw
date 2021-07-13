package main

import (
	"github.com/pulumi/pulumi-github/sdk/v4/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := github.NewRepository(ctx, "repository", &github.RepositoryArgs{
			AllowMergeCommit:    pulumi.Bool(true),
			AllowRebaseMerge:    pulumi.Bool(true),
			AllowSquashMerge:    pulumi.Bool(true),
			Archived:            pulumi.Bool(false),
			DeleteBranchOnMerge: pulumi.Bool(false),
			HasDownloads:        pulumi.Bool(true),
			HasIssues:           pulumi.Bool(true),
			HasProjects:         pulumi.Bool(true),
			HasWiki:             pulumi.Bool(true),
			Name:                pulumi.String("otus_golang_hw"),
			Template: &github.RepositoryTemplateArgs{
				Owner:      pulumi.String("OtusGolang"),
				Repository: pulumi.String("home_work"),
			},
		}, pulumi.Protect(true),
		)
		if err != nil {
			return err
		}
		return nil
	})

}
