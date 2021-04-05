function initTracker(baseUri) {
  let tracker = {}

  tracker.trackCustom = function (e, d) {
    fetch(`${baseUri}/v1/tracking/custom?cts=${Date.now()}&e=${e}&d=${d}`)
      .then(r => r.text());
      // .then(data => console.log(data));
  }

  tracker.trackPageView = function (d) {
    fetch(`${baseUri}/v1/tracking/pageview?cts=${Date.now()}&d=${d}`)
      .then(r => r.text());
      // .then(data => console.log(data));
  }
  
  tracker.trackClick = function (d) {
    fetch(`${baseUri}/v1/tracking/click?cts=${Date.now()}&d=${d}&uid=testuid`)
      .then(r => r.text());
      // .then(data => console.log(data));
  }

  return tracker
}
