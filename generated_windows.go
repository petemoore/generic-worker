// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package main

import (
	"encoding/json"
	"errors"

	tcclient "github.com/taskcluster/taskcluster-client-go"
)

type (
	// Requires scope `queue:get-artifact:<artifact-name>`
	ArtifactContent struct {

		// Max length: 1024
		Artifact string `json:"artifact"`

		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		TaskID string `json:"taskId"`
	}

	Content json.RawMessage

	FileMount struct {

		// Content of the file to be mounted
		Content Content `json:"content"`

		// The filesystem location to mount the file
		File string `json:"file"`
	}

	// This schema defines the structure of the `payload` property referred to in a
	// Taskcluster Task definition.
	GenericWorkerPayload struct {

		// Artifacts to be published. For example:
		// `{ "type": "file", "path": "builds\\firefox.exe", "expires": "2015-08-19T17:30:00.000Z" }`
		Artifacts []struct {

			// Explicitly set the value of the HTTP `Content-Type` response header when the artifact(s)
			// is/are served over HTTP(S). If not provided (this property is optional) the worker will
			// guess the content type of artifacts based on the filename extension of the file storing
			// the artifact content. It does this by looking at the system filename-to-mimetype mappings
			// defined in the Windows registry. Note, setting `contentType` on a directory artifact will
			// apply the same contentType to all files contained in the directory.
			//
			// See [mime.TypeByExtension](https://godoc.org/mime#TypeByExtension).
			ContentType string `json:"contentType,omitempty"`

			// Date when artifact should expire must be in the future, no earlier than task deadline, but
			// no later than task expiry. If not set, defaults to task expiry.
			Expires tcclient.Time `json:"expires,omitempty"`

			// Name of the artifact, as it will be published. If not set, `path` will be used.
			// Conventionally (although not enforced) path elements are forward slash separated. Example:
			// `public/build/a/house`. Note, no scopes are required to read artifacts beginning `public/`.
			// Artifact names not beginning `public/` are scope-protected (caller requires scopes to
			// download the artifact). See the Queue documentation for more information.
			Name string `json:"name,omitempty"`

			// Relative path of the file/directory from the task directory. Note this is not an absolute
			// path as is typically used in docker-worker, since the absolute task directory name is not
			// known when the task is submitted. Example: `dist\regedit.exe`. It doesn't matter if
			// forward slashes or backslashes are used.
			Path string `json:"path"`

			// Artifacts can be either an individual `file` or a `directory` containing
			// potentially multiple files with recursively included subdirectories.
			//
			// Possible values:
			//   * "file"
			//   * "directory"
			Type string `json:"type"`
		} `json:"artifacts,omitempty"`

		// One entry per command (consider each entry to be interpreted as a full line of
		// a Windows™ .bat file). For example:
		// `["set", "echo hello world > hello_world.txt", "set GOPATH=C:\\Go"]`.
		Command []string `json:"command"`

		// Env vars must be string to __string__ mappings (not number or boolean). For example:
		// ```
		// {
		//   "PATH": "C:\\Windows\\system32;C:\\Windows",
		//   "GOOS": "windows",
		//   "FOO_ENABLE": "true",
		//   "BAR_TOTAL": "3"
		// }
		// ```
		Env json.RawMessage `json:"env,omitempty"`

		// Feature flags enable additional functionality.
		Features struct {

			// An artifact named `public/chainOfTrust.json.asc` should be generated
			// which will include information for downstream tasks to build a level
			// of trust for the artifacts produced by the task and the environment
			// it ran in.
			ChainOfTrust bool `json:"chainOfTrust,omitempty"`
		} `json:"features,omitempty"`

		// Maximum time the task container can run in seconds
		//
		// Mininum:    1
		// Maximum:    86400
		MaxRunTime int64 `json:"maxRunTime"`

		// Directories and/or files to be mounted
		Mounts []Mount `json:"mounts,omitempty"`

		// A list of OS Groups that the task user should be a member of. Requires
		// scope `generic-worker:os-group:<os-group>` for each group listed.
		OSGroups []string `json:"osGroups,omitempty"`

		// URL of a service that can indicate tasks superseding this one; the current `taskId`
		// will be appended as a query argument `taskId`. The service should return an object with
		// a `supersedes` key containing a list of `taskId`s, including the supplied `taskId`. The
		// tasks should be ordered such that each task supersedes all tasks appearing later in the
		// list.  See
		// [superseding](https://docs.taskcluster.net/reference/platform/taskcluster-queue/docs/superseding)
		// for more detail.
		SupersederURL string `json:"supersederUrl,omitempty"`
	}

	Mount json.RawMessage

	ReadOnlyDirectory struct {

		// Contents of read only directory.
		Content Content `json:"content"`

		// The filesystem location to mount the directory volume
		Directory string `json:"directory"`

		// Archive format of content for read only directory
		//
		// Possible values:
		//   * "rar"
		//   * "tar.bz2"
		//   * "tar.gz"
		//   * "zip"
		Format string `json:"format"`
	}

	// URL to download content from
	URLContent struct {

		// URL to download content from
		URL string `json:"url"`
	}

	Var FileMount

	Var1 WritableDirectoryCache

	Var2 ReadOnlyDirectory

	WritableDirectoryCache struct {

		// Implies a read/write cache directory volume. A unique name for the cache volume. Requires scope `generic-worker:cache:<cache-name>`. Note if this cache is loaded from an artifact, you will also require scope `queue:get-artifact:<artifact-name>` to use this cache.
		CacheName string `json:"cacheName"`

		// Optional content to be preloaded when initially creating the cache (if set, `format` must also be provided).
		Content Content `json:"content,omitempty"`

		// The filesystem location to mount the directory volume
		Directory string `json:"directory"`

		// Archive format of the preloaded content (if `content` provided).
		//
		// Possible values:
		//   * "rar"
		//   * "tar.bz2"
		//   * "tar.gz"
		//   * "zip"
		Format string `json:"format,omitempty"`
	}
)

// MarshalJSON calls json.RawMessage method of the same name. Required since
// Content is of type json.RawMessage...
func (this *Content) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *Content) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("Content: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}

// MarshalJSON calls json.RawMessage method of the same name. Required since
// Mount is of type json.RawMessage...
func (this *Mount) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *Mount) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("Mount: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}

// Returns json schema for the payload part of the task definition. Please
// note we use a go string and do not load an external file, since we want this
// to be *part of the compiled executable*. If this sat in another file that
// was loaded at runtime, it would not be burned into the build, which would be
// bad for the following two reasons:
//  1) we could no longer distribute a single binary file that didn't require
//     installation/extraction
//  2) the payload schema is specific to the version of the code, therefore
//     should be versioned directly with the code and *frozen on build*.
//
// Run `generic-worker show-payload-schema` to output this schema to standard
// out.
func taskPayloadSchema() string {
	return `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "additionalProperties": false,
  "definitions": {
    "content": {
      "oneOf": [
        {
          "additionalProperties": false,
          "description": "Requires scope ` + "`" + `queue:get-artifact:\u003cartifact-name\u003e` + "`" + `",
          "properties": {
            "artifact": {
              "maxLength": 1024,
              "type": "string"
            },
            "taskId": {
              "pattern": "^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$",
              "type": "string"
            }
          },
          "required": [
            "taskId",
            "artifact"
          ],
          "title": "Artifact Content",
          "type": "object"
        },
        {
          "additionalProperties": false,
          "description": "URL to download content from",
          "properties": {
            "url": {
              "description": "URL to download content from",
              "format": "uri",
              "title": "URL",
              "type": "string"
            }
          },
          "required": [
            "url"
          ],
          "title": "URL Content",
          "type": "object"
        }
      ]
    },
    "fileMount": {
      "additionalProperties": false,
      "properties": {
        "content": {
          "$ref": "#/definitions/content",
          "description": "Content of the file to be mounted"
        },
        "file": {
          "description": "The filesystem location to mount the file",
          "title": "File",
          "type": "string"
        }
      },
      "required": [
        "file",
        "content"
      ],
      "title": "File Mount",
      "type": "object"
    },
    "mount": {
      "oneOf": [
        {
          "$ref": "#/definitions/fileMount"
        },
        {
          "$ref": "#/definitions/writableDirectoryCache"
        },
        {
          "$ref": "#/definitions/readOnlyDirectory"
        }
      ],
      "title": "Mount"
    },
    "readOnlyDirectory": {
      "additionalProperties": false,
      "properties": {
        "content": {
          "$ref": "#/definitions/content",
          "description": "Contents of read only directory.",
          "title": "Content"
        },
        "directory": {
          "description": "The filesystem location to mount the directory volume",
          "title": "Directory",
          "type": "string"
        },
        "format": {
          "description": "Archive format of content for read only directory",
          "enum": [
            "rar",
            "tar.bz2",
            "tar.gz",
            "zip"
          ],
          "title": "Format",
          "type": "string"
        }
      },
      "required": [
        "directory",
        "content",
        "format"
      ],
      "title": "Read Only Directory",
      "type": "object"
    },
    "writableDirectoryCache": {
      "additionalProperties": false,
      "dependencies": {
        "content": [
          "format"
        ],
        "format": [
          "content"
        ]
      },
      "properties": {
        "cacheName": {
          "description": "Implies a read/write cache directory volume. A unique name for the cache volume. Requires scope ` + "`" + `generic-worker:cache:\u003ccache-name\u003e` + "`" + `. Note if this cache is loaded from an artifact, you will also require scope ` + "`" + `queue:get-artifact:\u003cartifact-name\u003e` + "`" + ` to use this cache.",
          "title": "Cache Name",
          "type": "string"
        },
        "content": {
          "$ref": "#/definitions/content",
          "description": "Optional content to be preloaded when initially creating the cache (if set, ` + "`" + `format` + "`" + ` must also be provided).",
          "title": "Content"
        },
        "directory": {
          "description": "The filesystem location to mount the directory volume",
          "title": "Directory Volume",
          "type": "string"
        },
        "format": {
          "description": "Archive format of the preloaded content (if ` + "`" + `content` + "`" + ` provided).",
          "enum": [
            "rar",
            "tar.bz2",
            "tar.gz",
            "zip"
          ],
          "title": "Format",
          "type": "string"
        }
      },
      "required": [
        "directory",
        "cacheName"
      ],
      "title": "Writable Directory Cache",
      "type": "object"
    }
  },
  "description": "This schema defines the structure of the ` + "`" + `payload` + "`" + ` property referred to in a\nTaskcluster Task definition.",
  "id": "http://schemas.taskcluster.net/generic-worker/v1/payload.json#",
  "properties": {
    "artifacts": {
      "description": "Artifacts to be published. For example:\n` + "`" + `{ \"type\": \"file\", \"path\": \"builds\\\\firefox.exe\", \"expires\": \"2015-08-19T17:30:00.000Z\" }` + "`" + `",
      "items": {
        "additionalProperties": false,
        "properties": {
          "contentType": {
            "description": "Explicitly set the value of the HTTP ` + "`" + `Content-Type` + "`" + ` response header when the artifact(s)\nis/are served over HTTP(S). If not provided (this property is optional) the worker will\nguess the content type of artifacts based on the filename extension of the file storing\nthe artifact content. It does this by looking at the system filename-to-mimetype mappings\ndefined in the Windows registry. Note, setting ` + "`" + `contentType` + "`" + ` on a directory artifact will\napply the same contentType to all files contained in the directory.\n\nSee [mime.TypeByExtension](https://godoc.org/mime#TypeByExtension).",
            "title": "Content-Type header when serving artifact over HTTP",
            "type": "string"
          },
          "expires": {
            "description": "Date when artifact should expire must be in the future, no earlier than task deadline, but\nno later than task expiry. If not set, defaults to task expiry.",
            "format": "date-time",
            "title": "Expiry date and time",
            "type": "string"
          },
          "name": {
            "description": "Name of the artifact, as it will be published. If not set, ` + "`" + `path` + "`" + ` will be used.\nConventionally (although not enforced) path elements are forward slash separated. Example:\n` + "`" + `public/build/a/house` + "`" + `. Note, no scopes are required to read artifacts beginning ` + "`" + `public/` + "`" + `.\nArtifact names not beginning ` + "`" + `public/` + "`" + ` are scope-protected (caller requires scopes to\ndownload the artifact). See the Queue documentation for more information.",
            "title": "Name of the artifact",
            "type": "string"
          },
          "path": {
            "description": "Relative path of the file/directory from the task directory. Note this is not an absolute\npath as is typically used in docker-worker, since the absolute task directory name is not\nknown when the task is submitted. Example: ` + "`" + `dist\\regedit.exe` + "`" + `. It doesn't matter if\nforward slashes or backslashes are used.",
            "title": "Artifact location",
            "type": "string"
          },
          "type": {
            "description": "Artifacts can be either an individual ` + "`" + `file` + "`" + ` or a ` + "`" + `directory` + "`" + ` containing\npotentially multiple files with recursively included subdirectories.",
            "enum": [
              "file",
              "directory"
            ],
            "title": "Artifact upload type.",
            "type": "string"
          }
        },
        "required": [
          "type",
          "path"
        ],
        "title": "Artifact",
        "type": "object"
      },
      "title": "Artifacts to be published",
      "type": "array"
    },
    "command": {
      "description": "One entry per command (consider each entry to be interpreted as a full line of\na Windows™ .bat file). For example:\n` + "`" + `[\"set\", \"echo hello world \u003e hello_world.txt\", \"set GOPATH=C:\\\\Go\"]` + "`" + `.",
      "items": {
        "type": "string"
      },
      "minItems": 1,
      "title": "Commands to run",
      "type": "array"
    },
    "env": {
      "additionalProperties": {
        "type": "string"
      },
      "description": "Env vars must be string to __string__ mappings (not number or boolean). For example:\n` + "`" + `` + "`" + `` + "`" + `\n{\n  \"PATH\": \"C:\\\\Windows\\\\system32;C:\\\\Windows\",\n  \"GOOS\": \"windows\",\n  \"FOO_ENABLE\": \"true\",\n  \"BAR_TOTAL\": \"3\"\n}\n` + "`" + `` + "`" + `` + "`" + `",
      "title": "Env vars",
      "type": "object"
    },
    "features": {
      "additionalProperties": false,
      "description": "Feature flags enable additional functionality.",
      "properties": {
        "chainOfTrust": {
          "description": "An artifact named ` + "`" + `public/chainOfTrust.json.asc` + "`" + ` should be generated\nwhich will include information for downstream tasks to build a level\nof trust for the artifacts produced by the task and the environment\nit ran in.",
          "title": "Enable generation of a openpgp signed Chain of Trust artifact",
          "type": "boolean"
        }
      },
      "title": "Feature flags",
      "type": "object"
    },
    "maxRunTime": {
      "description": "Maximum time the task container can run in seconds",
      "maximum": 86400,
      "minimum": 1,
      "multipleOf": 1,
      "title": "Maximum run time in seconds",
      "type": "integer"
    },
    "mounts": {
      "description": "Directories and/or files to be mounted",
      "items": {
        "$ref": "#/definitions/mount",
        "title": "Mount"
      },
      "type": "array"
    },
    "osGroups": {
      "description": "A list of OS Groups that the task user should be a member of. Requires\nscope ` + "`" + `generic-worker:os-group:\u003cos-group\u003e` + "`" + ` for each group listed.",
      "items": {
        "type": "string"
      },
      "title": "OS Groups",
      "type": "array"
    },
    "supersederUrl": {
      "description": "URL of a service that can indicate tasks superseding this one; the current ` + "`" + `taskId` + "`" + `\nwill be appended as a query argument ` + "`" + `taskId` + "`" + `. The service should return an object with\na ` + "`" + `supersedes` + "`" + ` key containing a list of ` + "`" + `taskId` + "`" + `s, including the supplied ` + "`" + `taskId` + "`" + `. The\ntasks should be ordered such that each task supersedes all tasks appearing later in the\nlist.  See\n[superseding](https://docs.taskcluster.net/reference/platform/taskcluster-queue/docs/superseding)\nfor more detail.",
      "format": "uri",
      "title": "Superseder URL",
      "type": "string"
    }
  },
  "required": [
    "command",
    "maxRunTime"
  ],
  "title": "Generic worker payload",
  "type": "object"
}`
}
