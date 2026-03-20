import { useEffect, useState } from "react";
import LoginPage from "./pages/LoginPage";
import Dashboard from "./pages/Dashboard";
import { getToken, setToken, clearToken } from "./jwt";
import { getLinks } from "./api";

type AuthState = "authenticated" | "unauthenticated";

export default function App() {
	const [authState, setAuthState] = useState<AuthState>(getToken() ? "authenticated" : "unauthenticated");
	const [notWhitelisted, setNotWhitelisted] = useState(false);

	function handleLogout() {
		clearToken();
		setAuthState("unauthenticated");
	}

	useEffect(() => {
		const params = new URLSearchParams(window.location.search);

		const token = params.get("token");
		if (token) {
			setToken(token);
			setAuthState("authenticated");
			window.history.replaceState({}, "", window.location.pathname);
			return;
		}

		const error = params.get("error");
		if (error) {
			clearToken();
			setNotWhitelisted(true);
			setAuthState("unauthenticated");
			window.history.replaceState({}, "", window.location.pathname);
			return;
		}

		if (getToken()) {
			getLinks()
				.then(response => {
					if (response.status === 401) {
						clearToken();
						setNotWhitelisted(true);
						setAuthState("unauthenticated");
					}
				})
				.catch(() => setAuthState("unauthenticated"));
		}
	}, []);

	if (authState === "unauthenticated") {
		return <LoginPage unauthorized={notWhitelisted} />;
	}

	return <Dashboard onLogout={handleLogout} />;
}
