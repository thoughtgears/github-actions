# Thoughtgears GitHub Release

This action releases interesting things to GitHub.

## Inputs

### `github_token`

**Required** The GitHub token.

### `release_name`

**Required** The release name.

### `release_body`

The release body. Default is `" "`.

### `latest`

Set to `true` to create a latest release. Default is `true`.

### `prerelease`

Set to `true` to create a prerelease. Default is `false`.

## Example usage

```yaml
uses: thoughtgears/github-release@v1
with:
  github_token: ${{ secrets.GITHUB_TOKEN }}
  release_name: 'v1.0.0'
  release_body: 'This is the body of the release.'
  latest: true
  prerelease: false
```

## Author

- [Thoughtgears](https://www.thoughtgears.co.uk)
