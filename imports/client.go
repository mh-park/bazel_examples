package imports

import _ "github.com/minho.park/bazel_examples/imports/testutils"

type Client interface {
	Close() error
}
