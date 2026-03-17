const TOKEN_KEY = "token";

export function getToken(): string | null {
    return localStorage.getItem(TOKEN_KEY);
}

export function setToken(token: string) {
    localStorage.setItem(TOKEN_KEY, token);
}

export function clearToken() {
    localStorage.removeItem(TOKEN_KEY);
}

export function request(input: RequestInfo, init?: RequestInit): Promise<Response> {
    const token = getToken();
    return fetch(input, {
        ...init,
        headers: {
            ...init?.headers,
            ...(token ? { "Authorization": "Bearer " + token } : {}),
        },
    });
}
