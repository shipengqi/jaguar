from app.schemas import UserCreate, UserResponse

_users: list[UserResponse] = []
_next_id = 1


def list_users() -> list[UserResponse]:
    return _users


def create_user(data: UserCreate) -> UserResponse:
    global _next_id
    user = UserResponse(id=_next_id, **data.model_dump())
    _users.append(user)
    _next_id += 1
    return user
