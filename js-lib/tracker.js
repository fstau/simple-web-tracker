// Utilities
function uuidv4() {
  return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, function (c) {
    var r = (Math.random() * 16) | 0,
      v = c == "x" ? r : (r & 0x3) | 0x8;
    return v.toString(16);
  });
}

function getCookie(cname) {
  var name = cname + "=";
  var decodedCookie = decodeURIComponent(document.cookie);
  var ca = decodedCookie.split(";");
  for (var i = 0; i < ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) == " ") {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return null;
}

function setCookie(cname, cvalue, exdays) {
  var d = new Date();
  d.setTime(d.getTime() + exdays * 24 * 60 * 60 * 1000);
  var expires = "expires=" + d.toUTCString();
  document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

function getLocalStorage(key) {
  return window.localStorage.getItem(key);
}

function setLocalStorage(key, value) {
  return window.localStorage.setItem(key, value);
}

function getSessionStorage(key) {
  return window.sessionStorage.getItem(key);
}

function setSessionStorage(key, value) {
  return window.sessionStorage.setItem(key, value);
}

function getUserPreferences() {
  const settings = JSON.parse(getLocalStorage("swt_preferences"));
  return settings;
}

function setUserPreferences(track) {
  const settings = {
    track: track,
    preferences: {
      attribution: {
        user: true,
        session: true,
      },
      events: {
        pageviews: true,
        scrollpoints: true,
        clicks: true,
        custom: true,
      },
    },
  };
  setLocalStorage("swt_preferences", JSON.stringify(settings));
}

function openConsentModal() {
  const modal = document.querySelector("#consent_modal");
  modal.style.display = "block";
}

function closeConsentModal() {
  const modal = document.querySelector("#consent_modal");
  modal.style.display = "none";
}

function trackPageView(baseUri, user, session) {
  fetch(
    `${baseUri}/v1/tracking/pageview?cts=${Date.now()}&u=${user}&s=${session}`
  ).catch(console.error);
}

function getTracker() {
  const settings = getUserPreferences();

  let tracker = {
    baseUri: "http://localhost:5000",
    settings: settings,
  };

  tracker.initSettings = function () {
    this.track = this.settings.track;
    this.eventPreferences = this.settings.preferences.events;
  };

  tracker.initIdentity = function () {
    let user = getCookie("swt_user");
    if (
      !user &&
      this.settings.track &&
      this.settings.preferences.attribution.user
    ) {
      user = uuidv4();
      setCookie("swt_user", user);
    }
    if (!user) user = "anonymous";
    this.user = user;

    // Initialize session
    let session = getSessionStorage("swt_session");
    if (
      !session &&
      this.settings.track &&
      this.settings.preferences.attribution.session
    ) {
      session = uuidv4();
      setSessionStorage("swt_session", session);
    }
    if (!this.session) session = "anonymous";
    this.session = session;
  };

  tracker.init = function () {
    this.initSettings();
    this.initIdentity();
  };

  tracker.trackPageView = function () {
    if (this.track && this.eventPreferences.pageviews)
      trackPageView(this.baseUri, this.user, this.session);
  };

  tracker.updateTracker = function (track) {
    setUserPreferences(track);
    this.settings = getUserPreferences();
    this.init();
  };

  if (!tracker.settings) {
    openConsentModal();
    return tracker;
  }
  tracker.init();
  return tracker;
}

// Tracker
function initTracker(baseUri, cname, allowTracker) {
  let tracker = {};

  tracker.initUser = function () {
    // Set cookies
    if (!getCookie(cname)) {
      if (getLocalStorage(allowTracker) == "true") {
        setCookie(cname, uuidv4(), 90);
      } else {
        setCookie(cname, "", 0);
      }
    } else {
      setCookie(cname, "", 0);
    }
  };

  tracker.trackCustom = function (e, d) {
    fetch(
      `${baseUri}/v1/tracking/custom?cts=${Date.now()}&e=${e}&d=${d}&uid=${getCookie(
        cname
      )}`
    ).then((r) => r.text());
    // .then(data => console.log(data));
  };

  tracker.trackPageView = function (d) {
    fetch(
      `${baseUri}/v1/tracking/pageview?cts=${Date.now()}&d=${d}&uid=${getCookie(
        cname
      )}`
    ).then((r) => r.text());
    // .then(data => console.log(data));
  };

  tracker.trackClick = function (d) {
    fetch(
      `${baseUri}/v1/tracking/click?cts=${Date.now()}&d=${d}&uid=${getCookie(
        cname
      )}`
    ).then((r) => r.text());
    // .then(data => console.log(data));
  };

  tracker.registerUser = function () {
    fetch(
      `${baseUri}/v1/users/register?cts=${Date.now()}&uid=${getCookie(
        cname
      )}&ww=${window.screen.width}&wh=${window.screen.height}&waw=${
        window.screen.availWidth
      }&wah=${window.screen.availHeight}&o=${window.screen.orientation.type}`
    ).then((r) => r.text());
    // .then(data => console.log(data));
  };

  tracker.initUserWithPageView = function (d) {
    tracker.initUser();
    tracker.trackPageView(d);
  };

  tracker.initUser();

  return tracker;
}
