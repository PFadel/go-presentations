package withinterfaces_test

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/PFadel/go-presentations/gopherconbr2022/withinterfaces"
	"github.com/PFadel/go-presentations/gopherconbr2022/withinterfaces/mocks"
)

type coordinate struct {
	lat float64
	lng float64
}

func TestNoMock(t *testing.T) {

	calc := withinterfaces.NewCalculator()
	deli := withinterfaces.NewDeliverable()

	winnipeg := coordinate{49.895077, -97.138451}
	regina := coordinate{50.445210, -104.618896}

	got := deli.WithinDeliveryDistance(winnipeg.lat, winnipeg.lng, regina.lat, regina.lng, calc)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestWithMock(t *testing.T) {

	calc := mocks.NewCalculator(t)
	deli := withinterfaces.NewDeliverable()

	calc.Mock.
		On("Distance", 49.895077, -97.138451, 50.445210, -104.618896, "N").
		Return(100000.0).
		Once()

	winnipeg := coordinate{49.895077, -97.138451}
	regina := coordinate{50.445210, -104.618896}

	got := deli.WithinDeliveryDistance(winnipeg.lat, winnipeg.lng, regina.lat, regina.lng, calc)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	calc.Mock.
		On("Distance", 49.895077, -97.138451, 50.445210, -104.618896, "N").
		Return(1.0).
		Once()

	got = deli.WithinDeliveryDistance(winnipeg.lat, winnipeg.lng, regina.lat, regina.lng, calc)
	want = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	calc.Mock.
		On("Distance", mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.Anything).
		Return(2000.0).
		Once()

	got = deli.WithinDeliveryDistance(winnipeg.lat, winnipeg.lng, regina.lat, regina.lng, calc)
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
