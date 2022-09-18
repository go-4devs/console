package value

import (
	"fmt"
	"time"

	"gitoa.ru/go-4devs/console/input/errs"
)

var (
	_ Value = Read{}
	_ Value = Slice{}
)

var ErrWrongType = errs.ErrWrongType

type Read struct {
	ParseValue
}

func (r Read) String() string {
	sout, _ := r.ParseValue.ParseString()

	return sout
}

func (r Read) Int() int {
	iout, _ := r.ParseValue.ParseInt()

	return iout
}

func (r Read) Int64() int64 {
	iout, _ := r.ParseValue.ParseInt64()

	return iout
}

func (r Read) Uint() uint {
	uout, _ := r.ParseValue.ParseUint()

	return uout
}

func (r Read) Uint64() uint64 {
	uout, _ := r.ParseValue.ParseUint64()

	return uout
}

func (r Read) Float64() float64 {
	fout, _ := r.ParseValue.ParseFloat64()

	return fout
}

func (r Read) Bool() bool {
	bout, _ := r.ParseValue.ParseBool()

	return bout
}

func (r Read) Duration() time.Duration {
	dout, _ := r.ParseValue.ParseDuration()

	return dout
}

func (r Read) Time() time.Time {
	tout, _ := r.ParseValue.ParseTime()

	return tout
}

func (r Read) Strings() []string {
	return []string{r.String()}
}

func (r Read) Ints() []int {
	return []int{r.Int()}
}

func (r Read) Int64s() []int64 {
	return []int64{r.Int64()}
}

func (r Read) Uints() []uint {
	return []uint{r.Uint()}
}

func (r Read) Uint64s() []uint64 {
	return []uint64{r.Uint64()}
}

func (r Read) Float64s() []float64 {
	return []float64{r.Float64()}
}

func (r Read) Bools() []bool {
	return []bool{r.Bool()}
}

func (r Read) Durations() []time.Duration {
	return []time.Duration{r.Duration()}
}

func (r Read) Times() []time.Time {
	return []time.Time{r.Time()}
}

type Slice struct {
	SliceValue
}

func (s Slice) String() string {
	return ""
}

func (s Slice) Int() int {
	return 0
}

func (s Slice) Int64() int64 {
	return 0
}

func (s Slice) Uint() uint {
	return 0
}

func (s Slice) Uint64() uint64 {
	return 0
}

func (s Slice) Float64() float64 {
	return 0
}

func (s Slice) Bool() bool {
	return false
}

func (s Slice) Duration() time.Duration {
	return 0
}

func (s Slice) Time() time.Time {
	return time.Time{}
}

func (s Slice) wrongType() error {
	return fmt.Errorf("%w: for %T", ErrWrongType, s.SliceValue)
}

func (s Slice) ParseString() (string, error) {
	return "", s.wrongType()
}

func (s Slice) ParseInt() (int, error) {
	return 0, s.wrongType()
}

func (s Slice) ParseInt64() (int64, error) {
	return 0, s.wrongType()
}

func (s Slice) ParseUint() (uint, error) {
	return 0, s.wrongType()
}

func (s Slice) ParseUint64() (uint64, error) {
	return 0, s.wrongType()
}

func (s Slice) ParseFloat64() (float64, error) {
	return 0, s.wrongType()
}

func (s Slice) ParseBool() (bool, error) {
	return false, s.wrongType()
}

func (s Slice) ParseDuration() (time.Duration, error) {
	return 0, s.wrongType()
}

func (s Slice) ParseTime() (time.Time, error) {
	return time.Time{}, s.wrongType()
}
