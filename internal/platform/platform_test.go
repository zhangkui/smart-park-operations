package platform

import (
	"testing"
	"time"
)

func TestIssueAndParseToken(t *testing.T) {
	raw, err := IssueToken("secret", "u-1", []string{"ops"}, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	c, err := ParseToken("secret", raw)
	if err != nil || c.UserID != "u-1" {
		t.Fatalf("claims=%+v err=%v", c, err)
	}
}
