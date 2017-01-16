package fhir

import (
	"fmt"
	"testing"
)

const pid = "Tbt3KuCY0B5PSrJvCu2j-PlK.aiHsu2xUjUM8bWpetXoB"

func TestDevice(t *testing.T) {
	c := New("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/")
	_, err := c.GetDevice(pid)
	if err != nil {
		t.Error(err)
	}
}

func TestSearch(t *testing.T) {
	c := New("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/")
	pat, err := c.PatientSearch("family=Argonaut&given=Jason")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", pat)
}

/*
func TestOther(t *testing.T) {
	_, err = c.GetPatient(pid)
	if err != nil {
		t.Error(err)
	}
	_, err = c.GetDocumentReference(pid)
	if err != nil {
		t.Error(err)
	}
	_, err = c.GetCondition(pid)
	if err != nil {
		t.Error(err)
	}
	_, err = c.GetProcedure(pid)
	if err != nil {
		t.Error(err)
	}
	_, err = c.GetMedication(pid)
	if err != nil {
		t.Error(err)
	}
	_, err = c.GetObservation(pid)
	if err != nil {
		t.Error(err)
	}
	_, err = c.GetImmunization(pid)
	if err != nil {
		t.Error(err)
	}
	_, err = c.GetAllergyIntolerence(pid)
	if err != nil {
		t.Error(err)
	}
	_, err = c.GetFamilyMemberHistory(pid)
	if err != nil {
		t.Error(err)
	}
}
*/
