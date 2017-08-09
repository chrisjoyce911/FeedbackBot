# Slack to HipChat

[![Build Status](https://travis-ci.org/chrisjoyce911/slacktohip.svg?branch=master)](https://travis-ci.org/chrisjoyce911/slacktohip)
[![Coverage Status](https://coveralls.io/repos/github/chrisjoyce911/slacktohip/badge.svg?branch=master)](https://coveralls.io/github/chrisjoyce911/slacktohip?branch=master)
[![GitHub version](https://badge.fury.io/gh/chrisjoyce911%2Fslacktohip.svg)](https://badge.fury.io/gh/chrisjoyce911%2Fslacktohip)
![chriskoyce911/slacktohip](https://reposs.herokuapp.com/?path=chrisjoyce911/slacktohip)
[![Code Climate](http://img.shields.io/codeclimate/github/badges/badgerbadgerbadger.svg?style=flat-square)](https://codeclimate.com/github/chrisjoyce911/slacktohip)
[![Github Issues](http://githubbadges.herokuapp.com/chrisjoyce911/slacktohip/issues.svg?style=flat-square)](https://github.com/chrisjoyce911/slacktohip/issues)
[![Pending Pull-Requests](http://githubbadges.herokuapp.com/chrisjoyce911/slacktohip/pulls.svg?style=flat-square)](https://github.com/chrisjoyce911/slacktohip/pulls)

`slacktohip` is a very simple single way cross messageing bot.

```bash
slacktohip
```

## Outline

Bot will review any message in any channge that it has been invited into, should it be in a channel that it has a record of.

If the channel is one that the bot has been configured to forared messages for it will then run redirection tests

### Redirection

If the message contains matching text the bot will then redirect the message to a new channel

### Background

Backgrounds can be set based on matching text in the message
