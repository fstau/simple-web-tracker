# Web Tracker Concept

## Collected data

- Tracking across a single session
- Tracking across multiple sessions
- Traffic source via UTM parameters
- Pageviews
- Scrollpoints
- Click events
- Custom events (e.g. button clicks)

## Giving users the choice

- Consent: Users should have control over what tracking data they share with the site and be able to opt out of tracking entirely.
- Choice: Tracking of any category of event is a choice of the user.
- Visibility: Users should be able to see the data that was collected on their behavior.
- Ownership: If the user changes his mind about sharing his data, he can choose to trigger the anonymization (i.e. removing all session id information from the data points) or deletion of the collected data.

## Core concepts

- Preferences: The users tracking choices are saved in a localStorage entry to let tracker know (across sessions on the same device) which behavior to follow.
- User Agent: Represents a user on a device. If the user chooses to be tracked across multiple sessions his user agent is persisted in a cookie.
- Session: Represents on browser session (the users interaction with the site from first opening to closing the tab). Identification stored in session storage.
- Event: Interactions of the user with the site. This includes page views, scroll events, clicks and other custom events.
