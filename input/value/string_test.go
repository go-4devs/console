package value_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitoa.ru/go-4devs/console/input/value"
)

func TestStringUnmarshal(t *testing.T) {
	t.Parallel()

	st := value.New("test")
	sta := value.New([]string{"test1", "test2"})

	ac := ""
	require.NoError(t, st.Unmarshal(&ac))
	require.Equal(t, "test", ac)

	aca := []string{}
	require.NoError(t, sta.Unmarshal(&aca))
	require.Equal(t, []string{"test1", "test2"}, aca)

	require.ErrorIs(t, sta.Unmarshal(ac), value.ErrWrongType)
	require.ErrorIs(t, sta.Unmarshal(&ac), value.ErrWrongType)
	require.ErrorIs(t, st.Unmarshal(aca), value.ErrWrongType)
	require.ErrorIs(t, st.Unmarshal(&aca), value.ErrWrongType)
}
