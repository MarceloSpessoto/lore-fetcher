package fetcher

import (
  "testing"
)

func TestParsePatchTag(t *testing.T){
  got := parsePatchTag("http://lore.kernel.org/all/171949586279.9146.13526422228365247246.rtt-probe@aws-us-west-2-korg-lkml-1.web.codeaurora.org/")
  want := "171949586279.9146.13526422228365247246.rtt"
  if(want != got){
    t.Errorf("got %s, wanted %s", got, want)
  }
}


