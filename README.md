# Kafka to HipChat

[![Build Status](https://travis-ci.org/chrisjoyce911/slacktohip.svg?branch=Ready-for-Kafka)](https://travis-ci.org/chrisjoyce911/slacktohip)
[![Coverage Status](https://coveralls.io/repos/github/chrisjoyce911/slacktohip/badge.svg?branch=Ready-for-Kafka)](https://coveralls.io/github/chrisjoyce911/slacktohip?branch=Ready-for-Kafka)

`kafkatohip` uses a Kafka consumer to process feedback messages, and posts them to HipChat

```bash
kafkatohip
```

## Config

When first run a configuration file will be prodiced with examples

### Redirection

If the message contains matching text the bot will then redirect the message to a new channel

### Background

Backgrounds can be set based on matching text in the message
