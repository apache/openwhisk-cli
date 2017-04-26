# Contributing to go-whisk-cli

## Proposing new features

If you would like to implement a new feature, please [raise an issue](https://github.ibm.com/BlueMix-Fabric/go-whisk-cli) before sending a pull request so the feature can be discussed.
This is to avoid you spending your valuable time working on a feature that the project developers are not willing to accept into the code base.

## Fixing bugs

If you would like to fix a bug, please [raise an issue](https://github.ibm.com/BlueMix-Fabric/go-whisk-cli) before sending a pull request so it can be discussed.
If the fix is trivial or non controversial then this is not usually necessary.

## Merge approval

The project maintainers use LGTM (Looks Good To Me) in comments on the code review to
indicate acceptance. A change requires LGTMs from two of the maintainers of each
component affected.

## Communication
Please use [Slack channel #whisk-users](https://cloudplatform.slack.com/messages/whisk_cli).
## Setup
Project was written with `Go v1.5`.  It has a dependency on [go-whisk](https://github.ibm.com/BlueMix-Fabric/go-whisk), which has to be manually resolved (because of GHE).

## Testing

This repository needs unit tests.

Please provide information that helps the developer test any changes they make before submitting.

Should pass the cli integration test defined in the [main whisk project](https://github.rtp.raleigh.ibm.com/whisk-development/openwhisk/blob/master/tests/src/common/WskCli.java).

## Coding style guidelines

Use idomatic go.  (try to) Document exported functions.
