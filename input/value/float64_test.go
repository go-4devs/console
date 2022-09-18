package value_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"gitoa.ru/go-4devs/console/input/value"
)

func TestFloat64_Unmarshal(t *testing.T) {
	t.Parallel()

	f := value.Float64(math.Pi)

	var out float64

	require.NoError(t, f.Unmarshal(&out))
	require.Equal(t, math.Pi, out)
}

func TestFloat64_Any(t *testing.T) {
	t.Parallel()

	f := value.Float64(math.Pi)

	require.Equal(t, math.Pi, f.Any())
}

func TestFloat64s_Unmarshal(t *testing.T) {
	t.Parallel()

	f := value.Float64s{math.Pi, math.Sqrt2}

	var out []float64

	require.NoError(t, f.Unmarshal(&out))
	require.Equal(t, []float64{math.Pi, math.Sqrt2}, out)
}

func TestFloat64s_Any(t *testing.T) {
	t.Parallel()

	f := value.Float64s{math.Pi, math.Sqrt2}

	require.Equal(t, []float64{math.Pi, math.Sqrt2}, f.Any())
}
