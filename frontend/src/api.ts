import { getToken } from "./jwt";

function request(input: RequestInfo, init?: RequestInit): Promise<Response> {
	const token = getToken();
	return fetch(input, {
		...init,
		headers: {
			...init?.headers,
			...(token ? { "Authorization": "Bearer " + token } : {}),
		},
	});
}

export interface Link {
	id: number;
	path: string;
	url: string;
}

export function getLinks(): Promise<Response> {
	return request("/api/links");
}

export function createLink(url: string, code?: string): Promise<Response> {
	const body: { Link: string; Code?: string } = { Link: url };
	if (code) {
		body.Code = code;
	}
	
	return request("/api/links/link", {
		method: "POST",
		headers: { "Content-Type": "application/json" },
		body: JSON.stringify(body),
	});
}

export function deleteLink(path: string): Promise<Response> {
	return request(`/api/links/link/${path}`, { method: "DELETE" });
}

export function inviteUser(email: string): Promise<Response> {
	return request("/api/auth/invite", {
		method: "POST",
		headers: { "Content-Type": "application/json" },
		body: JSON.stringify({ Email: email }),
	});
}
