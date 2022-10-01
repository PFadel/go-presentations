package withinterfaces_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/PFadel/go-presentations/gopherconbr2022/withinterfaces"
	"github.com/PFadel/go-presentations/gopherconbr2022/withinterfaces/mocks"
)

type mockCalculator struct {
	mockDistanceReturn float64
}

func (m *mockCalculator) Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	return m.mockDistanceReturn
}

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

func TestWithBasicMock(t *testing.T) {
	calc := mockCalculator{1.0}
	deli := withinterfaces.NewDeliverable()

	winnipeg := coordinate{49.895077, -97.138451}
	regina := coordinate{50.445210, -104.618896}

	got := deli.WithinDeliveryDistance(winnipeg.lat, winnipeg.lng, regina.lat, regina.lng, &calc)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	calc.mockDistanceReturn = 1000000.0

	got = deli.WithinDeliveryDistance(winnipeg.lat, winnipeg.lng, regina.lat, regina.lng, &calc)
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestWithMockery(t *testing.T) {

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

func TestWithExpect(t *testing.T) {
	winnipeg := coordinate{49.895077, -97.138451}
	regina := coordinate{50.445210, -104.618896}

	calc := mocks.NewCalculator(t)
	calc.EXPECT().
		Distance(49.895077, -97.138451, 50.445210, -104.618896, "N").
		Run(func(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) {
			fmt.Println("parametros de latitude", lat1, lat2)
		}).
		Call.Return(
		func(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
			fmt.Println("parametros de longitude", lng1, lng2)
			return 10000.0
		})

	deli := withinterfaces.NewDeliverable()

	got := deli.WithinDeliveryDistance(winnipeg.lat, winnipeg.lng, regina.lat, regina.lng, calc)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
