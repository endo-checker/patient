@description('The name of the API')
param apiName string

@description('The display name to use in API manager')
param displayName string

@description('The name of backend service for this API')
param backendName string

@description('The API specification in openapi format')
param apiSpec string

@description('Name of the API manager')
param apiManagerName string = 'apim-endochecker-dev'

@description('Fully qualified domain name of the app')
param appFdqn string

// get reference to API manager
resource apiManager 'Microsoft.ApiManagement/service@2021-08-01' existing = {
  name: apiManagerName
}

// create backend for service
resource backend 'Microsoft.ApiManagement/service/backends@2021-08-01' = {
  name: backendName
  parent: apiManager
  properties: {
    url: 'https://${appFdqn}'
    protocol: 'http'
    tls: {
      validateCertificateChain: true
      validateCertificateName: true
    }
  }
}

// update API from swagger
resource api 'Microsoft.ApiManagement/service/apis@2021-08-01' = {
  name: apiName
  parent: apiManager
  properties: {
    displayName: displayName
    apiRevision: '1'
    // apiVersion: 'string'
    isCurrent: true
    path: '/${apiName}'
    type: 'http'
    protocols: [
      'https'
    ]
    subscriptionRequired: false
    format: 'swagger-json'
    value: apiSpec
  }
}

// set policies
resource policies 'Microsoft.ApiManagement/service/apis/policies@2021-08-01' = {
  name: 'policy'
  parent: api
  properties: {
    format: 'xml'
    value: loadTextContent('api-policies.xml')
  }
}
