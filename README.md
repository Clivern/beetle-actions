<p align="center">
    <img alt="Pushover Actions Logo" src="https://raw.githubusercontent.com/Clivern/beetle-actions/master/assets/images/logo.png" height="130" />
    <h3 align="center">Pushover Actions</h3>
    <p align="center">Push notifications for github repository changes through beetle</p>
    <p align="center">
        <a href="https://travis-ci.com/Clivern/beetle-actions"><img src="https://travis-ci.com/Clivern/beetle-actions.svg?branch=master"></a>
            <a href="https://github.com/Clivern/beetle-actions/actions"><img src="https://github.com/Clivern/beetle-actions/workflows/beetle-actions/badge.svg"></a>
        <a href="https://github.com/Clivern/beetle-actions/releases"><img src="https://img.shields.io/badge/Version-0.0.6-red.svg"></a>
         <a href="https://hub.docker.com/r/clivern/beetle-actions"><img src="https://img.shields.io/badge/Docker-Latest-green"></a>
        <a href="https://github.com/Clivern/beetle-actions/blob/master/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg"></a>
    </p>
</p>


## Documentation

#### Setup the Github Action:

1. First create [beetle application](https://beetle.net/) and get a token and user key.

2. Create github action by adding the following to your `workflow.yml`.

```yml
name: workflow_name

on:
    # may vary based on modules enabled

jobs:
  beetle-actions:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2 # required to clone your code
      - name: beetle-actions
        uses: clivern/beetle-actions@master
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          PUSHOVER_TOKEN: ${{ secrets.PUSHOVER_TOKEN }}
          PUSHOVER_USER: ${{ secrets.PUSHOVER_USER }}
```

3. Add beetle token as `PUSHOVER_TOKEN` and user key as `PUSHOVER_USER` in *github settings > secrets page*.

4. [Check this workflow as an example](https://github.com/Clivern/beetle-actions/blob/master/.github/workflows/beetle-actions.yml)

```yml
name: beetle-actions

on:
  issues:
    types: [opened, edited]

jobs:
  beetle-actions:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2 # required to clone your code
      - name: beetle-actions
        uses: clivern/beetle-actions@master
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          PUSHOVER_TOKEN: ${{ secrets.PUSHOVER_TOKEN }}
          PUSHOVER_USER: ${{ secrets.PUSHOVER_USER }}
```


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Pushover Actions is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/beetle-actions/releases) for changelogs for each release version of Pushover Actions. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/beetle-actions/issues


## Security Issues

If you discover a security vulnerability within Pushover Actions, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2021, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Pushover-Actions** is authored and maintained by [@Clivern](https://github.com/clivern).
