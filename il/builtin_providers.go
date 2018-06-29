package il

import (
	"github.com/terraform-providers/terraform-provider-archive/archive"
	"github.com/terraform-providers/terraform-provider-http/http"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
)

var builtinProviderInfo = map[string]*tfbridge.ProviderInfo {
	"archive": {
		P: archive.Provider().(*schema.Provider),
		Config: map[string]*tfbridge.SchemaInfo{},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"archive_file": {Tok: "archive:archive:archiveFile"},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"archive_file": {Tok: "archive:archive/archiveFile:ArchiveFile"},
		},
	},
	"http": {
		P: http.Provider().(*schema.Provider),
		Config: map[string]*tfbridge.SchemaInfo{},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"http": {Tok: "http:http:http"},
		},
		Resources: map[string]*tfbridge.ResourceInfo{},
	},
}