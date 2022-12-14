// ----------------------------------------------------------------------------
// Deploys a Container App
// ----------------------------------------------------------------------------

@description('Location of the Container Apps environment')
param location string = resourceGroup().location

@description('Name of Container Apps Managed Environment')
param managedEnvName string

@description('Container App HTTP port')
param appName string

@description('Container App image name')
param imageName string = 'ghcr.io/endo-checker/${appName}:latest'

@description('Container App HTTP port')
param httpPort int = 8080

@description('Container App gRPC port')
param grpcPort int = 8080

@secure()
@description('Mongo DB URI')
param mongoUri string

@secure()
@description('GitHub container registry user')
param ghcrUser string

@secure()
@description('GitHub container registry personal access token')
param ghcrPat string

// get a reference to the container apps environment
resource managedEnv 'Microsoft.App/managedEnvironments@2022-01-01-preview' existing = {
  name: managedEnvName
}

resource containerApp 'Microsoft.App/containerApps@2022-01-01-preview' = {
  name: 'ca-${appName}'
  location: location
  properties: {
    managedEnvironmentId: managedEnv.id
    configuration: {
      activeRevisionsMode: 'single'
      dapr: {
        appId: appName
        appPort: grpcPort
        appProtocol: 'grpc'
        enabled: true
      }
      ingress: {
        external: true
        targetPort: httpPort
        allowInsecure: false
        traffic: [
          {
            latestRevision: true
            weight: 100
          }
        ]
      }
      registries: [
        {
          server: 'ghcr.io'
          username: ghcrUser
          passwordSecretRef: 'ghcr-pat'
        }
      ]
      secrets: [
        {
          name: 'ghcr-pat'
          value: ghcrPat
        }
        {
          name: 'mongo-uri'
          value: mongoUri
        }
      ]
    }
    template: {
      containers: [
        {
          name: appName
          image: imageName
          resources: {
            cpu: any('0.5')
            memory: '1Gi'
          }
          env: [
            {
              name: 'mongoUri'
              secretRef: 'mongo-uri'
            }
          ]
        }
      ]
      scale: {
        minReplicas: 1
        maxReplicas: 1
      }
    }
  }
}

// used to set app as 'backend' in API manager
output fdqn string = containerApp.properties.configuration.ingress.fqdn
