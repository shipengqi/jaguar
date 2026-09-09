async def test_health(client):
    response = await client.get("/health")
    assert response.status_code == 200
    assert response.json() == {"status": "ok"}


async def test_list_users_empty(client):
    response = await client.get("/api/v1/users/")
    assert response.status_code == 200
    assert response.json() == []


async def test_create_user(client):
    response = await client.post("/api/v1/users/", json={"name": "Alice", "email": "alice@example.com"})
    assert response.status_code == 201
    data = response.json()
    assert data["name"] == "Alice"
    assert data["id"] == 1
