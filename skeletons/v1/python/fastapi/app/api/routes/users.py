from fastapi import APIRouter

from app.crud import create_user, list_users
from app.schemas import UserCreate, UserResponse

router = APIRouter()


@router.get("/", response_model=list[UserResponse])
async def get_users() -> list[UserResponse]:
    return list_users()


@router.post("/", response_model=UserResponse, status_code=201)
async def post_user(user: UserCreate) -> UserResponse:
    return create_user(user)
