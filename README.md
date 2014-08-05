[![Build Status](https://travis-ci.org/agerlic/hipchat-notifications-to-slack.svg?branch=master)](https://travis-ci.org/agerlic/hipchat-notifications-to-slack)

## Forward Hipchat notifications to Slack

This tool reformat and forward notifications from HipChat to Slack

## Use Cases

* Test Slack without rewriting all your notifications
* Smooth transition from HipChat to Slack
* Handle services only compatible with Hipchat notifications eg (http://intercom.io)
* Don't use polling to transfer messages from Hipchat to Slack

## Deploy on Heroku
* Create your Slack Inbound Webhook : https://my.slack.com/services/new/incoming-webhook
  * Login to your Slack Account
  * Go to Integrations Tab
  * Select Incoming Webhooks
  * Choose a channel then click on Add Incoming Webhook
  * Save your incoming webhook url : https://XXX.slack.com/services/hooks/incoming-webhook?token=...
* Clone the repository
```git clone git@github.com:agerlic/hipchat-notifications-to-slack.git```
* Install Heroku toolbelt : https://toolbelt.heroku.com/
* Create & configure for Heroku
```
heroku create hipchat-to-slack -b https://github.com/kr/heroku-buildpack-go.git
heroku config:set SLACK_CHANNEL="#channel"
heroku config:set SLACK_URL="https://XXX.slack.com/services/hooks/incoming-webhook?token=TOKEN"
git push heroku master
```
* Create your Hipchat Webhook: https://www.hipchat.com/docs/apiv2/method/create_webhook
```
FORWARD_APP_URL=http://XXX.herokuapp.com HIPCHAT_ROOM=myroom HIPCHAT_ADMIN_TOKEN=mytoken ./hipchat-webhook.sh --create
```
* Test
```
HIPCHAT_ROOM=myroom HIPCHAT_ADMIN_TOKEN=mytoken ./hipchat-webhook.sh --test
```
* Enjoy!
