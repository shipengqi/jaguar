[project]
name = "{{.App.NormalizedName}}"
version = "0.1.0"
description = "{{.App.Name}} - FastAPI backend"
requires-python = ">=3.12"
dependencies = [
    "fastapi[standard]>=0.115.0",
    "pydantic>=2.7.0",
    "pydantic-settings>=2.2.0",
]

[dependency-groups]
dev = [
    "pytest>=8.0.0",
    "pytest-asyncio>=0.23.0",
    "httpx>=0.27.0",
    "ruff>=0.6.0",
    "coverage[toml]>=7.4.0",
]

[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

[tool.uv]
dev-dependencies = [
    "pytest>=8.0.0",
    "pytest-asyncio>=0.23.0",
    "httpx>=0.27.0",
    "ruff>=0.6.0",
    "coverage[toml]>=7.4.0",
]

[tool.ruff]
line-length = 100
target-version = "py312"
exclude = ["alembic"]

[tool.ruff.lint]
select = [
    "E",     # pycodestyle errors
    "W",     # pycodestyle warnings
    "F",     # pyflakes
    "I",     # isort
    "B",     # flake8-bugbear
    "C4",    # flake8-comprehensions
    "UP",    # pyupgrade
    "ARG001",# unused arguments
]
ignore = [
    "E501",  # line too long, handled by formatter
    "B008",  # function calls in argument defaults (FastAPI Depends)
    "B904",  # raise from e (allow bare raises in HTTP exceptions)
]

[tool.ruff.format]
quote-style = "double"
indent-style = "space"

[tool.pytest.ini_options]
asyncio_mode = "auto"
testpaths = ["tests"]

[tool.coverage.run]
source = ["app"]

[tool.coverage.report]
show_missing = true
