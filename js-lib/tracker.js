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

function setUserPreferences(track, user, session) {
  const settings = {
    track: track,
    preferences: {
      attribution: {
        user: user,
        session: session,
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

function trackEvent(baseUri, user, session, event, page, query, data) {
  fetch(`${baseUri}/v1/track`, {
    method: "POST",
    body: JSON.stringify({
      cts: new Date().getTime(),
      u: user,
      s: session,
      e: event,
      p: page,
      q: query,
      d: data,
    }),
  }).catch(console.error);
}

function getTracker(baseUri) {
  const settings = getUserPreferences();

  const tracker = {
    baseUri,
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
    if (!session) session = "anonymous";
    this.session = session;
  };

  tracker.init = function () {
    this.initSettings();
    this.initIdentity();
  };

  tracker.trackPageView = function (page) {
    const query = window.location.search;
    let p = window.location.pathname;
    if (page) p = page;
    if (this.track && this.eventPreferences.pageviews)
      trackEvent(this.baseUri, this.user, this.session, "pageview", p, query);
  };

  tracker.trackClick = function (data) {
    const query = window.location.search;
    const page = window.location.pathname;
    if (this.track && this.eventPreferences.clicks)
      trackEvent(
        this.baseUri,
        this.user,
        this.session,
        "click",
        page,
        query,
        data
      );
  };

  tracker.updateTracker = function (track, user, session) {
    setUserPreferences(track, user, session);
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
