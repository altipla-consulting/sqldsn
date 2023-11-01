package sqldsn

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromURLCloudDev(t *testing.T) {
	require.Equal(t, FromURL("mysql://root:foo@clouddev/bar"), "root:foo@tcp(clouddev)/bar?parseTime=true")
}

func TestFromURLCloudSQLProxy(t *testing.T) {
	require.Equal(t, FromURL("mysql://foo-user:foo-pass@cloudsql/foo_db?socket=foo-project:europe-west1:foo-server"), "foo-user:foo-pass@cloudsql-mysql(foo-project:europe-west1:foo-server)/foo_db?parseTime=true")
}
