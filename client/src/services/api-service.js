const API_ENDPOINTS = "http://127.0.0.1:8080";

export async function loginRequest(username, password) {
  const response = await fetch(`${API_ENDPOINTS}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username,
      password,
    }),
  });

  return await response.json();
}

export async function registerRequest(username, password) {
  const response = await fetch(`${API_ENDPOINTS}/registration`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      username,
      password
    })
  });
  return await response.json();
}
