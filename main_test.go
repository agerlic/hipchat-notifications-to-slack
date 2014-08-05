package main

import (
	"testing"
)

func TestReformatNewlines(t *testing.T) {
	source := "<br/><br />hello<br/>"
	expected := "\n\nhello\n"
	reformatedMessage := reformat(source)
	if expected != reformatedMessage {
		t.Errorf("reformatedMessage(): Expected %s, got %s", expected, reformatedMessage)
	}
}

func TestReformatLinks(t *testing.T) {
	expected := "<http://slack.com|slack>"
	source := "<a href=\"http://slack.com\">slack</a>"
	reformatedMessage := reformat(source)
	if expected != reformatedMessage {
		t.Errorf("reformatedMessage(): Expected %s, got %s", expected, reformatedMessage)
	}
}
