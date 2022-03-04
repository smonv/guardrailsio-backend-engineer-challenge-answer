package worker

import (
	"io/fs"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScanFile(t *testing.T) {
	dir, err := ioutil.TempDir(os.TempDir(), "beca")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	secretFile, err := ioutil.TempFile(dir, "secret")
	require.NoError(t, err)

	nonSecretFile, err := ioutil.TempFile(dir, "non-secret")
	require.NoError(t, err)

	err = ioutil.WriteFile(secretFile.Name(), []byte("public_keyasdf\nprivate_keyfdas"), fs.ModeAppend)
	require.NoError(t, err)

	err = ioutil.WriteFile(nonSecretFile.Name(), []byte("not so secret"), fs.ModeAppend)
	require.NoError(t, err)

	t.Run("correct find error in file have secret", func(t *testing.T) {
		findings, err := scanFile(secretFile.Name())
		assert.NoError(t, err)
		assert.NotNil(t, findings)
		assert.NotEmpty(t, findings)
		assert.Equal(t, len(findings), 2)
	})

	t.Run("correct find non error in not secret file", func(t *testing.T) {
		findings, err := scanFile(nonSecretFile.Name())
		assert.NoError(t, err)
		assert.Empty(t, findings)
	})

	t.Run("correct return err on non exist file", func(t *testing.T) {
		_, err := scanFile("asdf")
		assert.Error(t, err)
	})
}
