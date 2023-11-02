package install

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gofulljs/gbook/global"
	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_parseBookJsonPlugins(t *testing.T) {
	tests := []struct {
		name        string
		content     []byte
		wantMPlugin map[string]struct{}
		wantErr     error
	}{
		// TODO: Add test cases.
		{
			name: "valid book.json",
			content: []byte(`
{
	"plugins": [
		"-lunr",
		"-search",
		"a-import",
		"code",
		"github"
	]
}			
			`),
			wantMPlugin: map[string]struct{}{
				"gitbook-plugin-a-import": {},
				"gitbook-plugin-code":     {},
				"gitbook-plugin-github":   {},
			},
		},
		{
			name: "invalid book.json",
			content: []byte(`
{
	"plugins": [
		"-lunr": {},
		"-search",
		"a-import",
		"code",
		"github"
	]
}			
			`),
			wantErr: errInvalidBookJson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMPlugin, err := parseBookJsonPlugins(tt.content)
			assert.Equal(t, tt.wantErr, errors.Unwrap(err))
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantMPlugin, gotMPlugin)
		})
	}
}

func Test_parseAlreadyPlugins(t *testing.T) {
	output := []byte(`{
		"problems": [
		  "extraneous: boolbase@1.0.0 D:\\my\\mynotes\\note_k8s_2hour_intro\\node_modules\\boolbase"
		],
		"dependencies": {
		  "entities": {
			"version": "1.1.2",
			"resolved": "https://registry.npmmirror.com/entities/-/entities-1.1.2.tgz",
			"extraneous": true,
			"problems": [
			  "extraneous: entities@1.1.2 D:\\my\\mynotes\\note_k8s_2hour_intro\\node_modules\\entities"
			]
		  },
		  "gitbook-plugin-3-ba": {
			"version": "0.9.0",
			"resolved": "https://registry.npmmirror.com/gitbook-plugin-3-ba/-/gitbook-plugin-3-ba-0.9.0.tgz",
			"extraneous": true,
			"problems": [
			  "extraneous: gitbook-plugin-3-ba@0.9.0 D:\\my\\mynotes\\note_k8s_2hour_intro\\node_modules\\gitbook-plugin-3-ba"
			]
		  },
		  "gitbook-plugin-a-import": {
			"version": "0.0.3",
			"extraneous": true,
			"problems": [
			  "extraneous: gitbook-plugin-a-import@0.0.3 D:\\my\\mynotes\\note_k8s_2hour_intro\\node_modules\\gitbook-plugin-a-import"
			]
		  },
		  "whatwg-url": {
			"version": "5.0.0",
			"extraneous": true,
			"problems": [
			  "extraneous: whatwg-url@5.0.0 D:\\my\\mynotes\\note_k8s_2hour_intro\\node_modules\\whatwg-url"
			]
		  }
		}
	  }`)

	m := parseAlreadyPlugins(output)

	assert.Equal(t, map[string]struct{}{
		"gitbook-plugin-3-ba":     {},
		"gitbook-plugin-a-import": {},
	}, m)
}

func TestParse(t *testing.T) {
	// versions := []string{
	// 	">=4.0.0-alpha.0",
	// 	">=2.5.0",
	// 	"*",
	// 	"*",
	// 	"*",
	// }

	constraints, err := version.NewConstraint(">=4.0.0-alpha.0")
	assert.NoError(t, err)

	// constraints, err := version.NewConstraint("*")
	// assert.NoError(t, err)

	// bookVersion, err := version.NewVersion(global.BOOK_VERSION)
	bookVersion, err := version.NewVersion("4.0.0")
	assert.NoError(t, err)

	fmt.Println(constraints.Check(bookVersion))
}

func Test_parseValidPluginName(t *testing.T) {
	tests := []struct {
		name            string
		pluginNoVersion string
		output          []byte
		bookVersion     string
		want            string
		wantErr         error
	}{
		// TODO: Add test cases.
		{
			name:            "github",
			pluginNoVersion: "gitbook-plugin-github",
			output: []byte(`gitbook-plugin-github@3.0.0 '>=4.0.0-alpha.0'
gitbook-plugin-github@2.0.0 '>=2.5.0'
gitbook-plugin-github@1.1.0 '*'
gitbook-plugin-github@1.0.2 '*'
gitbook-plugin-github@1.0.0 '*'`),
			bookVersion: global.BOOK_VERSION,
			want:        "gitbook-plugin-github@2.0.0",
		},
		{
			name:            "3-ba",
			pluginNoVersion: "gitbook-plugin-3-ba",
			output:          []byte(">=3.0.0"),
			bookVersion:     global.BOOK_VERSION,
			want:            "gitbook-plugin-3-ba",
		},
		{
			name:            "anchor-navigation-ex",
			pluginNoVersion: "gitbook-plugin-anchor-navigation-ex",
			output: []byte(`gitbook-plugin-anchor-navigation-ex@1.2.5 '>3.x.x'
gitbook-plugin-anchor-navigation-ex@1.0.13 '>=3.0.0'
gitbook-plugin-anchor-navigation-ex@1.0.14 '>=3.0.0'`),
			bookVersion: global.BOOK_VERSION,
			want:        "gitbook-plugin-anchor-navigation-ex@1.0.14",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseValidPluginName(tt.pluginNoVersion, tt.output, tt.bookVersion)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
