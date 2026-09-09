from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    model_config = SettingsConfigDict(env_file=".env", env_ignore_empty=True, extra="ignore")

    app_name: str = "{{.App.Name}}"
    debug: bool = False
    database_url: str = "sqlite:///./{{.App.NormalizedName}}.db"


settings = Settings()
