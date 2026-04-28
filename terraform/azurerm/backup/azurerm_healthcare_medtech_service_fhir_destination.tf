resource "azurerm_healthcare_medtech_service_fhir_destination" "pike_gen" {
  name                                 = "examplemtdes"
  location                             = "east us"
  medtech_service_id                   = "pike"
  destination_fhir_service_id          = "pike"
  destination_identity_resolution_type = "Create"

  destination_fhir_mapping_json = jsonencode({
    "templateType" : "CollectionFhirTemplate",
    "template" : [
      {
        "templateType" : "CodeValueFhir",
        "template" : {
          "codes" : [
            {
              "code" : "8867-4",
              "system" : "http://loinc.org",
              "display" : "Heart rate"
            }
          ],
          "periodInterval" : 60,
          "typeName" : "heartrate",
          "value" : {
            "defaultPeriod" : 5000,
            "unit" : "count/min",
            "valueName" : "hr",
            "valueType" : "SampledData"
          }
        }
      }
    ]
  })
}
