SECRET_KEY = "quickstart-dev-key-change-in-production"
DEBUG = True
ALLOWED_HOSTS = ["*"]

INSTALLED_APPS = [
    "rest_framework",
    "api",
]

ROOT_URLCONF = "taskforge.urls"
DEFAULT_AUTO_FIELD = "django.db.models.BigAutoField"
