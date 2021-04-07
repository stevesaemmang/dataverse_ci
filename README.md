# dataverse_ci
 tiny CLI tool for exporting/importing dataverse solution files.

## Usage

### Getting solution information
```
docker run stevesaemmang/dataverse_ci solutionlist 
--clientId 00000000-0000-0000-0000-000000000000 
--tenantId 00000000-0000-0000-0000-000000000000 
--clientSecret abcdefg1234567 
--environment https://mypowerapp.crm4.dynamics.com

```
### Set solution version
```
docker run stevesaemmang/dataverse_ci setsolutionversion 
--clientId 00000000-0000-0000-0000-000000000000 
--tenantId 00000000-0000-0000-0000-000000000000 
--clientSecret abcdefg1234567 
--environment https://mypowerapp.crm4.dynamics.com 
--solutionid 00000000-0000-0000-0000-000000000000 
--version 2021.04.06.1210

```
### Export unmanaged solution by name
```
docker run -v `pwd`:/output stevesaemmang/dataverse_ci exportsolution 
--clientId 00000000-0000-0000-0000-000000000000 
--tenantId 00000000-0000-0000-0000-000000000000 
--clientSecret abcdefg1234567 
--environment https://mypowerapp.crm4.dynamics.com 
--solutionname MySolution

```
### Export managed solution by name
```
docker run -v `pwd`:/output stevesaemmang/dataverse_ci exportsolution 
--clientId 00000000-0000-0000-0000-000000000000 
--tenantId 00000000-0000-0000-0000-000000000000 
--clientSecret abcdefg1234567 
--environment https://mypowerapp.crm4.dynamics.com 
--solutionname MySolution 
--managed

```
### Import solution
```
docker run -v `pwd`:/input stevesaemmang/dataverse_ci importsolution 
--clientId 00000000-0000-0000-0000-000000000000 
--tenantId 00000000-0000-0000-0000-000000000000 
--clientSecret abcdefg1234567 
--environment https://mypowerapp.crm4.dynamics.com 
--solutionfile input/Test_managed.zip 
--overwriteunmanaged 
--publishworkflows

```

## Use in GitHub Actions
### Deploy solution file
```
jobs:
  deploy_powerapp:
    name: DeployPowerAppSolution

    runs-on: ubuntu-latest
    container: stevesaemmang/dataverse_ci
    steps:
    - uses: actions/checkout@v2
    - name: Deploy
      run: dataverse_ci importsolution --clientId 00000000-0000-0000-0000-000000000000 --tenantId 00000000-0000-0000-0000-000000000000 --clientSecret abcdefg1234567  --environment https://mypowerapp.crm4.dynamics.com --solutionfile Test_managed.zip --overwriteunmanaged --publishworkflows
      working-directory: /.
```