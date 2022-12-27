# Goldcrest

Goldcrest is an application allowing for per-user color roles on our Discord server. It is a purely HTTP-based interactions bot.

Setup:
  - Served via [Google App Engine](https://cloud.google.com/appengine/)
  - Token storage is managed via [Berglas](https://github.com/GoogleCloudPlatform/berglas), internally utilizing [Google Secret Manager](https://cloud.google.com/secret-manager)
  - Persistent data is via [Google Cloud Firestore](https://cloud.google.com/firestore/)
