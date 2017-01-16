# Go FHIR [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/dshills/gofhir)

## Overview

> Fast Healthcare Interoperability Resources (FHIR, pronounced "fire") is a draft standard describing data formats and elements (known as "resources") and an Application Programming Interface (API) for exchanging Electronic health records. The standard was created by the Health Level Seven International (HL7) health-care standards organization.

> FHIR builds on previous data format standards from HL7, like HL7 version 2.x and HL7 version 3.x. But it is easier to implement because it uses a modern web-based suite of API technology, including a HTTP-based RESTful protocol, HTML and Cascading Style Sheets for user interface integration, a choice of JSON or XML for data representation, and Atom for results. One of its goals is to facilitate interoperation between legacy health care systems, to make it easy to provide health care information to health care providers and individuals on a wide variety of devices from computers to tablets to cell phones, and to allow third-party application developers to provide medical applications which can be easily integrated into existing systems.

> FHIR provides an alternative to document-centric approaches by directly exposing discrete data elements as services. For example, basic elements of healthcare like patients, admissions, diagnostic reports and medications can each be retrieved and manipulated via their own resource URLs. FHIR was supported at an American Medical Informatics Association meeting by companies like Cerner which value its open and extensible nature. - https://en.wikipedia.org/wiki/Fast_Healthcare_Interoperability_Resources

Go FHIR is a Go language client for healthcare FHIR. Specifications, testing and examples where pulled from https://open.epic.com and the HL7 FHIR documentation https://www.hl7.org/fhir/documentation.html.

## Supported Queries
- allergy
- careplan
- condition
- device
- diagnostic report
- document
- family history
- immunization
- observation
- patient
- procedure

## Instalation
```sh
go get -u github.com/dshills/gofhir
```

## Usage
```Go
c := New("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/")
pat, err := c.PatientSearch("family=Argonaut&given=Jason")
```
