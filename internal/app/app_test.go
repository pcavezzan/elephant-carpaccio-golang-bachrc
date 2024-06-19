package app

import (
	"testing"

	"github.com/bachrc/test-stonal/internal/state"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {

	t.Run("app should return utah in its states", func(t *testing.T) {
		app := New()

		states := app.ListStates()

		assert.Equal(t, []state.State{{
			Name: "UTAH",
		}}, states)
	})
}