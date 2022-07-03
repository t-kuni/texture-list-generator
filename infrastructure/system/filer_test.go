package system_test

import (
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
	"github.com/t-kuni/go-cli-app-skeleton/infrastructure/system"
	"testing"
)

func TestFiler(t *testing.T) {
	t.Run("FindFilePaths", func(t *testing.T) {
		t.Skip()

		//
		// Prepare
		//
		testee, err := system.NewFiler(do.New())
		if err != nil {
			t.Fatal(err)
		}

		actual, err := testee.FindFilePaths([]string{".go"})
		if err != nil {
			t.Fatal(err)
		}

		assert.NotEmpty(t, actual)
	})
}
