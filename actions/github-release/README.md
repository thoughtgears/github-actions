# Thoughtgears GitHub Release

This action releases the internal actions to GitHub, it only needs the action name, assuming the action is in the 
same repository and its located under `actions/action-name`. It will extract the version from the tag and generate
the release notes from the CHANGELOG.md file based on the version tagging.

## Inputs

### `github_token`

**Required** The GitHub token.

### `action`

**Required** The name of the action.

### `latest`

Set to `true` to create a latest release. Default is `true`.

### `prerelease`

Set to `true` to create a prerelease. Default is `false`.

## Example usage

```yaml
uses: thoughtgears/github-release@v1
with:
  github_token: ${{ secrets.GITHUB_TOKEN }}
  action: 'action-name'
  latest: true
  prerelease: false
```

## Author

- [Thoughtgears](https://www.thoughtgears.co.uk)
