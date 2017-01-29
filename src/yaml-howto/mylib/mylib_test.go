package mylib

import (
	"gopkg.in/yaml.v2"
	"log"
	"testing"
)

func TestValidateConfig_0(t *testing.T) {

	var data = `
info:  "blue"

data:
  source: http://intra
  destination:  /tmp 

run:
  -  id:  A1
     exe:  "run.a1"
     output:  "output.A1"
  
  -  id:  A2
     exe:  "run.a2"
     output:  "output.A2"
     `
	expected := true

	conf := Config{}

	err := yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		log.Fatal("problem...")
		log.Fatal(err)
	}

	ans := ValidateConfig(&conf)
	if expected != ans {
		t.Error("Exected", expected, "got", ans)
	}

}

func TestValidateConfig_1(t *testing.T) {

	var data = `


data:
  source: http://intra
  destination:  /tmp 

run:
  -  id:  A1
     exe:  "run.a1"
     output:  "output.A1"
  
  -  id:  A2
     exe:  "run.a2"
     output:  "output.A2"
     `
	expected := false

	conf := Config{}

	err := yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		log.Fatal("problem...")
		log.Fatal(err)
	}

	ans := ValidateConfig(&conf)
	if expected != ans {
		t.Error("Exected", expected, "got", ans)
	}

}

func TestAAA(t *testing.T) {
	expected := true
	ans := AAA()
	if expected != ans {
		t.Error("Expected", expected, "got", ans)
	}
}
