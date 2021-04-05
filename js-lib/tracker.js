function uuidv4() {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
    var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
    return v.toString(16);
  });
}

function getCookie(cname) {
  var name = cname + "=";
  var decodedCookie = decodeURIComponent(document.cookie);
  var ca = decodedCookie.split(';');
  for(var i = 0; i <ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) == ' ') {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return "";
}

function setCookie(cname, cvalue, exdays) {
  var d = new Date();
  d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
  var expires = "expires="+d.toUTCString();
  document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

function getLocalStorage(e) {
  return window.localStorage[e]
}


function initTracker(baseUri, cname, allowTracker) {
  let tracker = {}

  tracker.initUser = function() {
    // Set cookies
    if (!getCookie(cname)) {
      if (getLocalStorage(allowTracker) == 'true') {
        setCookie(cname, uuidv4(), 90)
      } else {
        setCookie(cname, '', 0)
      }
    } else {
      setCookie(cname, '', 0)
    }
  }

  tracker.trackCustom = function (e, d) {
    fetch(`${baseUri}/v1/tracking/custom?cts=${Date.now()}&e=${e}&d=${d}&uid=${getCookie(cname)}`)
      .then(r => r.text());
      // .then(data => console.log(data));
  }

  tracker.trackPageView = function (d) {
    fetch(`${baseUri}/v1/tracking/pageview?cts=${Date.now()}&d=${d}&uid=${getCookie(cname)}`)
      .then(r => r.text());
      // .then(data => console.log(data));
  }
  
  tracker.trackClick = function (d) {
    fetch(`${baseUri}/v1/tracking/click?cts=${Date.now()}&d=${d}&uid=${getCookie(cname)}`)
      .then(r => r.text());
      // .then(data => console.log(data));
  }

  tracker.registerUser = function () {
    fetch(`${baseUri}/v1/users/register?cts=${Date.now()}&uid=${getCookie(cname)}&ww=${window.screen.width}&wh=${window.screen.height}&waw=${window.screen.availWidth}&wah=${window.screen.availHeight}&o=${window.screen.orientation.type}`)
      .then(r => r.text());
      // .then(data => console.log(data));
  }

  tracker.initUserWithPageView = function(d) {
    tracker.initUser()
    tracker.trackPageView(d)
  }

  tracker.initUser()

  return tracker
}
