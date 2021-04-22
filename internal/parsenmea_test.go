package internal

import "testing"

func TestDecode(t *testing.T) {
	t1 := Decode("5501.82010", 2)
	t2 := Decode("01142.81950", 3)

	if t1 != 55.030335 {
		t.Logf("Expected 55.030335 but got %f", t1)
		t.Fail()
	}
	if t2 != 11.713650 {
		t.Logf("Expected 11.713650 but got %f", t2)
		t.Fail()
	}
}

func TestParse(t *testing.T) {
	s1 := "$GPRMC,103240.00,A,5501.81896,N,01142.81966,E,1.640,,180421,,,A*76"

	lat, lon, err := Parse(s1)
	if err != nil {
		t.Logf("Error while parsing valid RMC data")
		t.Fail()
	}
	if lat != 55.030335 {
		t.Logf("Expected 55.030335 for latitude but got: %f", lat)
	}
	if lon != 11.713650 {
		t.Logf("Expected 11.713650 for longitude but got: %f", lon)
	}

	s2 := "$GPGSV,3,3,10,31,18,310,20,32,30,261,34*73"
	lat, lon, err = Parse(s2)
	if err == nil {
		t.Logf("Expecting failure but no error was returned")
		t.Fail()
	}
}

func TestParseInvalid(t *testing.T) {
	s1 := "$GPRMC,,,,,,,,,,,"

	_, _, err := Parse(s1)
	if err == nil {
		t.Logf("Error with invalid format expected but not thrown")
		t.Fail()
	}
}
