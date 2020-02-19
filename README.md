# go-github-actions-playground
Go github actions playground


## Actions
### Pull Request Master
Runs linting (golangci-lint) and tests against commits in a Pull Request  
[pull_request-master.yml](./.github/workflows/pull_request-master.yml)


### Push Master
Runs tests against commits pushed to master  
[push-master.yml](./.github/workflows/push-master.yml)

### Release
Creates a draft release for every pushed tag (vX.Y.Z)  
[release.yml](./.github/workflows/release.yml)  

#### Flow
Create tag  
`git tag v0.0.1`  

Push tag  
`git push origin --tags`  

Fill in details for the created draft release and publish  