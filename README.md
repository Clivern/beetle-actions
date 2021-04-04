<p align="center">
    <img src="https://raw.githubusercontent.com/clivern/beetle-actions/main/assets/img/gopher.png?v=0.1.0" width="180" />
    <h3 align="center">Beetle Actions</h3>
    <p align="center">Deploy to Kubernetes Cluster Using Beetle</p>
    <p align="center">
        <a href="https://travis-ci.com/Clivern/beetle-actions"><img src="https://travis-ci.com/Clivern/beetle-actions.svg?branch=main"></a>
            <a href="https://github.com/Clivern/beetle-actions/actions"><img src="https://github.com/Clivern/beetle-actions/workflows/beetle-actions/badge.svg"></a>
        <a href="https://github.com/Clivern/beetle-actions/releases"><img src="https://img.shields.io/badge/Version-0.1.0-red.svg"></a>
         <a href="https://hub.docker.com/r/clivern/beetle-actions"><img src="https://img.shields.io/badge/Docker-Latest-green"></a>
        <a href="https://github.com/Clivern/beetle-actions/blob/main/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg"></a>
    </p>
</p>


## Documentation

#### Setup the Github Action:

1. Create github action by adding the following to your `workflow.yml`.

```yaml
name: workflow_name

on:
  push:
    tags:
      - '*'

jobs:
  beetle-actions:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2 # required to clone your code
      - name: beetle-actions
        uses: clivern/beetle-actions@main
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          BEETLE_NUMERIC_TAG: true
          BEETLE_WATCH_DEPLOYMENT: true
          BEETLE_API_URL: ${{secrets.BEETLE_API_URL }}
          BEETLE_API_KEY: ${{secrets.BEETLE_API_KEY }}
          BEETLE_CLUSTER_NAME: ${{secrets.BEETLE_CLUSTER_NAME }}
          BEETLE_NAMESPACE_NAME: ${{secrets.BEETLE_NAMESPACE_NAME }}
          BEETLE_APP_ID: ${{secrets.BEETLE_APP_ID }}
          BEETLE_DEPLOYMENT_STRATEGY: ${{secrets.BEETLE_DEPLOYMENT_STRATEGY }}
          BEETLE_MAX_SURGE: ${{secrets.BEETLE_MAX_SURGE }}
          BEETLE_MAX_UNAVAILABLE: ${{secrets.BEETLE_MAX_UNAVAILABLE }}
```

2. Add all beetle action secrets in github settings > secrets page.


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Beetle Actions is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/beetle-actions/releases) for changelogs for each release version of Beetle Actions. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/beetle-actions/issues


## Security Issues

If you discover a security vulnerability within Beetle Actions, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2021, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Beetle-Actions** is authored and maintained by [@Clivern](https://github.com/clivern).
