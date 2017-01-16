package fhir

import "testing"

const pid = "Tbt3KuCY0B5PSrJvCu2j-PlK.aiHsu2xUjUM8bWpetXoB"
const ordercode = "8310-5"

func TestDevice(t *testing.T) {
	c := New("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/")
	data, err := c.GetDevice(pid)
	if err != nil {
		t.Error(err)
	}
	if data.Total == 0 {
		t.Error("Expected > 0 got 0")
	}
}

func TestPatSearch(t *testing.T) {
	c := New("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/")
	data, err := c.PatientSearch("family=Argonaut&given=Jason")
	if err != nil {
		t.Fatal(err)
	}
	if data.Total == 0 {
		t.Error("Expected > 0 got 0")
	}
}

func TestPatient(t *testing.T) {
	c := New("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/")
	data, err := c.GetPatient(pid)
	if err != nil {
		t.Fatal(err)
	}
	if len(data.Name) == 0 {
		t.Error("Expected > 0 got 0")
	}
}

func TestDocument(t *testing.T) {
	c := New("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/")
	data, err := c.GetDocumentReference(pid)
	if err != nil {
		t.Fatal(err)
	}
	if data.Total == 0 {
		t.Error("Expected > 0 got 0")
	}
}

func TestCondition(t *testing.T) {
	c := New("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/")
	data, err := c.GetCondition(pid)
	if err != nil {
		t.Fatal(err)
	}
	if data.Total == 0 {
		t.Error("Expected > 0 got 0")
	}
}

func TestProcedure(t *testing.T) {
	c := New("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/")
	data, err := c.GetProcedure(pid)
	if err != nil {
		t.Fatal(err)
	}
	if data.Total == 0 {
		t.Error("Expected > 0 got 0")
	}
}

func TestMedication(t *testing.T) {
	c := New("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/")
	data, err := c.GetMedication(pid)
	if err != nil {
		t.Fatal(err)
	}
	if data.Total == 0 {
		t.Error("Expected > 0 got 0")
	}
}

func TestObservation(t *testing.T) {
	c := New("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/")
	data, err := c.GetObservation(pid, ordercode)
	if err != nil {
		t.Fatal(err)
	}
	if data.Total == 0 {
		t.Error("Expected > 0 got 0")
	}
}

func TestImmunization(t *testing.T) {
	c := New("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/")
	data, err := c.GetImmunization(pid)
	if err != nil {
		t.Fatal(err)
	}
	if data.Total == 0 {
		t.Error("Expected > 0 got 0")
	}
}

func TestAllergy(t *testing.T) {
	c := New("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/")
	data, err := c.GetAllergyIntolerence(pid)
	if err != nil {
		t.Fatal(err)
	}
	if data.Total == 0 {
		t.Error("Expected > 0 got 0")
	}
}

func TestFamilyHx(t *testing.T) {
	c := New("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/")
	data, err := c.GetFamilyMemberHistory(pid)
	if err != nil {
		t.Fatal(err)
	}
	if data.Total == 0 {
		t.Error("Expected > 0 got 0")
	}
}
